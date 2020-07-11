package entity

import (
	"time"
)

type UserMemorizedTerm struct {
	tableName struct{} `pg:"userMemorizedTerm"`

	Timestamp           time.Time `pg:"timestamp,pk"`
	UserID              UserID    `pg:"user_id,pk"`
	Term                Term      `pg:"term,pk"`
	CorrectlyTranslated bool      `pg:"correctly_translated"`
}

func NewUserMemorizedTerm(userID UserID, term Term, correctlyTranslated bool) *UserMemorizedTerm {
	return &UserMemorizedTerm{
		Timestamp:           time.Now(),
		UserID:              userID,
		Term:                term,
		CorrectlyTranslated: correctlyTranslated,
	}
}
