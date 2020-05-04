package telegram

import (
	"log"
	"math/rand"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/collection"
	"github.com/ravil23/lingualynda/telegrambot/collection/paulineunit1"
	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const (
	timeout                 = 10
	maxQuestionOptionsCount = 4
)

type API interface {
	SetMessagesHandler(handlerFunc func(*dao.Message) error)
	SetPollAnswersHandler(handlerFunc func(*dao.PollAnswer) error)
	ListenUpdates() error
	SendNextPoll(user *dao.User) error
}

var _ API = (*api)(nil)

type api struct {
	botAPI              *tgbotapi.BotAPI
	messageDAO          dao.MessageDAO
	pollDAO             dao.PollDAO
	pollAnswerDAO       dao.PollAnswerDAO
	userDAO             dao.UserDAO
	linkUserQuestionDAO dao.LinkUserQuestionDAO
	questionDAO         dao.QuestionDAO

	messagesHandler    func(update *tgbotapi.Update) error
	pollAnswersHandler func(update *tgbotapi.Update) error
}

func NewAPI(botToken string, conn *postgres.Connection) (*api, error) {
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	userDAO, err := dao.NewUserDAO(conn)
	if err != nil {
		return nil, err
	}
	questionDAO, err := dao.NewQuestionDAO(conn)
	if err != nil {
		return nil, err
	}
	messageDAO, err := dao.NewMessageDAO(conn)
	if err != nil {
		return nil, err
	}
	pollDAO, err := dao.NewPollDAO(conn)
	if err != nil {
		return nil, err
	}
	pollAnswerDAO, err := dao.NewPollAnswerDAO(conn)
	if err != nil {
		return nil, err
	}
	linkUserQuestionDAO, err := dao.NewLinkUserQuestionDAO(conn)
	if err != nil {
		return nil, err
	}
	return &api{
		botAPI:              botAPI,
		messageDAO:          messageDAO,
		pollDAO:             pollDAO,
		pollAnswerDAO:       pollAnswerDAO,
		userDAO:             userDAO,
		linkUserQuestionDAO: linkUserQuestionDAO,
		questionDAO:         questionDAO,
	}, nil
}

func (api *api) SetMessagesHandler(handlerFunc func(*dao.Message) error) {
	api.messagesHandler = func(update *tgbotapi.Update) error {
		user := dao.NewUser(update.Message.From)
		user.ChatID = dao.ChatID(update.Message.Chat.ID)
		if err := api.userDAO.Upsert(user); err != nil {
			return err
		}

		message := dao.NewMessage(update.Message, user)
		if err := api.messageDAO.Upsert(message); err != nil {
			return err
		}

		if err := handlerFunc(message); err != nil {
			return err
		}
		return nil
	}
}

func (api *api) SetPollAnswersHandler(handlerFunc func(*dao.PollAnswer) error) {
	api.pollAnswersHandler = func(update *tgbotapi.Update) error {
		user := dao.NewUser(&update.PollAnswer.User)
		if err := api.userDAO.Upsert(user); err != nil {
			return err
		}

		pollAnswer := dao.NewPollAnswer(update.PollAnswer, user)
		if err := api.pollAnswerDAO.Upsert(pollAnswer); err != nil {
			return err
		}

		return handlerFunc(pollAnswer)
	}
}

func (api *api) ListenUpdates() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	updates := api.botAPI.GetUpdatesChan(u)
	for update := range updates {
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

func (api *api) SendNextPoll(user *dao.User) error {
	poll, err := api.getNextPoll(user)
	if err != nil {
		return err
	}

	tgPoll := poll.ToChatable(user.ChatID)
	_, err = api.botAPI.Send(tgPoll)
	return err
}

func (api *api) getNextPoll(user *dao.User) (*dao.Poll, error) {
	question := generateRandomQuestion()
	if err := api.questionDAO.Upsert(question); err != nil {
		return nil, err
	}

	linkUserQuestion := dao.NewLinkUserQuestion(question.ID, user.ID)
	if err := api.linkUserQuestionDAO.Upsert(linkUserQuestion); err != nil {
		return nil, err
	}

	poll := &dao.Poll{
		Type:       dao.PollTypeQuiz,
		QuestionID: question.ID,
		Question:   question,
		IsPublic:   true,
	}
	if err := api.pollDAO.Upsert(poll); err != nil {
		return nil, err
	}
	return poll, nil
}

func generateRandomQuestion() *dao.Question {
	term := paulineunit1.VocabularyTotal.GetRandomTerm()
	correctTranslations := paulineunit1.VocabularyTotal.GetTranslations(term)
	correctTranslation := correctTranslations[rand.Intn(len(correctTranslations))]
	question := &dao.Question{
		Text: term.String(),
		Options: []dao.Option{
			{Text: correctTranslation.String(), IsCorrect: true},
		},
	}
	for len(question.Options) < maxQuestionOptionsCount {
		randomTranslation := collection.VocabularyTotal.GetRandomTranslation()
		isValidTranslation := true
		for _, correctTranslation := range correctTranslations {
			if randomTranslation == correctTranslation {
				isValidTranslation = false
				break
			}
		}
		if isValidTranslation {
			question.Options = append(question.Options, dao.Option{
				Text:      randomTranslation.String(),
				IsCorrect: false,
			})
		}
	}
	rand.Shuffle(len(question.Options), func(i, j int) {
		question.Options[i], question.Options[j] = question.Options[j], question.Options[i]
	})
	return question
}

func GetBotTokenOrPanic() string {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Panic("bot token is empty")
	}
	return botToken
}
