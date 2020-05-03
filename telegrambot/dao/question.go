package dao

import (
	"log"

	"github.com/go-pg/pg/v9/orm"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type QuestionID int64

func (p QuestionID) String() string {
	return string(p)
}

type Option struct {
	Text      string
	IsCorrect bool
}

type Question struct {
	tableName struct{} `pg:"question"`

	ID      QuestionID `pg:"id,pk"`
	Text    string     `pg:"text,notnull"`
	Options []Option   `pg:"options,array"`
}

type QuestionDAO interface {
	Upsert(question *Question) (*Question, error)
}

var _ QuestionDAO = (*questionDAO)(nil)

type questionDAO struct {
	conn *postgres.Connection
}

func NewQuestionDAO(conn *postgres.Connection) (*questionDAO, error) {
	dao := &questionDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *questionDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*Question)(nil), options)
}

func (dao *questionDAO) Upsert(question *Question) (*Question, error) {
	log.Printf("add question %d", question.ID)
	_, err := dao.conn.Model(question).
		OnConflict("(id) DO NOTHING").
		Insert(question)
	if err != nil {
		return nil, err
	}
	return question, nil
}
