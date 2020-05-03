package dao

import (
	"log"

	"github.com/go-pg/pg/v9/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type PollAnswer struct {
	tableName struct{} `pg:"poll_answer"`

	ChosenOptions []PollOptionID `pg:"chosen_options,array"`

	PollID PollID `pg:"poll_id,pk"`
	//Poll   *Poll  `pg:"fk:poll_id"` // TODO: uncomment

	UserID UserID `pg:"user_id,pk"`
	User   *User  `pg:"fk:user_id"`
}

func NewPollAnswer(tgPollAnswer *tgbotapi.PollAnswer, user *User) *PollAnswer {
	pollAnswer := &PollAnswer{
		PollID:        PollID(tgPollAnswer.PollID),
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
	Upsert(pollAnswer *PollAnswer) (*PollAnswer, error)
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

func (dao *pollAnswerDAO) Upsert(pollAnswer *PollAnswer) (*PollAnswer, error) {
	log.Printf("[user=%d][chat=%d] answer for poll %s", pollAnswer.UserID, pollAnswer.User.ChatID, pollAnswer.PollID)
	_, err := dao.conn.Model(pollAnswer).
		OnConflict("(user_id, poll_id) DO NOTHING").
		Insert(pollAnswer)
	if err != nil {
		return nil, err
	}
	return pollAnswer, nil
}
