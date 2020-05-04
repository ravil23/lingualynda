package paulineunit1

import (
	"github.com/ravil23/lingualynda/telegrambot/collection"
)

var VocabularyAdjectives = collection.NewVocabulary(
	map[collection.Term][]collection.Translation{
		"apprehensive": {"опасающийся", "боязливый"},
		"assertive":    {"напористый", "настойчивый"},
		"clumzy":       {"неуклюжий", "неловкий"},
	},
)

var VocabularyNouns = collection.NewVocabulary(
	map[collection.Term][]collection.Translation{
		"adolescent":     {"подросток", "отрок"},
		"characteristic": {"характеристика", "особенность", "свойство", "признак", "черта", "качество"},
		"trait":          {"черта", "признак", "особенность", "свойство"},
	},
)

var VocabularyTotal *collection.Vocabulary

func init() {
	VocabularyTotal = collection.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns)
}
