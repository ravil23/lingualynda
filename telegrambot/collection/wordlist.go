package collection

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/body"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
	"github.com/ravil23/lingualynda/telegrambot/collection/superlativeadjectives"
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary
var AllVocabularies []*schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(pauline.VocabularyEngToRus).
		Update(phrasalverbs.VocabularyEngToRus).
		Update(superlativeadjectives.VocabularyEngToRus).
		Update(body.VocabularyEngToRus)
	VocabularyRusToEng = schema.NewEmptyVocabulary().
		Update(pauline.VocabularyRusToEng).
		Update(phrasalverbs.VocabularyRusToEng).
		Update(superlativeadjectives.VocabularyRusToEng).
		Update(body.VocabularyRusToEng)
	AllVocabularies = []*schema.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
