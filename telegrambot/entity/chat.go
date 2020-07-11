package entity

import (
	"github.com/ravil23/lingualynda/telegrambot/collection"
	"github.com/ravil23/lingualynda/telegrambot/collection/body"
	"github.com/ravil23/lingualynda/telegrambot/collection/idioms"
	"github.com/ravil23/lingualynda/telegrambot/collection/lesson"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
	"github.com/ravil23/lingualynda/telegrambot/collection/superlativeadjectives"
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
	vocabularies []*schema.Vocabulary
}

func NewChat(chatID ChatID) *Chat {
	c := &Chat{
		id:         chatID,
		mode:       ChatModeRandom,
		vocabulary: ChatVocabularyAll,
	}
	c.refreshVocabularies()
	return c
}

func (c *Chat) GetID() ChatID {
	return c.id
}

func (c *Chat) SetMode(mode ChatMode) {
	c.mode = mode
	c.refreshVocabularies()
}

func (c *Chat) GetMode() ChatMode {
	return c.mode
}

func (c *Chat) SetVocabulary(vocabulary ChatVocabulary) {
	c.vocabulary = vocabulary
	c.refreshVocabularies()
}

func (c *Chat) GetVocabulary() ChatVocabulary {
	return c.vocabulary
}

func (c *Chat) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Chat) IsDebuggingEnabled() bool {
	return c.debug
}

func (c *Chat) GetVocabularies() []*schema.Vocabulary {
	return c.vocabularies
}

func (c *Chat) refreshVocabularies() {
	switch c.vocabulary {
	case ChatVocabularyAll:
		switch c.mode {
		case ChatModeEngToRus:
			c.vocabularies = []*schema.Vocabulary{collection.VocabularyEngToRus}
		case ChatModeRusToEng:
			c.vocabularies = []*schema.Vocabulary{collection.VocabularyRusToEng}
		default:
			c.vocabularies = collection.AllVocabularies
		}
	case ChatVocabularyPauline:
		switch c.mode {
		case ChatModeEngToRus:
			c.vocabularies = []*schema.Vocabulary{pauline.VocabularyEngToRus}
		case ChatModeRusToEng:
			c.vocabularies = []*schema.Vocabulary{pauline.VocabularyRusToEng}
		default:
			c.vocabularies = pauline.AllVocabularies
		}
	case ChatVocabularyPhrasalVerbs:
		switch c.mode {
		case ChatModeEngToRus:
			c.vocabularies = []*schema.Vocabulary{phrasalverbs.VocabularyEngToRus}
		case ChatModeRusToEng:
			c.vocabularies = []*schema.Vocabulary{phrasalverbs.VocabularyRusToEng}
		default:
			c.vocabularies = phrasalverbs.AllVocabularies
		}
	case ChatVocabularySuperlativeAdjectives:
		switch c.mode {
		case ChatModeEngToRus:
			c.vocabularies = []*schema.Vocabulary{superlativeadjectives.VocabularyEngToRus}
		case ChatModeRusToEng:
			c.vocabularies = []*schema.Vocabulary{superlativeadjectives.VocabularyRusToEng}
		default:
			c.vocabularies = superlativeadjectives.AllVocabularies
		}
	case ChatVocabularyBody:
		switch c.mode {
		case ChatModeEngToRus:
			c.vocabularies = []*schema.Vocabulary{body.VocabularyEngToRus}
		case ChatModeRusToEng:
			c.vocabularies = []*schema.Vocabulary{body.VocabularyRusToEng}
		default:
			c.vocabularies = body.AllVocabularies
		}
	case ChatVocabularyIdioms:
		switch c.mode {
		case ChatModeEngToRus:
			c.vocabularies = []*schema.Vocabulary{idioms.VocabularyEngToRus}
		case ChatModeRusToEng:
			c.vocabularies = []*schema.Vocabulary{idioms.VocabularyRusToEng}
		default:
			c.vocabularies = idioms.AllVocabularies
		}
	case ChatVocabularyLesson:
		switch c.mode {
		case ChatModeEngToRus:
			c.vocabularies = []*schema.Vocabulary{lesson.VocabularyEngToRus}
		case ChatModeRusToEng:
			c.vocabularies = []*schema.Vocabulary{lesson.VocabularyRusToEng}
		default:
			c.vocabularies = lesson.AllVocabularies
		}
	default:
		c.vocabularies = collection.AllVocabularies
	}
}
