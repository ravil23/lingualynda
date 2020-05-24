package paulineunit1

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
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

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"adolescent":     {"подросток", "отрок"},
		"characteristic": {"характеристика", "особенность", "свойство", "признак", "черта", "качество"},
		"trait":          {"черта", "признак", "особенность", "свойство"},
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
