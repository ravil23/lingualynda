package paulineunit2

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"abrupt":        {"резкий", "внезапный", "крутой"},
		"bygone":        {"ушедшее в прошлое"},
		"fleeting":      {"мимолетный"},
		"immense":       {"огромный", "колоссальный", "гигантский"},
		"infinitesimal": {"бесконечно малый"},
		"moderate":      {"умеренный", "средний", "небольшой"},
		"periodic":      {"периодический"},
		"preceding":     {"предыдущий", "предшествующий"},
		"prehistoric":   {"доисторический"},
		"profound":      {"глубокий", "мудрый"},
		"rapid":         {"быстрый", "оперативный", "стремительный", "ускоренный", "резкий", "скоростной"},
		"topical":       {"актуальный", "тематический"},
		"tremendous":    {"огромный", "колоссальный", "потрясающий", "громадный"},
		"turbulent":     {"бурный", "неспокойный", "турбулентный"},
	},
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"excavate": {"выкопать", "перекопать", "раскапывать", "производить земляные работы"},
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
