package paulineunit8

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyPhrases = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"leader of the":     {"руководитель", "глава", "предводитель"},
		"opposition":        {"оппозиция", "противодействие", "сопротивление"},
		"lose control of":   {"потерять контроль над"},
		"means tested":      {"нуждаемость"},
		"social welfare":    {"социальное обеспечение", "социальная защита"},
		"social well-being": {"социальное благостояние", "социальное благополучие"},
	},
)

var VocabularyNouns = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"bureaucrat":       {"бюрократ", "чиновник"},
		"civil servant":    {"гражданский служащий", "госслужащий"},
		"community":        {"сообщество", "община", "население", "общественность"},
		"helathcare":       {"здравоохранение", "медицинское обслуживание"},
		"mayor":            {"мэр", "майор", "градоначальник"},
		"member of":        {"участник", "сотрудник"},
		"parliament":       {"парламент", "депутат", "собрание"},
		"military service": {"военная служба"},
		"notion":           {"понятие", "концепция", "идея", "представление", "мысль"},
		"old-age pension":  {"пенсия по старости"},
		"safety net":       {"сеть безопасности"},
		"social services":  {"социальные услуги", "социальное обеспечение"},
		"state":            {"государство", "штат", "состояние"},
		"unemployment":     {"безработица", "увольнения"},
		"benefit":          {"пособие", "благо", "выгода", "польза", "преимущество", "льгота"},
	},
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"miscalculate":  {"неверно рассчитывать"},
		"misinform":     {"дезинформировать", "ввести в заблуждение"},
		"misdiagnose":   {"неправильно диагностировать", "ставить неверный диагноз"},
		"mismanage":     {"неумело справляться"},
		"misunderstand": {"неправильно понять"},
		"misinterpret":  {"искажать", "извращать", "неправильно истолковывать"},
		"mislead":       {"ввести в заблуждение", "обмануть", "запутать"},
		"misjudge":      {"недооценивать"},
		"mistrust":      {"не доверять"},
		"subsidise":     {"субсидировать"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyNouns).
		Update(VocabularyPhrases).
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
}
