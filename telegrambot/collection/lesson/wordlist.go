package lesson

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"aforementioned": {"вышеупомянутый", "указанный"},
		"amused":         {"развеселенный"},
		"arrogant":       {"высокомерный"},
		"awkward":        {"неуклюжий"},
		"bored":          {"скучающий"},
		"cheerful":       {"жизнерадостный"},
		"confident":      {"уверенный"},
		"curious":        {"любопытный"},
		"depressed":      {"депрессивный"},
		"dishonest":      {"бесчестный"},
		"dissatisfied":   {"недовольный"},
		"edible":         {"съедобный"},
		"enthusiastic":   {"восторженный"},
		"folk":           {"народный"},
		"foreign":        {"иностранный"},
		"frustrated":     {"расстроенный", "разочарованный"},
		"greasy":         {"жирный"},
		"grumpy":         {"сердитый", "ворчливый"},
		"guilty":         {"виноватый"},
		"hazardous":      {"опасный"},
		"hostile":        {"враждебный", "вражеский"},
		"impatient":      {"нетерпеливый"},
		"insecure":       {"неуверенный"},
		"lonely":         {"одинокий"},
		"marvellous":     {"чудесный", "удивительный", "изумительный"},
		"needy":          {"неполноценный"},
		"panicky":        {"паникующий"},
		"rebellious":     {"бунтарящий"},
		"resilient":      {"жизнерадостный"},
		"upset":          {"встревоженный", "расстроенный"},
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
		"admission":     {"входная плата"},
		"anxiety":       {"тревожность", "тревога"},
		"benefit":       {"польза", "преимущество"},
		"betrayal":      {"предательство"},
		"chariot":       {"колесница"},
		"companionship": {"дружеские отношения"},
		"confluent":     {"приток реки"},
		"contentment":   {"удовлетворенность"},
		"gossip":        {"сплетни"},
		"guilt":         {"вина"},
		"hide-and-seek": {"прятки"},
		"lack":          {"отсутствие", "нехватка", "недостаток"},
		"leather":       {"кожанный материал"},
		"loneliness":    {"одиночество"},
		"manual labour": {"ручной труд"},
		"mate":          {"приятель"},
		"misery":        {"страдание", "горе", "беда", "мучение"},
		"mortgage":      {"ипотека"},
		"nausea":        {"тошнота"},
		"outcast":       {"изгой", "изгнанник"},
		"overcrowding":  {"перенаселенность"},
		"pleasure":      {"удовольствие", "радость", "наслаждение", "удовлетворение"},
		"pollution":     {"загрязнение"},
		"proving":       {"доказательство"},
		"sentiment":     {"настроение", "чувство", "мнение"},
		"slave":         {"раб"},
		"stroll":        {"прогулка"},
		"tears":         {"слезы"},
		"trainee":       {"стажер", "практикант"},
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
		"encounter":    {"сталкиваться", "встретиться"},
		"devote":       {"посвятить", "уделять"},
		"gain":         {"получать", "добиться", "завоевать", "приобрести"},
		"get stuck":    {"застрять"},
		"have a row":   {"ссориться"},
		"perceive":     {"воспринимать", "ощущать"},
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
