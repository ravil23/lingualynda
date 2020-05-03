package dao

import (
	"log"

	"github.com/go-pg/pg/v9/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type PollID string
type PollType string
type PollOptionID int64

const (
	PollTypeQuiz = "quiz"
)

func (p PollID) String() string {
	return string(p)
}

func (t PollType) String() string {
	return string(t)
}

type Poll struct {
	tableName struct{} `pg:"poll"`

	ID       PollID   `pg:"poll_id,pk"`
	Type     PollType `pg:"type,notnull"`
	IsPublic bool     `pg:"is_public,use_zero"`

	QuestionID QuestionID `pg:"question_id,notnull"`
	Question   *Question  `pg:"fk:question_id"`
}

func (p *Poll) ToChatable(chatID ChatID) *tgbotapi.SendPollConfig {
	correctOptionID := -1
	tgOptions := make([]string, 0, len(p.Question.Options))
	for i, option := range p.Question.Options {
		tgOptions = append(tgOptions, option.Text)
		if option.IsCorrect {
			correctOptionID = i
		}
	}
	tgPoll := tgbotapi.NewPoll(int64(chatID), p.Question.Text, tgOptions...)
	tgPoll.CorrectOptionID = int64(correctOptionID)
	tgPoll.Type = p.Type.String()
	tgPoll.IsAnonymous = !p.IsPublic
	return &tgPoll
}

type PollDAO interface {
	Upsert(poll *Poll) (*Poll, error)
}

var _ PollDAO = (*pollDAO)(nil)

type pollDAO struct {
	conn *postgres.Connection
}

func NewPollDAO(conn *postgres.Connection) (*pollDAO, error) {
	dao := &pollDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *pollDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*Poll)(nil), options)
}

func (dao *pollDAO) Upsert(poll *Poll) (*Poll, error) {
	log.Printf("add poll `%s` of type `%s`", poll.ID, poll.Type)
	_, err := dao.conn.Model(poll).
		OnConflict("(poll_id) DO NOTHING").
		Insert(poll)
	if err != nil {
		return nil, err
	}
	return poll, nil
}
