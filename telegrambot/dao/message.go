package dao

import (
	"github.com/go-pg/pg/v9/orm"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type Message struct {
	ID     int
	ChatID int64
	Author string
	Text   string
}

type MessageDAO interface {
	CreateMessage(id int, chatID int64, author string, text string) (*Message, error)
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

func (dao *messageDAO) CreateMessage(id int, chatID int64, author string, text string) (*Message, error) {
	message := &Message{
		ID:     id,
		ChatID: chatID,
		Author: author,
		Text:   text,
	}
	if err := dao.conn.Insert(message); err != nil {
		return nil, err
	} else {
		return message, nil
	}
}
