package paulineunit9

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"appalling":     {"ужасающий", "чудовищный"},
		"aquatic":       {"водный"},
		"hypocritical":  {"лицемерный"},
		"impartial":     {"беспристрастный", "объективный"},
		"irresponsible": {"безответственный"},
		"nocturnal":     {"ночной"},
		"outrageous":    {"возмутительный", "вопиющий"},
		"risky":         {"рискованный"},
		"venomous":      {"ядовитый", "злой"},
	},
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"burrow":        {"нора"},
		"colony":        {"поселение", "колония"},
		"deforestation": {"вырубка лесов", "уничтожение лесов"},
		"degradation":   {"деградация6 разложение", "разрушение", "упадок"},
		"demise":        {"кончина", "смерть", "гибель,крах"},
		"devastation":   {"разрушение", "опустошение", "ущерб"},
		"flooding":      {"наводнение", "затопление"},
		"habitat":       {"местообитание", "среда обитания", "естественная среда"},
		"harm":          {"ущерб", "вред", "урон", "страдания"},
		"herd":          {"стадо"},
		"hive":          {"улей"},
		"lifespan":      {"продолжительности жизни", "срок службы"},
		"pack":          {"стая", "пачка", "пакет", "рюкзак", "упаковка"},
		"parasite":      {"паразит", "тунеядец"},
		"prey":          {"добыча", "жертва", "дичь"},
		"rival":         {"соперник", "конкурент"},
		"swarm":         {"рой"},
		"threshold":     {"порог", "предел", "критерий"},
	},
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"flourish":  {"процветать", "расцвести"},
		"forage":    {"добывать себе"},
		"hibernate": {"зимовать", "проспать", "быть в спящем режиме"},
		"lay":       {"лежать", "отложить", "закладывать"},
		"poach":     {"переманить", "увести", "спереть"},
		"thrive":    {"процветать", "преуспевать"},
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
