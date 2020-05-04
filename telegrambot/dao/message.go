package dao

import (
	"time"

	"github.com/go-pg/pg/v9/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type MessageID int
type ChatID int64

type Message struct {
	tableName struct{} `pg:"message"`

	ID        MessageID         `pg:"id,pk"`
	ChatID    ChatID            `pg:"chat_id,pk"`
	Timestamp time.Time         `pg:"timestamp,notnull"`
	Text      string            `pg:"text"`
	Raw       *tgbotapi.Message `pg:"raw,notnull"`

	UserID UserID `pg:"user_id,notnull"`
	User   *User  `pg:"fk:user_id"`
}

func NewMessage(tgMessage *tgbotapi.Message, user *User) *Message {
	return &Message{
		ID:        MessageID(tgMessage.MessageID),
		ChatID:    ChatID(tgMessage.Chat.ID),
		Timestamp: tgMessage.Time(),
		UserID:    user.ID,
		User:      user,
		Text:      tgMessage.Text,
		Raw:       tgMessage,
	}
}

type MessageDAO interface {
	Upsert(message *Message) error
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
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*Message)(nil), options)
}

func (dao *messageDAO) Upsert(message *Message) error {
	_, err := dao.conn.Model(message).
		OnConflict("(id, chat_id) DO NOTHING").
		Insert()
	return err
}
