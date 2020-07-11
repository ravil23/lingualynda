package idioms

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyIdioms = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"elbow grease":       {"работа до 7 пота"},
		"get off my chest":   {"избавиться от тяжелого бремени"},
		"hard to stomach":    {"тяжело терпеть"},
		"rule of thumb":      {"как правило", "по опыту"},
		"to pay lip service": {"пустые разговоры"},
		"toe the line":       {"стоять по струнке смирно"},
		"what a cheek":       {"какая наглость"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyIdioms)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
