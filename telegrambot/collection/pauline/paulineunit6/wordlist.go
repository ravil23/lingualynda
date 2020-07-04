package paulineunit6

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyAdjectives = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"catchy":      {"броский"},
		"competing":   {"конкурирующий", "соперничающий"},
		"distracting": {"отвлекающий"},
		"disturbing":  {"тревожный", "беспокойный"},
		"infuriating": {"бесящий", "приводящий в ярость"},
		"invasive":    {"агрессивный", "инвазивный"},
		"irritating":  {"раздражающий"},
		"monetary":    {"денежный", "монетарный", "валютный"},
		"persuasive":  {"убедительный"},
		"promotional": {"рекламный"},
		"repetitive":  {"повторяющийся"},
		"slick":       {"ловкий", "скользкий", "хитрый"},
		"ubiquitous":  {"повсеместный", "вездесущий"},
		"unavoidable": {"неизбежный", "неминуемый"},
	},
)

var VocabularyNouns = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"billboard":            {"рекламный щит", "билборд", "плакат", "афиша", "доска объявлений"},
		"branding":             {"бренд", "клеймо"},
		"distributor":          {"дистрибьютор", "распространитель", "распределитель", "поставщик"},
		"flyer":                {"флаер", "листовка", "объявление"},
		"gimmick":              {"трюк", "уловка"},
		"jingle":               {"заставка", "песенка"},
		"manufacturer":         {"изготовитель", "производитель", "промышленник"},
		"marketing":            {"маркетинг", "продажа", "торговля", "реклама"},
		"online retailer":      {"онлайн продавец"},
		"ploy":                 {"уловка", "тактический ход"},
		"rebate":               {"скидка", "возврат"},
		"sales representative": {"торговый представитель"},
		"slogan":               {"лозунг", "слоган", "девиз"},
		"telemarketing":        {"телемаркетинг"},
		"vendor":               {"поставщик", "продавец", "подрядчик", "торговец"},
	},
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{
		"assure":     {"заверять", "гарантировать", "обеспечивать", "убеждать"},
		"cultivate":  {"культивировать", "развивать", "выращивать"},
		"entice":     {"соблазнить", "заманить", "побудить"},
		"induce":     {"побудить", "заставить", "вызывать", "стимулировать", "вынудить", "склонить"},
		"oblige":     {"обязать", "заставить", "вынудить"},
		"pressurise": {"оказывать давление"},
		"reassure":   {"заверить", "успокоить", "убедить", "обнадежить"},
		"tempt":      {"искушать", "соблазнять"},
		"urge":       {"просить", "призывать", "побудить", "убедить", "обращаться к"},
	},
)

var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
