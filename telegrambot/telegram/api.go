package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const timeout = 10

type API interface {
	SetMessagesHandler(handlerFunc func(*dao.Message) error)
	SetPollAnswersHandler(handlerFunc func(*dao.PollAnswer) error)
	ListenUpdates() error
	SendNextPoll(user *dao.User) error
}

var _ API = (*api)(nil)

type api struct {
	botAPI        *tgbotapi.BotAPI
	messageDAO    dao.MessageDAO
	pollDAO       dao.PollDAO
	pollAnswerDAO dao.PollAnswerDAO
	userDAO       dao.UserDAO

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
	return &api{
		botAPI:        botAPI,
		messageDAO:    messageDAO,
		pollDAO:       pollDAO,
		pollAnswerDAO: pollAnswerDAO,
		userDAO:       userDAO,
	}, nil
}

func (api *api) SetMessagesHandler(handlerFunc func(*dao.Message) error) {
	api.messagesHandler = func(update *tgbotapi.Update) error {
		user := dao.NewUser(update.Message.From)
		user.ChatID = dao.ChatID(update.Message.Chat.ID)
		user, err := api.userDAO.Upsert(user)
		if err != nil {
			return err
		}
		message := dao.NewMessage(update.Message, user)
		message, err = api.messageDAO.Upsert(message)
		if err != nil {
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
		user, err := api.userDAO.Upsert(user)
		if err != nil {
			return err
		}
		user, err = api.userDAO.Find(user.ID) // TODO: merge upsert and find
		if err != nil {
			return err
		}
		pollAnswer := dao.NewPollAnswer(update.PollAnswer, user)
		pollAnswer, err = api.pollAnswerDAO.Upsert(pollAnswer)
		if err != nil {
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
	if poll, err = api.pollDAO.Upsert(poll); err != nil { // FIXME: remove and write to users_poll table
		return err
	}
	tgPoll := poll.ToChatable(user.ChatID)
	_, err = api.botAPI.Send(tgPoll)
	return err
}

func (api *api) getNextPoll(user *dao.User) (*dao.Poll, error) {
	poll := &dao.Poll{
		ID:              "test",
		Type:            dao.PollTypeQuiz,
		Question:        "Test question",
		Options:         map[dao.PollOptionID]string{0: "a", 1: "b"},
		CorrectOptionID: 0,
		IsPublic:        true,
	}
	return poll, nil
}

func GetBotTokenOrPanic() string {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Panic("bot token is empty")
	}
	return botToken
}
