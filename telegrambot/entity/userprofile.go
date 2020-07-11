package entity

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

type MemorizationWeight int

type UserProfile struct {
	userID                    UserID
	correctlyTranslatedTerms  map[schema.Term]int
	mistakenlyTranslatedTerms map[schema.Term]int
}

func NewUserProfile(userID UserID) *UserProfile {
	return &UserProfile{
		userID:                    userID,
		correctlyTranslatedTerms:  make(map[schema.Term]int),
		mistakenlyTranslatedTerms: make(map[schema.Term]int),
	}
}

func (p *UserProfile) AddCorrectlyTranslatedTerm(term schema.Term) {
	p.correctlyTranslatedTerms[term]++
}

func (p *UserProfile) AddMistakenlyTranslatedTerm(term schema.Term) {
	p.mistakenlyTranslatedTerms[term]++
}

func (p *UserProfile) GetMemorizationWeight(term schema.Term) MemorizationWeight {
	return MemorizationWeight(p.correctlyTranslatedTerms[term] - p.mistakenlyTranslatedTerms[term])
}
