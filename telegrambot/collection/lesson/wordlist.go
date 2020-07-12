package lesson

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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
		"folk":         {"народный"},
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

var VocabularyAdverbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"in advance":     {"заранее", "заблаговременно", "наперед"},
		"overwhelmingly": {"преимущественно"},
	},
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"betrayal":      {"предательство"},
		"chariot":       {"колесница"},
		"contentment":   {"удовлетворенность"},
		"gossip":        {"сплетни"},
		"guilt":         {"вина"},
		"hide-and-seek": {"прятки"},
		"lack":          {"отсутствие", "нехватка", "недостаток"},
		"leather":       {"кожанный материал"},
		"manual labour": {"ручной труд"},
		"mate":          {"приятель"},
		"misery":        {"страдание", "горе", "беда", "мучение"},
		"nausea":        {"тошнота"},
		"proving":       {"доказательство"},
		"slave":         {"раб"},
		"stroll":        {"прогулка"},
		"tears":         {"слезы"},
		"violence":      {"насилие", "жестокость"},
		"well":          {"колодец"},
	},
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyAdverbs).
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
