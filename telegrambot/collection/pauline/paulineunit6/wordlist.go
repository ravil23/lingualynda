package paulineunit6

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyAdjectives = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
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

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyAdjectives).
		Update(VocabularyNouns).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
