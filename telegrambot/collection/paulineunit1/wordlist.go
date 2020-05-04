package paulineunit1

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"apprehensive": {"опасающийся", "боязливый"},
		"assertive":    {"напористый", "настойчивый"},
		"clumzy":       {"неуклюжий", "неловкий"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"adolescent":     {"подросток", "отрок"},
		"characteristic": {"характеристика", "особенность", "свойство", "признак", "черта", "качество"},
		"trait":          {"черта", "признак", "особенность", "свойство"},
	},
)

var VocabularyTotal *schema.Vocabulary

func init() {
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns)
}
