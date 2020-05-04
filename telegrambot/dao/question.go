package dao

import (
	"time"

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

	ID        QuestionID `pg:"id,pk"`
	Text      string     `pg:"text,notnull"`
	Options   []Option   `pg:"options,array"`
	CreatedAt time.Time  `pg:"created_at,default:now()"`
	UpdatedAt time.Time  `pg:"updated_at,default:now()"`
}

type QuestionDAO interface {
	Find(questionID QuestionID) (*Question, error)
	Delete(questionID QuestionID) error
	Upsert(question *Question) error
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

func (dao *questionDAO) Find(questionID QuestionID) (*Question, error) {
	question := &Question{ID: questionID}
	err := dao.conn.Select(question)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (dao *questionDAO) Delete(questionID QuestionID) error {
	question := &Question{ID: questionID}
	return dao.conn.Delete(question)
}

func (dao *questionDAO) Upsert(question *Question) error {
	_, err := dao.conn.Model(question).
		OnConflict("(id) DO UPDATE").
		Set("updated_at = now()").
		Set("text = EXCLUDED.text").
		Set("options = EXCLUDED.options").
		Insert()
	return err
}
