package introductory

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyStart = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"To tell the truth":      {"говоря по правде"},
		"First of all":           {"прежде всего", "в первую очередь"},
		"By the way":             {"кстати", "между прочим"},
		"As far as I know":       {"насколько я знаю"},
		"I am convinced":         {"я уверен", "убеждён"},
		"I would like to stress": {"я бы хотел отметить"},
		"I suppose":              {"я полагаю что"},
		"I believe":              {"я полагаю что"},
		"I consider":             {"я полагаю что"},
		"I guess":                {"я полагаю что"},
		"If you ask me":          {"если вы спрашиваете меня"},
		"To my way of thinking":  {"по-моему", "мне кажется"},
		"As a matter of fact":    {"по сути"},
		"It goes without saying": {"само собой разумеется"},
		"To begin with":          {"для начала", "сначала", "начнём с того что"},
	},
)

var VocabularyLink = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"In other words":              {"другими словами", "иначе говоря"},
		"As I have already mentioned": {"как я уже сказал"},
		"On the one hand":             {"с одной стороны"},
		"On the other hand":           {"с другой стороны"},
		"Moreover":                    {"более того", "кроме того"},
		"In addition":                 {"помимо этого"},
		"As to":                       {"что касается"},
		"However":                     {"однако"},
		"In this case":                {"в этом случае"},
		"Although":                    {"хотя"},
		"Nevertheless":                {"тем не менее", "однако"},
	},
)

var VocabularyFinish = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"To make a long story short": {"короче говоря", "в двух словах"},
		"In a nutshell":              {"короче говоря", "в двух словах"},
		"That is why":                {"поэтому"},
		"In conclusion":              {"в заключение"},
		"To conclude":                {"в заключение"},
		"As a result of":             {"в результате"},
		"Therefore":                  {"следовательно"},
		"All in all":                 {"в конечном счёте", "в общем"},
		"Finally":                    {"наконец", "в конце концов"},
		"To sum it up":               {"в итоге"},
		"To draw the conclusion":     {"подводя итог"},
		"To top it off":              {"наконец", "в завершение всего"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyStart).
		Update(VocabularyLink).
		Update(VocabularyFinish)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
