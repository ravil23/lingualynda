package collection

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/paulineunit1"
	"github.com/ravil23/lingualynda/telegrambot/collection/paulineunit2"
	"github.com/ravil23/lingualynda/telegrambot/collection/paulineunit3"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyTotal *schema.Vocabulary

func init() {
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(phrasalverbs.VocabularyTotal).
		Update(paulineunit1.VocabularyTotal).
		Update(paulineunit2.VocabularyTotal).
		Update(paulineunit3.VocabularyTotal)
}
