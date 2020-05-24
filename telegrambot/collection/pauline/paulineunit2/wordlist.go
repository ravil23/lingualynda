package paulineunit2

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"abrupt":        {"резкий", "внезапный", "крутой"},
		"bygone":        {"ушедшее в прошлое"},
		"fleeting":      {"мимолетный"},
		"immense":       {"огромный", "колоссальный", "гигантский"},
		"infinitesimal": {"бесконечно малый"},
		"moderate":      {"умеренный", "средний", "небольшой"},
		"periodic":      {"периодический"},
		"preceding":     {"предыдущий", "предшествующий"},
		"prehistoric":   {"доисторический"},
		"profound":      {"глубокий", "мдрый"},
		"rapid":         {"быстрый", "оперативный", "стремительный", "ускоренный", "резкий", "скоростной"},
		"topical":       {"актуальный", "тематический"},
		"tremendous":    {"огромный", "колоссальный", "потрясающий", "громадный"},
		"turbulent":     {"бурный", "неспокойный", "турбулентный"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"era":        {"эра", "эпоха", "период", "этап"},
		"evolution":  {"эволюция", "развитие", "изменение", "динамика", "преобразование"},
		"finds":      {"находки"},
		"pioneer":    {"пионер", "первопроходец", "инициатор", "новатор", "первооткрыватель"},
		"remnants":   {"остатки"},
		"retrospect": {"взгляд в прошлое"},
		"transition": {"переход", "преобразование"},
		"trend":      {"тенденция", "тренд", "динамика", "направление", "веяние"},
	},
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"excavate": {"выкопать", "перекопать", "раскапывать", "производить земляные работы"},
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
