package phrasalverbs

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

var VocabularyVerbs = schema.NewVocabulary(
	map[schema.Term][]schema.Translation{"look out": {"будь осторожен"},
		"blow out":                 {"задуть"},
		"break down":               {"перестать работать", "сломаться"},
		"bring in":                 {"принести"},
		"bring up":                 {"воспитывать"},
		"carry out":                {"выполнять"},
		"cash in on":               {"заработать на"},
		"catch up":                 {"не отставать"},
		"check in":                 {"зарегистрироваться"},
		"clean out":                {"обчистить", "обворовать"},
		"come in":                  {"войти"},
		"contend with":             {"справляться с"},
		"count out":                {"не включать"},
		"crawl out":                {"выползать"},
		"cross out":                {"зачеркнуть"},
		"cut in":                   {"прерывать"},
		"cut out":                  {"вырезать"},
		"deprive of":               {"лишать кого-то"},
		"dive in":                  {"погружаться в воду"},
		"draw out":                 {"чересчур растягивать во времени"},
		"drop in":                  {"заходить непреднамеренно"},
		"drop out":                 {"прекратить дело"},
		"eat out":                  {"поесть вне дома"},
		"factor in":                {"включить в расчеты"},
		"fall out":                 {"ссоритья"},
		"fill in":                  {"заполнять форму"},
		"fill out":                 {"заполнять форму"},
		"find out about":           {"получить информацию"},
		"find out":                 {"понять"},
		"fit in with":              {"сливаться c"},
		"get by":                   {"управлять"},
		"get on":                   {"справиться", "ладить", "садиться в транспорт"},
		"get out":                  {"избегать дела"},
		"give out":                 {"раздавать"},
		"give up":                  {"сдаваться"},
		"go out":                   {"перестать гореть"},
		"hand out":                 {"раздавать"},
		"join in":                  {"присоединиться"},
		"jut out":                  {"торчать", "выпирать"},
		"keep up with":             {"не отставать"},
		"kick away":                {"отбрасывать"},
		"lash out":                 {"отчитывать"},
		"lead-in":                  {"введение"},
		"leave out from":           {"исключать из"},
		"let in":                   {"впустить"},
		"look forward to":          {"ждать с нетерпением чего-либо"},
		"look up at":               {"поднять взгляд на"},
		"move in":                  {"заезжать"},
		"move out":                 {"съезжать"},
		"pay off":                  {"заплатить", "окупаться"},
		"pick out":                 {"отделять", "перебирать"},
		"pick up":                  {"подвозить, подбирать"},
		"pitch in":                 {"вносить свой вклад", "помогать"},
		"plug in":                  {"подключить к электричеству"},
		"point out":                {"отметить", "подчеркнуть"},
		"pull out":                 {"вытягивать"},
		"put on":                   {"накидывать"},
		"put out":                  {"тушить"},
		"put up with":              {"терпеть"},
		"put up":                   {"поднимать"},
		"rub out":                  {"стереть"},
		"run away from":            {"убегать от"},
		"run out":                  {"закончиться"},
		"set in":                   {"происходить"},
		"share out":                {"распространять"},
		"show in":                  {"представить"},
		"sort out":                 {"решать проблему"},
		"speak out":                {"высказаться откровенно"},
		"spread out":               {"распределять"},
		"squash in":                {"втиснуться"},
		"stand out from the crowd": {"выделяться из толпы"},
		"take in":                  {"обмануть"},
		"take off":                 {"взлетать", "снимать одежду"},
		"tell off":                 {"ругать"},
		"test out":                 {"протестировать"},
		"throw away":               {"выбрасывать"},
		"to be boxed in":           {"ограничивать"},
		"try out":                  {"испробовать"},
		"turn on":                  {"включать"},
		"turn out to be":           {"оказаться"},
		"turn out":                 {"выключить", "обернуться", "оказаться"},
		"usher in":                 {"влиться в, начинать"},
		"wake up":                  {"разбудить"},
		"work out":                 {"делать физические упражнения", "развиваться"},
		"work something out":       {"производить какие-то расчеты"},
		"yell out":                 {"вскрикивать для привлечения внимания"},
	},
)

var VocabularyTotal *schema.Vocabulary
var VocabularyEngToRus *schema.Vocabulary
var VocabularyRusToEng *schema.Vocabulary

func init() {
	VocabularyEngToRus = schema.NewEmptyVocabulary().
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	VocabularyTotal = schema.NewEmptyVocabulary().
		Update(VocabularyEngToRus).
		Update(VocabularyRusToEng)
}
