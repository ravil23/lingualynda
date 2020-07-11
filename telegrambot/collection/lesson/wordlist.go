package lesson

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"arrogant":   {"высокомерный"},
		"cheerful":   {"жизнерадостный"},
		"dishonest":  {"бесчестный"},
		"edible":     {"съедобный"},
		"grumpy":     {"сердитый", "ворчливый"},
		"guilty":     {"виноватый"},
		"foreign":    {"иностанный"},
		"greasy":     {"жирный"},
		"hazardous":  {"опасный"},
		"needy":      {"неполноценный"},
		"rebellious": {"бунтарь"},
		"resilient":  {"жизнерадостный"},
	},
)

var VocabularyAdverbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"in advance": {"заранее", "заблаговременно", "наперед"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"betrayal":      {"предательство"},
		"chariot":       {"колесница"},
		"gossip":        {"сплетни"},
		"guilt":         {"вина"},
		"hide-and-seek": {"прятки"},
		"leather":       {"кожанный материал"},
		"manual labour": {"ручной труд"},
		"mate":          {"приятель"},
		"nausea":        {"тошнота"},
		"slave":         {"раб"},
		"stroll":        {"прогулка"},
		"tears":         {"слезы"},
		"well":          {"колодец"},
	},
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"be in tears":  {"рыдать"},
		"betray":       {"предать"},
		"bury":         {"хоронить"},
		"cost fortune": {"стоит целое состояние"},
		"devote":       {"посвятить", "уделять"},
		"get stuck":    {"застрять"},
		"have a row":   {"ссориться"},
		"resemble":     {"быть похожим"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary
var AllVocabularies []*schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyAdverbs).
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*schema.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
