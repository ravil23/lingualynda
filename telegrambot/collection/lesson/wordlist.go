package lesson

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"amused":       {"развеселенный"},
		"arrogant":     {"высокомерный"},
		"awkward":      {"неуклюжий"},
		"bored":        {"скучающий"},
		"cheerful":     {"жизнерадостный"},
		"confident":    {"уверенный"},
		"curious":      {"любопытный"},
		"depressed":    {"депрессивный"},
		"dishonest":    {"бесчестный"},
		"dissatisfied": {"недовольный"},
		"edible":       {"съедобный"},
		"enthusiastic": {"восторженный"},
		"foreign":      {"иностранный"},
		"frustrated":   {"расстроенный", "разочарованный"},
		"greasy":       {"жирный"},
		"grumpy":       {"сердитый", "ворчливый"},
		"guilty":       {"виноватый"},
		"hazardous":    {"опасный"},
		"impatient":    {"нетерпеливый"},
		"insecure":     {"неуверенный"},
		"lonely":       {"одинокий"},
		"marvellous":   {"чудесный", "удивительный", "изумительный"},
		"needy":        {"неполноценный"},
		"panicky":      {"паникующий"},
		"rebellious":   {"бунтарящий"},
		"resilient":    {"жизнерадостный"},
		"upset":        {"встревоженный", "расстроенный"},
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
		"lack":          {"отсутствие", "нехватка", "недостаток"},
		"leather":       {"кожанный материал"},
		"manual labour": {"ручной труд"},
		"mate":          {"приятель"},
		"misery":        {"страдание", "горе", "беда", "мучение"},
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
