package adjectives

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"ancient":           {"древний", "старинный", "античный"},
		"argumentative":     {"спорящий"},
		"arrogant":          {"высокомерный"},
		"astounding":        {"удивительный", "поразительный"},
		"attention-seeking": {"показушный"},
		"awful":             {"ужасный"},
		"boiling":           {"кипящий"},
		"calm":              {"спокойный", "тихий", "невозмутимый"},
		"cheerful":          {"жизнерадостный"},
		"excellent":         {"замечательный", "прекрасный", "потрясающий"},
		"exhausted":         {"измученный"},
		"extrovert":         {"экстроверт"},
		"fantastic":         {"фантастический"},
		"fascinating":       {"удивительный", "очаровательный"},
		"filthy":            {"мерзко грязный"},
		"freezing":          {"леденящий"},
		"furious":           {"взбешенный"},
		"giant":             {"гигантский"},
		"gigantic":          {"гигантский"},
		"good in a team":    {"командный"},
		"gorgeous":          {"великолепный", "прекрасный"},
		"hideous":           {"отвратительный", "ужасный", "безобразный"},
		"hilarious":         {"очень смешной"},
		"horrible":          {"ужасный"},
		"hot-tempered":      {"вспыльчивый"},
		"huge":              {"гигантский"},
		"humorous":          {"шутливый"},
		"individualistic":   {"индивидуальный"},
		"introvert":         {"интроверт"},
		"jealous":           {"ревнивый", "завистливый"},
		"needy":             {"неполноценный"},
		"rebellious":        {"бунтарский", "непослушный"},
		"resilient":         {"жизнерадостный"},
		"self-sufficient":   {"самодостаточный"},
		"spotless":          {"безупречно чистый"},
		"starving":          {"безумно голодный"},
		"talkative":         {"разговорчивый", "болтливый"},
		"terrible":          {"ужасный"},
		"terrific":          {"прекрасный"},
		"terrifying":        {"ужасный"},
		"tiny":              {"малюсенький"},
		"uncommunicative":   {"необщительный", "замкнутый"},
		"wonderful":         {"замечательный", "прекрасный", "чудесный"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary
var AllVocabularies []*schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*schema.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
