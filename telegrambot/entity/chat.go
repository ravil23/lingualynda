package entity

import (
	"fmt"
)

type ChatID int64
type ChatMode string
type ChatVocabulary string

const (
	ChatModeRandom   = ChatMode("random")
	ChatModeEngToRus = ChatMode("eng2rus")
	ChatModeRusToEng = ChatMode("rus2eng")

	ChatVocabularyAll                   = ChatVocabulary("all")
	ChatVocabularyPauline               = ChatVocabulary("pauline")
	ChatVocabularyPhrasalVerbs          = ChatVocabulary("phrasalverbs")
	ChatVocabularySuperlativeAdjectives = ChatVocabulary("superlativeadjectives")
	ChatVocabularyBody                  = ChatVocabulary("body")
	ChatVocabularyIdioms                = ChatVocabulary("idioms")
	ChatVocabularyLesson                = ChatVocabulary("lesson")
)

type Chat struct {
	id           ChatID
	mode         ChatMode
	vocabulary   ChatVocabulary
	debug        bool
	vocabularies []*Vocabulary
}

func NewChat(chatID ChatID) *Chat {
	return &Chat{
		id:         chatID,
		mode:       ChatModeRandom,
		vocabulary: ChatVocabularyAll,
	}
}

func (c *Chat) Configure(text string) {
	switch text {
	case "/debug true":
		c.debug = true
	case "/debug false":
		c.debug = false
	case fmt.Sprintf("/%s", ChatModeRandom):
		c.mode = ChatModeRandom
	case fmt.Sprintf("/%s", ChatModeEngToRus):
		c.mode = ChatModeEngToRus
	case fmt.Sprintf("/%s", ChatModeRusToEng):
		c.mode = ChatModeRusToEng
	case fmt.Sprintf("/%s", ChatVocabularyAll):
		c.vocabulary = ChatVocabularyAll
	case fmt.Sprintf("/%s", ChatVocabularyPauline):
		c.vocabulary = ChatVocabularyPauline
	case fmt.Sprintf("/%s", ChatVocabularyPhrasalVerbs):
		c.vocabulary = ChatVocabularyPhrasalVerbs
	case fmt.Sprintf("/%s", ChatVocabularySuperlativeAdjectives):
		c.vocabulary = ChatVocabularySuperlativeAdjectives
	case fmt.Sprintf("/%s", ChatVocabularyBody):
		c.vocabulary = ChatVocabularyBody
	case fmt.Sprintf("/%s", ChatVocabularyIdioms):
		c.vocabulary = ChatVocabularyIdioms
	case fmt.Sprintf("/%s", ChatVocabularyLesson):
		c.vocabulary = ChatVocabularyLesson
	}
}

func (c *Chat) GetID() ChatID {
	return c.id
}

func (c *Chat) GetMode() ChatMode {
	return c.mode
}

func (c *Chat) GetVocabulary() ChatVocabulary {
	return c.vocabulary
}

func (c *Chat) IsDebuggingEnabled() bool {
	return c.debug
}

func (c *Chat) GetVocabularies() []*Vocabulary {
	return c.vocabularies
}

func (c *Chat) SetVocabularies(vocabularies ...*Vocabulary) {
	c.vocabularies = vocabularies
}
