package telegram

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/entity"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const (
	timeout             = 10
	maxPollOptionsCount = 4
	alertsChatID        = -1001142742669
	botNickName         = "LinguaLyndaBot"
	botMention          = "@" + botNickName
)

type API interface {
	SetMessagesHandler(handlerFunc func(*entity.Message) error)
	SetPollAnswersHandler(handlerFunc func(*entity.User, *entity.PollAnswer) error)
	ListenUpdates() error
	SendNextPoll(user *entity.User) error
	SendAlert(text string)
	SendMessage(chatID entity.ChatID, text string)
	SendHTMLMessage(chatID entity.ChatID, text string)
	SendProgress(user *entity.User)
	UpdateInternalState(message *entity.Message)
}

var _ API = (*api)(nil)

type api struct {
	hostName string

	tgAPI     *tgbotapi.BotAPI
	tgUpdates tgbotapi.UpdatesChannel

	userDAO dao.UserDAO

	chatManager        *ChatManager
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

		chatManager:        NewChatManager(),
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
	poll, found := api.getNextPoll(user)
	if !found {
		chat := api.chatManager.GetChatOrCreate(user.ChatID)
		headerText := fmt.Sprintf(
			"<b>Congratulations!</b>\nYou have memorized all terms from /%s vocabulary in /%s mode.",
			chat.GetVocabularyType(),
			chat.GetMode(),
		)
		progressText := api.getProgressByUser(user)
		text := strings.Join([]string{
			headerText,
			progressText,
			"",
			"Please change vocabulary or mode. Type /help to see instructions.",
		}, "\n")
		api.SendHTMLMessage(user.ChatID, text)
		return nil
	}
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
	if chat := api.chatManager.GetChatOrCreate(user.ChatID); chat.IsDebuggingEnabled() {
		api.sendDebugMessage(chat, user, poll)
	}
	return nil
}

func (api *api) getNextPoll(user *entity.User) (*entity.Poll, bool) {
	chat := api.chatManager.GetChatOrCreate(user.ChatID)
	listOfVocabularies := chat.GetListOfVocabularies()
	vocabularyIndex := rand.Intn(len(listOfVocabularies))
	selectedVocabulary := listOfVocabularies[vocabularyIndex]
	var term entity.Term
	var weight float64
	var finished bool
	if userProfile, found := api.userProfileManager.GetUserProfile(user.ID); found {
		for i := range listOfVocabularies {
			selectedVocabulary = listOfVocabularies[(vocabularyIndex+i)%len(listOfVocabularies)]
			term, weight, finished = selectedVocabulary.GetTermByUserProfile(userProfile)
			if !finished {
				break
			}
		}
		if finished {
			return nil, false
		}
	} else {
		term = selectedVocabulary.GetRandomTerm()
	}
	correctTranslations := selectedVocabulary.GetTranslations(term)
	correctTranslation := correctTranslations[rand.Intn(len(correctTranslations))]
	poll := &entity.Poll{
		Term:     term,
		Weight:   weight,
		Type:     entity.PollTypeQuiz,
		IsPublic: true,
		Options: []*entity.PollOption{
			{Translation: correctTranslation, IsCorrect: true},
		},
	}
	for len(poll.Options) < maxPollOptionsCount {
		randomTranslation := selectedVocabulary.GetRandomTranslation()
		if poll.IsExistedOption(randomTranslation) {
			continue
		}
		poll.Options = append(poll.Options, &entity.PollOption{
			Translation: randomTranslation,
			IsCorrect:   false,
		})
	}
	rand.Shuffle(len(poll.Options), func(i, j int) {
		poll.Options[i], poll.Options[j] = poll.Options[j], poll.Options[i]
	})
	return poll, true
}

func (api *api) SendProgress(user *entity.User) {
	api.SendMessage(user.ChatID, api.getProgressByUser(user))
}

func (api *api) getProgressByUser(user *entity.User) string {
	chat := api.chatManager.GetChatOrCreate(user.ChatID)
	userProfile, _ := api.userProfileManager.GetUserProfile(user.ID)
	currentText := api.getProgressByChat(chat, userProfile, false)
	otherTexts := make([]string, 0, len(entity.AllChatVocabularyTypes))
	for _, vocabularyType := range entity.AllChatVocabularyTypes {
		if vocabularyType == chat.GetVocabularyType() {
			continue
		}
		otherChat := entity.NewChat(0)
		api.chatManager.SetupChatConfiguration(otherChat, entity.ChatModeAllDirections, vocabularyType)
		otherTexts = append(otherTexts, api.getProgressByChat(otherChat, userProfile, true))
	}
	return strings.Join([]string{
		currentText,
		"",
		"Other vocabularies:",
		strings.Join(otherTexts, "\n"),
	}, "\n")
}

func (api *api) getProgressByChat(chat *entity.Chat, userProfile *entity.UserProfile, short bool) string {
	totalTermsCount := 0
	correctMemorizedTermsCount := 0
	for _, vocabulary := range chat.GetListOfVocabularies() {
		totalTermsCount += vocabulary.GetTermsCount()
		if userProfile != nil {
			correctMemorizedTermsCount += vocabulary.GetCorrectMemorizedTermsCount(userProfile)
		}
	}
	if short {
		return fmt.Sprintf(
			"/%s: %s (%d terms from %d memorized)",
			chat.GetVocabularyType(),
			fmt.Sprintf("%.1f%%", 100*float64(correctMemorizedTermsCount)/float64(totalTermsCount)),
			correctMemorizedTermsCount,
			totalTermsCount,
		)
	} else {
		return fmt.Sprintf(
			"Progress of /%s vocabulary in /%s mode is %s (%d terms from %d memorized)",
			chat.GetVocabularyType(),
			chat.GetMode(),
			fmt.Sprintf("%.1f%%", 100*float64(correctMemorizedTermsCount)/float64(totalTermsCount)),
			correctMemorizedTermsCount,
			totalTermsCount,
		)
	}
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

func (api *api) UpdateInternalState(message *entity.Message) {
	api.chatManager.UpdateChatConfigurations(message.ChatID, message.Text)
}

func (api *api) sendDebugMessage(chat *entity.Chat, user *entity.User, poll *entity.Poll) {
	debugMessage := fmt.Sprintf("\nUser: %s", user.GetFormattedName())
	debugMessage += fmt.Sprintf("\nChat ID: %d", chat.GetID())
	debugMessage += fmt.Sprintf("\nSelected mode: %s", chat.GetMode())
	debugMessage += fmt.Sprintf("\nSelected vocabulary type: %s", chat.GetVocabularyType())
	debugMessage += fmt.Sprintf("\nSelected vocabularies count: %d", len(chat.GetListOfVocabularies()))
	if poll != nil {
		debugMessage += fmt.Sprintf("\nPoll: %+v with weight %.3f", poll.Term, poll.Weight)
	}
	api.SendAlert(debugMessage)
}

func GetBotTokenOrPanic() string {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Panic("bot token is empty")
	}
	return botToken
}
