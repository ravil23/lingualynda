package paulineunit7

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"breathtaking": {"захватывающий"},
		"charming":     {"очаровательный", "обаятельный"},
		"comfortable":  {"удобный", "комфортный", "уютный"},
		"dramatic":     {"резкий", "значительный"},
		"magnificent":  {"великолепный", "роскошный"},
		"memorable":    {"памятный", "незабываемый"},
		"mountainous":  {"горный"},
		"precarious":   {"ненадежный", "нестабильный"},
		"remote":       {"отдаленный", "дистанционный"},
		"rustic":       {"деревенский", "сельский"},
		"spectacular":  {"зрелищный", "впечатляющий", "эффектный"},
		"stunning":     {"сногсшибательный", "великолепный", "ошеломительный"},
	},
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"attract":  {"привлечь"},
		"damage":   {"повредить", "подорвать"},
		"outweigh": {"перевесить"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
