package pauline

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit1"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit2"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit3"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyTotal *schema.Vocabulary

func init() {
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(paulineunit1.VocabularyTotal).
		Update(paulineunit2.VocabularyTotal).
		Update(paulineunit3.VocabularyTotal)
}
