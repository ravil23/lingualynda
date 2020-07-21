package lesson

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"aforementioned": {"вышеупомянутый", "указанный"},
		"amused":         {"развеселенный"},
		"appropriate":    {"уместный", "подходящий"},
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
		"engaged":        {"помолвлены", "обручены"},
		"enthusiastic":   {"восторженный"},
		"erratic":        {"неустойчивый", "беспорядочный"},
		"folk":           {"народный"},
		"foreign":        {"иностранный"},
		"frustrated":     {"расстроенный", "разочарованный"},
		"further":        {"дальнейший", "дополнительный"},
		"grateful":       {"благодарный", "признательный"},
		"greasy":         {"жирный"},
		"grumpy":         {"сердитый", "ворчливый"},
		"guilty":         {"виноватый"},
		"handheld":       {"портативный", "карманный"},
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
		"loudly":         {"громко", "погромче", "громогласно"},
		"plenty of":      {"полно", "предостаточно"},
		"overwhelmingly": {"преимущественно"},
	},
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"absence":       {"отсутствие", "недостаток"},
		"admission":     {"входная плата"},
		"anxiety":       {"тревожность", "тревога"},
		"appointment":   {"запись", "прием"},
		"benefit":       {"польза", "преимущество"},
		"betrayal":      {"предательство"},
		"bowl":          {"супница"},
		"chariot":       {"колесница"},
		"chores":        {"работа по дому"},
		"coal mine":     {"угольная шахта"},
		"companionship": {"дружеские отношения"},
		"confluent":     {"приток реки"},
		"contentment":   {"удовлетворенность"},
		"cup":           {"чайная чашка"},
		"dish":          {"блюдо"},
		"drudgery":      {"нудная работа", "рутинная работа"},
		"evidence":      {"улика", "подтверждение"},
		"fork":          {"вилка"},
		"frontier":      {"государственная граница"},
		"glass":         {"стакан"},
		"gossip":        {"сплетни"},
		"guilt":         {"вина"},
		"hide-and-seek": {"прятки"},
		"knife":         {"нож"},
		"labour":        {"труд"},
		"lack":          {"отсутствие", "нехватка", "недостаток"},
		"landline":      {"стационарный телефон"},
		"leather":       {"кожанный материал"},
		"loneliness":    {"одиночество"},
		"manual labour": {"ручной труд"},
		"mate":          {"приятель"},
		"matter":        {"вопрос", "дело", "предмет"},
		"misery":        {"страдание", "горе", "беда", "мучение"},
		"mishap":        {"несчастный случай", "неудача", "происшествие", "инцидент", "неприятность", "несчастье", "недоразумение"},
		"mortgage":      {"ипотека"},
		"mother-in-law": {"свекровь", "теща"},
		"mug":           {"кружка с ручкой большая"},
		"nausea":        {"тошнота"},
		"outcast":       {"изгой", "изгнанник"},
		"overcrowding":  {"перенаселенность"},
		"pathway":       {"путь", "специализация"},
		"plate":         {"тарелка"},
		"pleasure":      {"удовольствие", "радость", "наслаждение", "удовлетворение"},
		"pollution":     {"загрязнение"},
		"proving":       {"доказательство"},
		"sentiment":     {"настроение", "чувство", "мнение"},
		"slave":         {"раб"},
		"snoring":       {"храп"},
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
		"bang":          {"стукнуть", "ударить"},
		"be in tears":   {"рыдать"},
		"betray":        {"предать"},
		"borrow":        {"одолжить", "позаимствовать"},
		"bury":          {"хоронить"},
		"convince":      {"убеждать", "уговаривать"},
		"cope":          {"справляться", "совладать", "бороться"},
		"cost fortune":  {"стоит целое состояние"},
		"devote":        {"посвятить", "уделять"},
		"encounter":     {"сталкиваться", "встретиться"},
		"gain":          {"получать", "добиться", "завоевать", "приобрести"},
		"get confused":  {"запутаться", "перепутать"},
		"get lost":      {"заблудиться", "потеряться"},
		"get stuck":     {"застрять"},
		"give a notice": {"предупреждать"},
		"have a row":    {"ссориться"},
		"nap":           {"вздремнуть"},
		"perceive":      {"воспринимать", "ощущать"},
		"resemble":      {"быть похожим"},
		"settle in":     {"обосновываться", "обживаться", "обустраиваться"},
		"slip":          {"проскользить", "проскочить"},
		"spill":         {"пролить", "разлить"},
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
