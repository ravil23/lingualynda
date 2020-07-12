package entity

import (
	"fmt"
)

type ChatID int64
type ChatMode string
type ChatVocabularyType string

const (
	ChatModeAllTasks = ChatMode("alltasks")
	ChatModeEngToRus = ChatMode("eng2rus")
	ChatModeRusToEng = ChatMode("rus2eng")

	ChatVocabularyTypeAllVocabularies       = ChatVocabularyType("allvocabularies")
	ChatVocabularyTypePauline               = ChatVocabularyType("pauline")
	ChatVocabularyTypePhrasalVerbs          = ChatVocabularyType("phrasalverbs")
	ChatVocabularyTypeSuperlativeAdjectives = ChatVocabularyType("superlativeadjectives")
	ChatVocabularyTypeBody                  = ChatVocabularyType("body")
	ChatVocabularyTypeIdioms                = ChatVocabularyType("idioms")
	ChatVocabularyTypeLesson                = ChatVocabularyType("lesson")

	defaultChatMode           = ChatModeAllTasks
	defaultChatVocabularyType = ChatVocabularyTypeAllVocabularies
)

var AllChatVocabularyTypes = []ChatVocabularyType{
	ChatVocabularyTypeAllVocabularies,
	ChatVocabularyTypePauline,
	ChatVocabularyTypePhrasalVerbs,
	ChatVocabularyTypeSuperlativeAdjectives,
	ChatVocabularyTypeBody,
	ChatVocabularyTypeIdioms,
	ChatVocabularyTypeLesson,
}

type Chat struct {
	id             ChatID
	mode           ChatMode
	vocabularyType ChatVocabularyType
	debug          bool
	vocabularies   []*Vocabulary
}

func NewChat(chatID ChatID) *Chat {
	return &Chat{
		id:             chatID,
		mode:           defaultChatMode,
		vocabularyType: defaultChatVocabularyType,
	}
}

func (c *Chat) Configure(debug bool, mode ChatMode, vocabularyType ChatVocabularyType) {
	c.debug = debug
	c.mode = mode
	c.vocabularyType = vocabularyType
}

func (c *Chat) ConfigureFromText(text string) {
	switch text {
	case "/debug true":
		c.debug = true
	case "/debug false":
		c.debug = false
	case fmt.Sprintf("/%s", ChatModeAllTasks):
		c.mode = ChatModeAllTasks
	case fmt.Sprintf("/%s", ChatModeEngToRus):
		c.mode = ChatModeEngToRus
	case fmt.Sprintf("/%s", ChatModeRusToEng):
		c.mode = ChatModeRusToEng
	case fmt.Sprintf("/%s", ChatVocabularyTypeAllVocabularies):
		c.vocabularyType = ChatVocabularyTypeAllVocabularies
	case fmt.Sprintf("/%s", ChatVocabularyTypePauline):
		c.vocabularyType = ChatVocabularyTypePauline
	case fmt.Sprintf("/%s", ChatVocabularyTypePhrasalVerbs):
		c.vocabularyType = ChatVocabularyTypePhrasalVerbs
	case fmt.Sprintf("/%s", ChatVocabularyTypeSuperlativeAdjectives):
		c.vocabularyType = ChatVocabularyTypeSuperlativeAdjectives
	case fmt.Sprintf("/%s", ChatVocabularyTypeBody):
		c.vocabularyType = ChatVocabularyTypeBody
	case fmt.Sprintf("/%s", ChatVocabularyTypeIdioms):
		c.vocabularyType = ChatVocabularyTypeIdioms
	case fmt.Sprintf("/%s", ChatVocabularyTypeLesson):
		c.vocabularyType = ChatVocabularyTypeLesson
	}
}

func (c *Chat) GetID() ChatID {
	return c.id
}

func (c *Chat) GetMode() ChatMode {
	if c == nil {
		return defaultChatMode
	} else {
		return c.mode
	}
}

func (c *Chat) GetVocabularyType() ChatVocabularyType {
	if c == nil {
		return defaultChatVocabularyType
	} else {
		return c.vocabularyType
	}
}

func (c *Chat) IsDebuggingEnabled() bool {
	return c.debug
}

func (c *Chat) GetListOfVocabularies() []*Vocabulary {
	return c.vocabularies
}

func (c *Chat) SetVocabularies(vocabularies ...*Vocabulary) {
	c.vocabularies = vocabularies
}
