package paulineunit10

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"astronomical":  {"астрономический"},
		"climatic":      {"климатический"},
		"colossal":      {"колоссальный", "огромный", "гигантский", "грандиозный"},
		"fascinating":   {"увлекательный", "очаровательный", "удивительный", "интересный", "обворожительный"},
		"fleeting":      {"мимолетный"},
		"immeasurable":  {"неизмеримый", "неоценимый"},
		"immense":       {"огромный", "колоссальный", "гигантский", "необъятный"},
		"imperceptible": {"незаметный"},
		"infinitesimal": {"бесконечно малый"},
		"magnetic":      {"магнитный", "магнетический"},
		"minuscule":     {"крохотный", "ничтожный"},
		"transient":     {"переходный", "временный", "кратковременный"},
		"vast":          {"огромный", "обширный", "большой", "колоссальный"},
	},
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"condensation": {"коненсация", "конденсат"},
		"debris":       {"мусор", "обломок", "осколок", "развалины"},
		"eclipse":      {"затмение"},
		"evaporation":  {"испарение", "выпаривание"},
		"gravity":      {"гравитация", "тяжесть", "сила тяжести", "притяжение"},
		"harbour":      {"гавань", "порт", "залив"},
		"light year":   {"световой год"},
	},
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"absorb":    {"поглащать", "покрыть", "осваивать", "покрывать", "впитать"},
		"activate":  {"активировать", "включить", "задействовать", "запустить"},
		"burst":     {"лопнуть", "взорваться", "разрушать"},
		"collide":   {"сталкиваться", "столкнуться"},
		"dilute":    {"разбавить", "ослабить"},
		"dissolve":  {"распустить", "расторгнуть", "упразднить"},
		"penetrate": {"проникнуть"},
		"pierce":    {"проколоть", "пронзить", "пробить", "проткнуть", "прокалывать"},
		"reflect":   {"отражать", "свидетельствовать", "подумать", "учитывать"},
		"release":   {"освободить", "выпустить", "отпустить"},
		"repel":     {"отражать", "отталкивать", "противостоять"},
		"solidify":  {"укрепить", "закрепить", "упрочить"},
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
