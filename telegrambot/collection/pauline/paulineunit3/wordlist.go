package paulineunit3

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"anti-social":    {"антиобщественный"},
		"conservative":   {"консервативный"},
		"conventional":   {"обычный", "традиционный", "договорной", "общепринятый"},
		"harmonious":     {"гармоничный"},
		"influential":    {"влиятельный"},
		"multicultural":  {"многокультурный"},
		"suburban":       {"пригородный"},
		"unconventional": {"нетрадиционный", "необычный", "нестандартный", "нешаблонный"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"behaviour":    {"поведение", "действия", "выходки"},
		"demographics": {"демография"},
		"heritage":     {"наследие", "достояние", "происхождение", "корни", "наследство"},
		"interaction":  {"взаимодействие", "сотрудничество", "общение", "взаимоотношения", "взаимосвязи"},
		"minority":     {"меньшинство", "несовершеннолетие"},
		"norm":         {"норма", "правило", "норматив"},
		"peer":         {"сверстник", "ровня"},
		"pressure":     {"давление", "нагрузка", "воздействие", "напряжение", "принуждение"},
		"segment":      {"сегмент", "часть", "совещание", "участок", "раздел"},
		"standard":     {"стандарт", "уровень", "норма", "критерий", "эталон"},
		"status":       {"статус", "положение", "состояние", "ситуация"},
	},
)

var VocabularyPhrases = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"follow accepted behaviour": {"следовать принятому поведению"},
		"shun mainstream values":    {"избегать массовых ценностей"},
	},
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"conform": {"соответствовать", "подчиняться"},
		"exclude": {"исключать", "не допускать", "лишить"},
		"flaunt":  {"выставлять напоказ", "щеголять", "красоваться", "афишировать", "хвастаться"},
		"obey":    {"подчиняться", "повиноваться", "слушаться", "соблюдать", "следовать", "выполнять"},
		"skew":    {"исказить", "наклонить"},
	},
)

var VocabularyTotal *schema.Vocabulary

func init() {
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns).
		Update(VocabularyPhrases).
		Update(VocabularyVerbs)
}
