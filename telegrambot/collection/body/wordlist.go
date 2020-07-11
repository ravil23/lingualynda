package body

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
