package dao

import (
	"log"
	"sort"

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

	ID              PollID                  `pg:"id,pk"` // FIXME: local id which doesn't send to api
	Type            PollType                `pg:"type,notnull"`
	Question        string                  `pg:"question,notnull"`
	Options         map[PollOptionID]string `pg:"options,notnull"`
	CorrectOptionID PollOptionID            `pg:"correct_option_id,use_zero"`
	IsPublic        bool                    `pg:"is_public,use_zero"`
}

func (p *Poll) ToChatable(chatID ChatID) *tgbotapi.SendPollConfig {
	tgOptions := p.getSortedOptions()
	tgPoll := tgbotapi.NewPoll(int64(chatID), p.Question, tgOptions...)
	tgPoll.CorrectOptionID = int64(p.CorrectOptionID)
	tgPoll.Type = p.Type.String()
	tgPoll.IsAnonymous = !p.IsPublic
	return &tgPoll
}

func (p *Poll) getSortedOptions() []string {
	optionIDs := make([]PollOptionID, 0, len(p.Options))
	for optionID := range p.Options {
		optionIDs = append(optionIDs, optionID)
	}
	sort.Slice(optionIDs, func(i, j int) bool {
		return optionIDs[i] < optionIDs[j]
	})
	options := make([]string, 0, len(p.Options))
	for _, optionID := range optionIDs {
		options = append(options, p.Options[optionID])
	}
	return options
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
	log.Printf("add poll %s of type %s", poll.ID, poll.Type)
	_, err := dao.conn.Model(poll).
		OnConflict("(id) DO NOTHING").
		Insert(poll)
	if err != nil {
		return nil, err
	}
	return poll, nil
}
