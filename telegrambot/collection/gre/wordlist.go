package gre

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var Vocabulary = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"abash":       {"смущать", "конфузить"},
		"abate":       {"уменьшать", "ослаблять"},
		"abdicate":    {"отрекаться (от власти)"},
		"aberrant":    {"сбившийся с пути", "ненормальный"},
		"abet":        {"помогать", "поощрять", "содействовать"},
		"abeyance":    {"состояние неопределенности"},
		"abhor":       {"ненавидеть", "презирать"},
		"abject":      {"подлый", "низкий", "ужасный"},
		"abjure":      {"отказываться", "отрекаться"},
		"abnegation":  {"отрицание", "непринятие"},
		"abomination": {"неприязнь", "враждебность"},
		"abortive":    {"бесплодный", "бесполезный"},
		"abound":      {"изобиловать"},
		"aboveboard":  {"открытый", "честный", "прямой"},
		"abridge":     {"сокращать", "уменьшать", "снижать"},
		//"abrogate":    {"отменять", "аннулировать"},
		//"abrupt":      {"внезапный", "обрывистый", "грубый"},
		//"abscond":     {"избегать", "уходить от"},
		//"absolve":     {"освобождать", "прощать"},
		//"abstemious":  {"умеренный", "скромный"},
		//"abstruse":    {"трудный для понимания", "глубокий"},
		//"abut":        {"примыкать", "граничить"},
		//"abysmal":     {"низкий", "отвратительный"},
		//"acclaim":     {"приветствовать", "объявлять"},
		//"accolade":    {"похвала", "одобрение"},
		//"accomplice":  {"сообщник", "соучастник"},
		//"accost":      {"приставать", "обращаться"},
		//"accretion":   {"прирост", "увеличение"},
		//"acerbic":     {"едкий", "язвительный"},
		//"acidulous":   {"едкий", "язвительный", "кисловатый"},
		//"acme":        {"кульминация"},
		//"acolyte":     {"помощник", "последователь"},
		//"acquiesce":   {"уступать", "молча соглашаться"},
		//"acquit":      {"оправдывать"},
		//"acrimonious": {"язвительный"},
		//"acumen":      {"проницательность"},
		//"acute":       {"острый", "резкий", "проницательный"},
		//"adage":       {"афоризм", "поговорка"},
		//"adamant":     {"непреклонный", "категоричный"},
		//"adherent":    {"приверженец", "адепт"},
		//"adjunct":     {"дополнение", "принадлежность"},
		//"admonish":    {"предостерегать"},
		//"adroit":      {"умелый", "искусный"},
		//"adulation":   {"низкопоклонство", "лесть"},
		//"adulterate":  {"портить примесями", "ухудшать"},
		//"advent":      {"наступление", "начало"},
		//"adverse":     {"неблагоприятный", "враждебный"},
		//"affinity":    {"близость", "родство", "симпатия"},
		//"afflict":     {"причинять боль", "страдания"},
		//"affluent":    {"богатый", "обильный"},
		//"aggrandize":   {"увеличивать", "возвеличивать"},
		//"alacrity":     {"готовность", "проворство"},
		//"allay":        {"успокаивать", "подавлять"},
		//"alleged":      {"предполагаемый"},
		//"alleviate":    {"облегчать", "смягчать"},
		//"alloy":        {"сплав", "примесь"},
		//"allude":       {"намекать", "указывать"},
		//"allure":       {"очарование", "привлекательность"},
		//"aloof":        {"отчужденный", "равнодушный"},
		//"alter":        {"изменять"},
		//"altercate":    {"препираться", "ссорится"},
		//"amalgamate":   {"объединять", "смешивать"},
		//"ambulatory":   {"ходячий", "передвижной"},
		//"ameliorate":   {"улучшать"},
		//"amicable":     {"дружелюбный"},
		//"amorous":      {"влюбленный", "влюбчивый"},
		//"ample":        {"обильный", "щедрый"},
		//"amplify":      {"расширять", "увеличивать"},
		//"amuck":        {"дикий", "бешеный", "неистовый"},
		//"ancillary":    {"вспомогательный", "добавочный"},
		//"anguish":      {"мука", "страдания"},
		//"anodyne":      {"болеутоляющее средство"},
		//"antecedent":   {"предок", "прародитель"},
		//"antedate":     {"предшествовать"},
		//"antediluvian": {"допотопный", "старомодный"},
		//"anterior":     {"предшествующий", "передний"},
		//"apace":        {"быстрый"},
		//"aperture":     {"отверстие", "дыра", "щель"},
		//"apex":         {"высшая точка"},
		//"aplomb":       {"невозмутимость", "самообладание"},
		//"apocryphal":   {"недостоверный", "сомнительный"},
		//"apostasy":     {"отступничество", "измена"},
		//"apothegm":     {"меткое изречение", "апофегма"},
		//"appalled":     {"потрясенный", "шокированный"},
		//"appellation":  {"имя", "название"},
		//"appendage":    {"придаток", "привесок", "дополнение"},
		//"apposite":     {"подходящий", "уместный", "удачный"},
		//"appraise":     {"оценивать"},
		//"apprehension": {"опасение", "мрачное предчувствие"},
		//"approbation":  {"одобрение"},
		//"aptitude":     {"пригодность", "склонность"},
		//"arbitrary":    {"произвольный"},
		//"arcadia":      {"идиллия", "утопия"},
		//"arcane":       {"загадочный", "непонятный"},
		//"ardent":       {"пылкий", "горячий"},
		//"ardor":        {"пыл", "страсть", "рвение"},
		//"arduous":      {"трудный", "тяжелый"},
		//"arrant":       {"настоящий", "отъявленный"},
		//"array":        {"набор", "строй", "количество"},
		//"arriviste":    {"карьерист", "честолюбец", "выскочка"},
		//"artful":       {"ловкий", "хитрый"},
		//"artifice":     {"уловка", "трюк", "прием"},
		//"artless":      {"безвкусный", "простой"},
		//"ascend":       {"подниматься", "восходить"},
		//"ascertain":    {"устанавливать", "выявлять"},
		//"ascetic":      {"аскет"},
		//"asinine":      {"глупый", "неумный", "неразумный"},
		//"askance":      {"наклонный", "косой"},
		//"asperity":     {"суровость", "лишения"},
		//"aspersion":    {"клевета"},
		//"aspirant":     {"кандидат"},
		//"aspire":       {"стремиться к чему-либо"},
		//"asset":        {"актив", "ценность"},
		//"assiduous":    {"усердный", "прилежный"},
		//"assuage":      {"смягчать", "ослаблять"},
		//"astringent":   {"суровый", "строгий"},
		//"astute":       {"хитрый", "ловкий"},
		//"attain":       {"достигать", "добираться"},
		//"attenuated":   {"тонкий", "ослабленный"},
		//"attest":       {"подтверждать"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(Vocabulary)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
