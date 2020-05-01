package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const timeout = 10

type API interface {
	GetNickName() string
	ListenMessages(handlerFunc func(message *dao.Message) error) error
	Reply(message *dao.Message, text string) error
}

var _ API = (*api)(nil)

type api struct {
	botAPI     *tgbotapi.BotAPI
	messageDAO dao.MessageDAO
}

func NewAPI(conn *postgres.Connection) (*api, error) {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, os.ErrInvalid
	}
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	messageDAO, err := dao.NewMessageDAO(conn)
	if err != nil {
		return nil, err
	}
	return &api{
		botAPI:     botAPI,
		messageDAO: messageDAO,
	}, nil
}

func (api *api) GetNickName() string {
	return api.botAPI.Self.UserName
}

func (api *api) ListenMessages(handlerFunc func(message *dao.Message) error) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	updates, err := api.botAPI.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		msg := update.Message
		if msg == nil {
			return nil
		}
		log.Printf("[%s] %s", msg.From.UserName, msg.Text)
		message, err := api.messageDAO.CreateMessage(
			msg.MessageID,
			msg.Chat.ID,
			msg.From.UserName,
			msg.Text,
		)
		if err != nil {
			return err
		}
		if err := handlerFunc(message); err != nil {
			return err
		}
	}
	return nil
}

func (api *api) Reply(message *dao.Message, text string) error {
	replicationMessage := tgbotapi.NewMessage(message.ChatID, text)
	replicationMessage.ReplyToMessageID = message.ID
	_, err := api.botAPI.Send(replicationMessage)
	return err
}
