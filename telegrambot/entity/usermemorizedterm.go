package entity

import (
	"time"

	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

type UserMemorizedTerm struct {
	tableName struct{} `pg:"userMemorizedTerm"`

	Timestamp           time.Time   `pg:"timestamp,pk"`
	UserID              UserID      `pg:"user_id,pk"`
	Term                schema.Term `pg:"term,pk"`
	CorrectlyTranslated bool        `pg:"correctly_translated"`
}

func NewUserMemorizedTerm(userID UserID, term schema.Term, correctlyTranslated bool) *UserMemorizedTerm {
	return &UserMemorizedTerm{
		Timestamp:           time.Now(),
		UserID:              userID,
		Term:                term,
		CorrectlyTranslated: correctlyTranslated,
	}
}
