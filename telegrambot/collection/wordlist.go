package collection

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/paulineunit1"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyTotal *schema.Vocabulary

func init() {
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(paulineunit1.VocabularyTotal)
}
