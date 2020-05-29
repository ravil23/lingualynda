package telegram

import (
	"github.com/ravil23/lingualynda/telegrambot/collection"
	"github.com/ravil23/lingualynda/telegrambot/collection/adjectives"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
	"github.com/ravil23/lingualynda/telegrambot/dao"
)

type mode string
type vocabulary string

const (
	modeRandom   = mode("random")
	modeEngToRus = mode("eng2rus")
	modeRusToEng = mode("rus2eng")

	vocabularyAll          = vocabulary("all")
	vocabularyPauline      = vocabulary("pauline")
	vocabularyPhrasalVerbs = vocabulary("phrasalverbs")
	vocabularyAdjectives   = vocabulary("adjectives")
)

var chatsStates = map[dao.ChatID]*ChatsState{}

type ChatsState struct {
	mode         mode
	vocabulary   vocabulary
	debug        bool
	vocabularies []*schema.Vocabulary
}

func (s *ChatsState) SetMode(mode mode) {
	s.mode = mode
	s.refreshVocabularies()
}

func (s *ChatsState) SetVocabulary(vocabulary vocabulary) {
	s.vocabulary = vocabulary
	s.refreshVocabularies()
}

func (s *ChatsState) refreshVocabularies() {
	switch s.vocabulary {
	case vocabularyAll:
		switch s.mode {
		case modeEngToRus:
			s.vocabularies = []*schema.Vocabulary{collection.VocabularyEngToRus}
		case modeRusToEng:
			s.vocabularies = []*schema.Vocabulary{collection.VocabularyRusToEng}
		default:
			s.vocabularies = collection.AllVocabularies
		}
	case vocabularyPauline:
		switch s.mode {
		case modeEngToRus:
			s.vocabularies = []*schema.Vocabulary{pauline.VocabularyEngToRus}
		case modeRusToEng:
			s.vocabularies = []*schema.Vocabulary{pauline.VocabularyRusToEng}
		default:
			s.vocabularies = pauline.AllVocabularies
		}
	case vocabularyPhrasalVerbs:
		switch s.mode {
		case modeEngToRus:
			s.vocabularies = []*schema.Vocabulary{phrasalverbs.VocabularyEngToRus}
		case modeRusToEng:
			s.vocabularies = []*schema.Vocabulary{phrasalverbs.VocabularyRusToEng}
		default:
			s.vocabularies = phrasalverbs.AllVocabularies
		}
	case vocabularyAdjectives:
		switch s.mode {
		case modeEngToRus:
			s.vocabularies = []*schema.Vocabulary{adjectives.VocabularyEngToRus}
		case modeRusToEng:
			s.vocabularies = []*schema.Vocabulary{adjectives.VocabularyRusToEng}
		default:
			s.vocabularies = adjectives.AllVocabularies
		}
	default:
		s.vocabularies = collection.AllVocabularies
	}
}
