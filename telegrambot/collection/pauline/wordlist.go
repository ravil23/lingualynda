package pauline

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit1"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit2"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit3"
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyTotal *schema.Vocabulary
var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(paulineunit1.VocabularyEngToRus).
		Update(paulineunit2.VocabularyEngToRus).
		Update(paulineunit3.VocabularyEngToRus)
	VocabularyRusToEng = schema.NewEmptyVocabulary().
		Update(paulineunit1.VocabularyRusToEng).
		Update(paulineunit2.VocabularyRusToEng).
		Update(paulineunit3.VocabularyRusToEng)
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(paulineunit1.VocabularyTotal).
		Update(paulineunit2.VocabularyTotal).
		Update(paulineunit3.VocabularyTotal)
}
