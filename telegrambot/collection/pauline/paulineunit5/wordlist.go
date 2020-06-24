package paulineunit5

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"analytical":   {"аналитический"},
		"blue-collar":  {"рабочий класс", "работники физического труда"},
		"conceptual":   {"концептуальный"},
		"hypothetical": {"гипотетический"},
		"indicative":   {"ориентировочный", "примерный", "показательный", "индикативный"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"accounting":        {"учет", "отчетность", "бухгалтерия"},
		"apprenticeship":    {"ученичество", "обучение", "стажировки"},
		"conceptualisation": {"концептуализация"},
		"consistency":       {"последовательность", "согласованность", "соответствие", "единообразие", "совместимость"},
		"establishment":     {"создание", "учреждение", "установление", "разработка", "формирование", "заведение", "организация", "предприятие"},
		"formulation":       {"разработка", "формулировка", "постановка", "подготовка"},
		"hypothesis":        {"гипотеза", "теория", "предположение", "вариант"},
		"inconsistency":     {"несоответствие", "непоследовательность", "несогласованность", "противоречие"},
		"insignificance":    {"незначительность"},
		"interpretation":    {"толкование", "интерпретация", "трактовка", "объяснение"},
		"technician":        {"техник", "лаборант", "специалист"},
		"theorist":          {"теоретик"},
		"vocation":          {"призвание", "предназначение"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
