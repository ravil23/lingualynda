package dao

import (
	"time"

	"github.com/go-pg/pg/v9/orm"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type LinkUserQuestion struct {
	tableName struct{} `pg:"userQuestionLink"`

	Timestamp  time.Time  `pg:"timestamp,pk"`
	UserID     UserID     `pg:"user_id,pk"`
	QuestionID QuestionID `pg:"question_id,pk"`
}

func NewLinkUserQuestion(questionID QuestionID, userID UserID) *LinkUserQuestion {
	return &LinkUserQuestion{
		Timestamp:  time.Now(),
		QuestionID: questionID,
		UserID:     userID,
	}
}

type LinkUserQuestionDAO interface {
	Upsert(userQuestionLink *LinkUserQuestion) error
}

var _ LinkUserQuestionDAO = (*userQuestionLinkDAO)(nil)

type userQuestionLinkDAO struct {
	conn *postgres.Connection
}

func NewLinkUserQuestionDAO(conn *postgres.Connection) (*userQuestionLinkDAO, error) {
	dao := &userQuestionLinkDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *userQuestionLinkDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*LinkUserQuestion)(nil), options)
}

func (dao *userQuestionLinkDAO) Upsert(linkUserQuestion *LinkUserQuestion) error {
	_, err := dao.conn.Model(linkUserQuestion).
		OnConflict("(timestamp, user_id, question_id) DO NOTHING").
		Insert()
	return err
}
