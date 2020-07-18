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
		"disastrous":     {"катастрофический", "разрушительный"},
		"dishonest":      {"бесчестный"},
		"dissatisfied":   {"недовольный"},
		"edible":         {"съедобный"},
		"enthusiastic":   {"восторженный"},
		"erratic":        {"неустойчивый", "беспорядочный"},
		"folk":           {"народный"},
		"foreign":        {"иностранный"},
		"frustrated":     {"расстроенный", "разочарованный"},
		"further":        {"дальнейший", "дополнительный"},
		"greasy":         {"жирный"},
		"grumpy":         {"сердитый", "ворчливый"},
		"guilty":         {"виноватый"},
		"hazardous":      {"опасный"},
		"hostile":        {"враждебный", "вражеский"},
		"impatient":      {"нетерпеливый"},
		"insecure":       {"неуверенный"},
		"lonely":         {"одинокий"},
		"marvellous":     {"чудесный", "удивительный", "изумительный"},
		"monotonous":     {"монотонный", "однообразный"},
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
		"plenty of":      {"полно", "предостаточно"},
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
		"chores":        {"работа по дому"},
		"companionship": {"дружеские отношения"},
		"confluent":     {"приток реки"},
		"contentment":   {"удовлетворенность"},
		"drudgery":      {"нудная работа", "рутинная работа"},
		"evidence":      {"улика", "подтверждение"},
		"frontier":      {"государственная граница"},
		"gossip":        {"сплетни"},
		"guilt":         {"вина"},
		"hide-and-seek": {"прятки"},
		"labour":        {"труд"},
		"lack":          {"отсутствие", "нехватка", "недостаток"},
		"leather":       {"кожанный материал"},
		"loneliness":    {"одиночество"},
		"manual labour": {"ручной труд"},
		"mate":          {"приятель"},
		"misery":        {"страдание", "горе", "беда", "мучение"},
		"mortgage":      {"ипотека"},
		"mother-in-law": {"свекровь", "теща"},
		"nausea":        {"тошнота"},
		"outcast":       {"изгой", "изгнанник"},
		"overcrowding":  {"перенаселенность"},
		"pathway":       {"путь", "специализация"},
		"pleasure":      {"удовольствие", "радость", "наслаждение", "удовлетворение"},
		"pollution":     {"загрязнение"},
		"proving":       {"доказательство"},
		"sentiment":     {"настроение", "чувство", "мнение"},
		"slave":         {"раб"},
		"solitary":      {"отшельник"},
		"stroll":        {"прогулка"},
		"tears":         {"слезы"},
		"trainee":       {"стажер", "практикант"},
		"violence":      {"насилие", "жестокость"},
		"well":          {"колодец"},
	},
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"be in tears":   {"рыдать"},
		"betray":        {"предать"},
		"bury":          {"хоронить"},
		"cope":          {"справляться", "совладать", "бороться"},
		"cost fortune":  {"стоит целое состояние"},
		"devote":        {"посвятить", "уделять"},
		"encounter":     {"сталкиваться", "встретиться"},
		"gain":          {"получать", "добиться", "завоевать", "приобрести"},
		"get stuck":     {"застрять"},
		"give a notice": {"предупреждать"},
		"have a row":    {"ссориться"},
		"nap":           {"вздремнуть"},
		"perceive":      {"воспринимать", "ощущать"},
		"resemble":      {"быть похожим"},
		"settle in":     {"обосновываться", "обживаться", "обустраиваться"},
		"suffer":        {"страдать", "мучиться"},
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
