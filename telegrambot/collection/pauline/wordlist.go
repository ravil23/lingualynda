package pauline

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit1"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit2"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit3"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit4"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit5"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline/paulineunit6"
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(paulineunit1.VocabularyEngToRus).
		Update(paulineunit2.VocabularyEngToRus).
		Update(paulineunit3.VocabularyEngToRus).
		Update(paulineunit4.VocabularyEngToRus).
		Update(paulineunit5.VocabularyEngToRus).
		Update(paulineunit6.VocabularyEngToRus)
	VocabularyRusToEng = entity.NewEmptyVocabulary().
		Update(paulineunit1.VocabularyRusToEng).
		Update(paulineunit2.VocabularyRusToEng).
		Update(paulineunit3.VocabularyRusToEng).
		Update(paulineunit4.VocabularyRusToEng).
		Update(paulineunit5.VocabularyRusToEng).
		Update(paulineunit6.VocabularyRusToEng)
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
