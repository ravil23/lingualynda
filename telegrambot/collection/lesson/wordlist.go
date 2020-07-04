package lesson

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"arrogant":   {"высокомерный"},
		"cheerful":   {"жизнерадостный"},
		"foreign":    {"иностанный"},
		"greasy":     {"жирный"},
		"hazardous":  {"опасный"},
		"needy":      {"неполноценный"},
		"rebellious": {"бунтарь"},
		"resilient":  {"жизнерадостный"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"betrayal":      {"предательство"},
		"chariot":       {"колесница"},
		"gossip":        {"сплетни"},
		"hide-and-seek": {"прятки"},
		"leather":       {"кожанный материал"},
		"manual labour": {"ручной труд"},
		"mate":          {"приятель"},
		"nausea":        {"тошнота"},
		"slave":         {"раб"},
		"stroll":        {"прогулка"},
		"well":          {"колодец"},
	},
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"betray":       {"предать"},
		"bury":         {"хоронить"},
		"cost fortune": {"стоит целое состояние"},
		"devote":       {"посвятить", "уделять"},
		"get stuck":    {"застрять"},
		"resemble":     {"быть похожим"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary
var AllVocabularies []*schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*schema.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
