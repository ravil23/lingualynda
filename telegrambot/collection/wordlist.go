package collection

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyTotal *schema.Vocabulary
var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(pauline.VocabularyEngToRus).
		Update(phrasalverbs.VocabularyEngToRus)
	VocabularyRusToEng = schema.NewEmptyVocabulary().
		Update(pauline.VocabularyRusToEng).
		Update(phrasalverbs.VocabularyRusToEng)
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(pauline.VocabularyTotal).
		Update(phrasalverbs.VocabularyTotal)
}
