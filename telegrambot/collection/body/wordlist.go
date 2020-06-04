package body

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"ankle":      {"лодыжка"},
		"armpit":     {"подмышка"},
		"calf":       {"икроножная мышца"},
		"cheek":      {"щека"},
		"chest":      {"грудь"},
		"chin":       {"подбородок"},
		"earlobe":    {"мочка уха"},
		"elbow":      {"локоть"},
		"eyebrow":    {"бровь"},
		"eyelash":    {"ресница"},
		"forehead":   {"лоб"},
		"heel":       {"пятка"},
		"hip":        {"бок"},
		"intestines": {"кишечник"},
		"jaw":        {"челюсть"},
		"kidney":     {"почка"},
		"knuckle":    {"кулак"},
		"lip":        {"губа"},
		"liver":      {"печень"},
		"lung":       {"легкое"},
		"neck":       {"шея"},
		"nostril":    {"ноздря"},
		"palm":       {"ладонь"},
		"pelvis":     {"таз"},
		"rib":        {"ребро"},
		"shin":       {"голень"},
		"sole":       {"ступня"},
		"spine":      {"позвоночник"},
		"stomach":    {"желудок"},
		"temple":     {"висок"},
		"thigh":      {"бедро"},
		"throat":     {"горло"},
		"thumb":      {"большой палец"},
		"veins":      {"вены"},
		"waist":      {"талия"},
		"wrist":      {"запястье"},
	},
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"clap":    {"хлопать"},
		"frisk":   {"обыскивать"},
		"hug":     {"обнять"},
		"nod":     {"кивать головой"},
		"nudge":   {"подталкивать"},
		"pat":     {"похлопать"},
		"pinch":   {"зажимать"},
		"rub":     {"тереть"},
		"shove":   {"совать", "толкнуть"},
		"slap":    {"шлепать"},
		"smack":   {"хлопать", "причмокивать"},
		"sniff":   {"сопеть"},
		"spit":    {"плюнуть"},
		"squeeze": {"выжимать"},
		"stroke":  {"погладить"},
		"swallow": {"глотать"},
		"thump":   {"стучать"},
		"tickle":  {"щекотать"},
		"wink":    {"подмигивать"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary
var AllVocabularies []*schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*schema.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
