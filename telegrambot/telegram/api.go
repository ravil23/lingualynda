package telegram

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/collection"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/entity"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const (
	timeout                 = 10
	maxQuestionOptionsCount = 4
	alertsChatID            = -1001142742669
	botNickName             = "LinguaLyndaBot"
	botMention              = "@" + botNickName
)

type API interface {
	SetMessagesHandler(handlerFunc func(*entity.Message) error)
	SetPollAnswersHandler(handlerFunc func(*entity.User, *entity.PollAnswer) error)
	ListenUpdates() error
	SendNextPoll(user *entity.User) error
	SendAlert(text string)
	SendMessage(chatID entity.ChatID, text string)
	SendHTMLMessage(chatID entity.ChatID, text string)
}

var _ API = (*api)(nil)

type api struct {
	hostName string

	tgAPI     *tgbotapi.BotAPI
	tgUpdates tgbotapi.UpdatesChannel

	userDAO dao.UserDAO

	userProfileManager *UserProfileManager

	messagesHandler    func(update *tgbotapi.Update) error
	pollAnswersHandler func(update *tgbotapi.Update) error
}

func NewAPI(botToken string, conn *postgres.Connection) (*api, error) {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown_host"
	}

	tgAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	tgUpdates := tgAPI.GetUpdatesChan(u)

	userDAO, err := dao.NewUserDAO(conn)
	if err != nil {
		return nil, err
	}

	userProfileManager, err := NewUserProfileManager(conn, userDAO)
	if err != nil {
		return nil, err
	}

	return &api{
		hostName:  hostName,
		tgUpdates: tgUpdates,
		tgAPI:     tgAPI,

		userDAO: userDAO,

		userProfileManager: userProfileManager,
	}, nil
}

func (api *api) SetMessagesHandler(handlerFunc func(*entity.Message) error) {
	api.messagesHandler = func(update *tgbotapi.Update) error {
		user := entity.NewUser(update.Message.From)
		user.ChatID = entity.ChatID(update.Message.Chat.ID)
		if err := api.userDAO.Upsert(user); err != nil {
			return err
		}

		if update.Message.Command() == "start" {
			api.SendAlert(fmt.Sprintf("%s started conversation with %s", user.GetFormattedName(), botMention))
		}

		message := entity.NewMessage(update.Message, user)

		if err := handlerFunc(message); err != nil {
			return err
		}
		return nil
	}
}

func (api *api) SetPollAnswersHandler(handlerFunc func(*entity.User, *entity.PollAnswer) error) {
	api.pollAnswersHandler = func(update *tgbotapi.Update) error {
		user, err := api.userDAO.Find(entity.UserID(update.PollAnswer.User.ID))
		if err != nil {
			return err
		}

		pollAnswer := entity.NewPollAnswer(update.PollAnswer)

		if err := api.userProfileManager.AddPollAnswer(user.ID, pollAnswer); err != nil {
			return err
		}

		return handlerFunc(user, pollAnswer)
	}
}

func (api *api) ListenUpdates() error {
	for update := range api.tgUpdates {
		log.Printf("Handle update: %+v", update)
		if update.Message != nil {
			if err := api.messagesHandler(&update); err != nil {
				return err
			}
		}
		if update.PollAnswer != nil {
			if err := api.pollAnswersHandler(&update); err != nil {
				return err
			}
		}

	}
	return nil
}

func (api *api) SendNextPoll(user *entity.User) error {
	poll := api.getNextPoll(user)
	tgPoll := poll.ToChatable(user.ChatID)
	tgMessage, err := api.tgAPI.Send(tgPoll)
	if err != nil {
		return err
	}
	if tgMessage.Poll == nil {
		return fmt.Errorf("returned message does not contain poll: %+v", tgMessage)
	}
	poll.ID = entity.PollID(tgMessage.Poll.ID)
	api.userProfileManager.AddPoll(poll)
	return nil
}

func (api *api) getNextPoll(user *entity.User) *entity.Poll {
	var listOfVocabularies []*schema.Vocabulary
	if chat, found := chatsStates[user.ChatID]; found && len(chat.GetVocabularies()) > 0 {
		listOfVocabularies = chat.GetVocabularies()
	} else {
		listOfVocabularies = []*schema.Vocabulary{collection.VocabularyEngToRus, collection.VocabularyRusToEng}
	}
	selectedVocabulary := listOfVocabularies[rand.Intn(len(listOfVocabularies))]
	term := selectedVocabulary.GetRandomTerm()
	correctTranslations := selectedVocabulary.GetTranslations(term)
	correctTranslation := correctTranslations[rand.Intn(len(correctTranslations))]
	poll := &entity.Poll{
		Term:     term,
		Type:     entity.PollTypeQuiz,
		IsPublic: true,
		Options: []*entity.PollOption{
			{Translation: correctTranslation, IsCorrect: true},
		},
	}
	for len(poll.Options) < maxQuestionOptionsCount {
		randomTranslation := selectedVocabulary.GetRandomTranslation()
		isValidTranslation := true
		for _, correctTranslation := range correctTranslations {
			if randomTranslation == correctTranslation {
				isValidTranslation = false
				break
			}
		}
		if isValidTranslation {
			poll.Options = append(poll.Options, &entity.PollOption{
				Translation: randomTranslation,
				IsCorrect:   false,
			})
		}
	}
	rand.Shuffle(len(poll.Options), func(i, j int) {
		poll.Options[i], poll.Options[j] = poll.Options[j], poll.Options[i]
	})
	return poll
}

func (api *api) SendAlert(text string) {
	api.SendMessage(alertsChatID, fmt.Sprintf("[%s] %s", api.hostName, text))
}

func (api *api) SendMessage(chatID entity.ChatID, text string) {
	api.sendMessage(chatID, text, "")
}

func (api *api) SendHTMLMessage(chatID entity.ChatID, text string) {
	api.sendMessage(chatID, text, tgbotapi.ModeHTML)
}

func (api *api) sendMessage(chatID entity.ChatID, text string, parseMode string) {
	log.Printf("Chat ID: %d, Parse mode: %s, Text: %s", chatID, parseMode, text)
	tgMessage := tgbotapi.NewMessage(int64(chatID), text)
	tgMessage.ParseMode = parseMode
	_, err := api.tgAPI.Send(tgMessage)
	if err != nil {
		log.Printf("Error on sending message: %s", err)
	}
}

func GetBotTokenOrPanic() string {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Panic("bot token is empty")
	}
	return botToken
}
