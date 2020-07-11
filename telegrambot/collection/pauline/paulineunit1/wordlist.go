package paulineunit1

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"apprehensive":        {"опасающийся", "боязливый"},
		"assertive":           {"напористый", "настойчивый"},
		"clumsy":              {"неуклюжий", "неловкий"},
		"cynical":             {"циничный"},
		"desirable":           {"желательный", "целесообразный"},
		"eccentric":           {"эксцентричный", "чудаковатый"},
		"egotistical":         {"эгоистичный", "самовлюбленный"},
		"gullible":            {"доверчивый", "наивный", "легковерный"},
		"idealised":           {"идеализированный"},
		"inconsiderate":       {"невнимательный", "бесцеремонный", "равнодушный", "неделикатный"},
		"indecisive":          {"нерешительный"},
		"self-assured":        {"самоуверенный", "уверенный в себе"},
		"self-absorbed":       {"эгоистичный", "самовлюбленный"},
		"self-centred":        {"эгоцентричный"},
		"self-congratulatory": {"самодовольный"},
		"self-deprecating":    {"самокритичный", "самоуничижительный"},
		"self-important":      {"самовлюбленный", "самодовольный"},
		"tactful":             {"тактичный", "дипломатичный"},
		"well-adjusted":       {"уравновешенный"},
		"well-bred":           {"хорошо воспитанный", "благородный"},
		"well-brought-up":     {"хорошо воспитанный"},
		"well-dressed":        {"хорошо одетый"},
	},
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"adolescent":     {"подросток", "отрок"},
		"characteristic": {"характеристика", "особенность", "свойство", "признак", "черта", "качество"},
		"trait":          {"черта", "признак", "особенность", "свойство"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
