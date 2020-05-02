package dao

import (
	"encoding/json"

	"github.com/go-pg/pg/v9/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type Message struct {
	ID        int
	Timestamp int
	ChatID    int64
	UserID    int
	UserName  string
	Text      string
	Dump      string
}

type MessageDAO interface {
	CreateMessage(message *tgbotapi.Message) (*Message, error)
}

var _ MessageDAO = (*messageDAO)(nil)

type messageDAO struct {
	conn *postgres.Connection
}

func NewMessageDAO(conn *postgres.Connection) (*messageDAO, error) {
	dao := &messageDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *messageDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	return dao.conn.CreateTable((*Message)(nil), options)
}

func (dao *messageDAO) CreateMessage(msg *tgbotapi.Message) (*Message, error) {
	dump, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	message := &Message{
		ID:        msg.MessageID,
		Timestamp: msg.Date,
		ChatID:    msg.Chat.ID,
		UserID:    msg.From.ID,
		UserName:  msg.From.UserName,
		Text:      msg.Text,
		Dump:      string(dump),
	}
	if err := dao.conn.Insert(message); err != nil {
		return nil, err
	} else {
		return message, nil
	}
}
