package paulineunit4

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"absorb":      {"поглощать", "осваивать", "покрывать", "впитать"},
		"contaminate": {"загрязнять", "заражать", "испортить"},
		"release":     {"выпустить", "освободить"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
