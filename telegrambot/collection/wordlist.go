package collection

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/body"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/superlativeadjectives"
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(pauline.VocabularyEngToRus).
		Update(phrasalverbs.VocabularyEngToRus).
		Update(superlativeadjectives.VocabularyEngToRus).
		Update(body.VocabularyEngToRus)
	VocabularyRusToEng = entity.NewEmptyVocabulary().
		Update(pauline.VocabularyRusToEng).
		Update(phrasalverbs.VocabularyRusToEng).
		Update(superlativeadjectives.VocabularyRusToEng).
		Update(body.VocabularyRusToEng)
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
