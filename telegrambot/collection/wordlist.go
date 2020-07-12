package collection

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/body"
	"github.com/ravil23/lingualynda/telegrambot/collection/idioms"
	"github.com/ravil23/lingualynda/telegrambot/collection/lesson"
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
		Update(body.VocabularyEngToRus).
		Update(idioms.VocabularyEngToRus).
		Update(lesson.VocabularyEngToRus).
		Update(pauline.VocabularyEngToRus).
		Update(phrasalverbs.VocabularyEngToRus).
		Update(superlativeadjectives.VocabularyEngToRus)
	VocabularyRusToEng = entity.NewEmptyVocabulary().
		Update(body.VocabularyRusToEng).
		Update(idioms.VocabularyRusToEng).
		Update(lesson.VocabularyRusToEng).
		Update(pauline.VocabularyRusToEng).
		Update(phrasalverbs.VocabularyRusToEng).
		Update(superlativeadjectives.VocabularyRusToEng)
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
