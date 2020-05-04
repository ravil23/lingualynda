package dao

import (
	"github.com/go-pg/pg/v9/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type ServerPollID string

func (p ServerPollID) String() string {
	return string(p)
}

type PollAnswer struct {
	tableName struct{} `pg:"poll_answer"`

	ChosenOptions []PollOptionID `pg:"chosen_options,array"`

	ServerPollID ServerPollID `pg:"server_poll_id,pk"`

	UserID UserID `pg:"user_id,pk"`
	User   *User  `pg:"fk:user_id"`
}

func NewPollAnswer(tgPollAnswer *tgbotapi.PollAnswer, user *User) *PollAnswer {
	pollAnswer := &PollAnswer{
		ServerPollID:  ServerPollID(tgPollAnswer.PollID),
		UserID:        user.ID,
		User:          user,
		ChosenOptions: make([]PollOptionID, 0, len(tgPollAnswer.OptionIDs)),
	}
	for _, tgOptionID := range tgPollAnswer.OptionIDs {
		pollAnswer.ChosenOptions = append(pollAnswer.ChosenOptions, PollOptionID(tgOptionID))
	}
	return pollAnswer
}

type PollAnswerDAO interface {
	Upsert(pollAnswer *PollAnswer) error
}

var _ PollAnswerDAO = (*pollAnswerDAO)(nil)

type pollAnswerDAO struct {
	conn *postgres.Connection
}

func NewPollAnswerDAO(conn *postgres.Connection) (*pollAnswerDAO, error) {
	dao := &pollAnswerDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *pollAnswerDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*PollAnswer)(nil), options)
}

func (dao *pollAnswerDAO) Upsert(pollAnswer *PollAnswer) error {
	_, err := dao.conn.Model(pollAnswer).
		OnConflict("(user_id, server_poll_id) DO NOTHING").
		Insert()
	return err
}
