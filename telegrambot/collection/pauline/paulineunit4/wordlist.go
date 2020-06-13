package paulineunit4

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"bacterial":     {"бактериальный"},
		"contagious":    {"заразный", "инфекционный"},
		"contaminated":  {"загрязненный"},
		"controversial": {"спорный", "противоречивый", "неоднозначный", "скандальный"},
		"crucial":       {"важнейший", "решающий", "ключевой", "критический"},
		"essential":     {"необходимый", "основной"},
		"natural":       {"природный", "естественный", "натуральный"},
		"organic":       {"органический", "натуральный"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"capsule":        {"капсула", "коробочка", "оболочка"},
		"carbon dioxide": {"углекислый газ"},
		"compound":       {"комплекс", "соединение", "смесь", "вещество"},
		"consequences":   {"последствия"},
		"dose":           {"дозировка"},
		"emissions":      {"выбросы"},
		"exposure":       {"воздействие", "подверженность"},
		"food chain":     {"пищевая цепь"},
		"infection":      {"инфекция", "заражение", "заболевание"},
		"interaction":    {"взаимодействие", "сотрудничество"},
		"pesticide":      {"пестицид", "вещество против вредителей"},
		"petrochemicals": {"нефтехимия", "нефтепродукты"},
		"protein":        {"белок", "протеин"},
		"reaction":       {"реакция", "реагирование", "отклик"},
		"side effects":   {"побочные эффекты"},
	},
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"absorb":      {"поглощать", "осваивать", "покрывать", "впитать"},
		"contaminate": {"загрязнять", "заражать", "испортить"},
		"release":     {"выпустить", "освободить"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
