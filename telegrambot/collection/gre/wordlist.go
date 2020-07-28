package gre

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var Vocabulary = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"abash":       {"смущать", "конфузить"},
		"abate":       {"уменьшать", "ослаблять"},
		"abdicate":    {"отрекаться (от власти)"},
		"aberrant":    {"сбившийся с пути", "ненормальный"},
		"abet":        {"помогать", "поощрять", "содействовать"},
		"abeyance":    {"состояние неопределенности"},
		"abhor":       {"ненавидеть", "презирать"},
		"abject":      {"подлый", "низкий", "ужасный"},
		"abjure":      {"отказываться", "отрекаться"},
		"abnegation":  {"отрицание", "непринятие"},
		"abomination": {"неприязнь", "враждебность"},
		"abortive":    {"бесплодный", "бесполезный"},
		"abound":      {"изобиловать"},
		"aboveboard":  {"открытый", "честный", "прямой"},
		"abridge":     {"сокращать", "уменьшать", "снижать"},
		"abrogate":    {"отменять", "аннулировать"},
		"abrupt":      {"внезапный", "обрывистый", "грубый"},
		"abscond":     {"избегать", "уходить от"},
		"absolve":     {"освобождать", "прощать"},
		"abstemious":  {"умеренный", "скромный"},
		"abstruse":    {"трудный для понимания", "глубокий"},
		"abut":        {"примыкать", "граничить"},
		"abysmal":     {"низкий", "отвратительный"},
		"acclaim":     {"приветствовать", "объявлять"},
		"accolade":    {"похвала", "одобрение"},
		"accomplice":  {"сообщник", "соучастник"},
		"accost":      {"приставать", "обращаться"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(Vocabulary)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
