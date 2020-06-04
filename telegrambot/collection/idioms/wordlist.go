package idioms

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyIdioms = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"elbow grease":       {"работа до 7 пота"},
		"get off my chest":   {"избавиться от тяжелого бремени"},
		"hard to stomach":    {"тяжело терпеть"},
		"rule of thumb":      {"как правило", "по опыту"},
		"to pay lip service": {"пустые разговоры"},
		"toe the line":       {"стоять по струнке смирно"},
		"what a cheek":       {"какая наглость"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary
var AllVocabularies []*schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyIdioms)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*schema.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
