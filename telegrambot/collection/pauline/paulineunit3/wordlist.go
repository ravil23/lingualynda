package paulineunit3

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyPhrases = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"follow accepted behaviour": {"следовать принятому поведению"},
		"shun mainstream values":    {"избегать массовых ценностей"},
	},
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"conform": {"соответствовать", "подчиняться"},
		"exclude": {"исключать", "не допускать", "лишить"},
		"flaunt":  {"выставлять напоказ", "щеголять", "красоваться", "афишировать", "хвастаться"},
		"obey":    {"подчиняться", "повиноваться", "слушаться", "соблюдать", "следовать", "выполнять"},
		"skew":    {"исказить", "наклонить"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns).
		Update(VocabularyPhrases).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
