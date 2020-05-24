package collection

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyTotal *schema.Vocabulary

func init() {
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(pauline.VocabularyTotal).
		Update(phrasalverbs.VocabularyTotal)
}
