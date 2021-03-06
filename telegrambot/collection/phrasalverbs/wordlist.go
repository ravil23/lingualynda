package phrasalverbs

import (
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

var VocabularyVerbs = entity.NewVocabulary(
	map[entity.Term][]entity.Translation{
		"add up":                   {"быть правдоподобным", "сходиться"},
		"ask after somebody":       {"спрашивать о ком-то"},
		"ask somebody out":         {"пригласить на свидание"},
		"back off":                 {"отвали"},
		"be back":                  {"вернуться"},
		"be cut off":               {"быть отрезанным от"},
		"be knocked down":          {"быть снесенным"},
		"be off":                   {"уезжать откуда-то"},
		"bend down":                {"наклониться", "нагнуться"},
		"blow away":                {"сдуть"},
		"blow out":                 {"задуть"},
		"blow something down":      {"свалить что-то"},
		"blow up":                  {"клеветать", "раздуть"},
		"break down into parts":    {"делить на части"},
		"break down":               {"перестать работать", "сломаться"},
		"break up":                 {"расставаться", "разломать"},
		"bring in":                 {"принести"},
		"bring up":                 {"воспитывать"},
		"burn down":                {"быть сожженым"},
		"call (somebody) back":     {"перезвонить"},
		"call something off":       {"отменить что-то"},
		"calm (somebody) down":     {"успокоиться"},
		"carry on (with)":          {"продолжать делать"},
		"carry out":                {"выполнять"},
		"cash in on":               {"заработать на"},
		"catch somebody up":        {"догнать"},
		"catch up (with somebody)": {"догнать"},
		"catch up":                 {"не отставать"},
		"check in":                 {"зарегистрироваться"},
		"clean out":                {"обчистить", "обворовать"},
		"clean up":                 {"чистить"},
		"clear up":                 {"чистить"},
		"close down":               {"прекратить бизнес"},
		"come down":                {"передавать через поколения", "согласиться на пониженную цену"},
		"come in":                  {"войти"},
		"come off":                 {"быть успешным", "сходить (пятно)"},
		"come over":                {"прийти"},
		"come up":                  {"подходить"},
		"contend with":             {"справляться с"},
		"cordon off something":     {"оцеплять территорию"},
		"count on something":       {"рассчитывать на что-то"},
		"count out":                {"не включать"},
		"crawl out":                {"выползать"},
		"cross out":                {"зачеркнуть"},
		"cut down (on something)":  {"делать что-то реже"},
		"cut in":                   {"прерывать"},
		"cut off something":        {"перестать поставки чего-либо"},
		"cut out":                  {"вырезать"},
		"cut something down":       {"срезать что-то"},
		"depend on something":      {"зависеть от чего-либо"},
		"deprive of":               {"лишать кого-то"},
		"dive in":                  {"погружаться в воду"},
		"doze off":                 {"засыпать"},
		"draw out":                 {"чересчур растягивать во времени"},
		"drive away":               {"уехать"},
		"drive off":                {"проехать"},
		"drive on":                 {"продолжать ехать"},
		"drone on":                 {"нудно рассказывать"},
		"drop by":                  {"зайти", "навестить"},
		"drop in":                  {"заходить непреднамеренно"},
		"drop off":                 {"засыпать"},
		"drop out":                 {"прекратить дело"},
		"dwell on something":       {"долго обсуждать одну тему"},
		"dwell on":                 {"зацикливаться"},
		"eat out":                  {"поесть вне дома"},
		"end up doing":             {"в итоге сделать"},
		"end up somewhere":         {"в итоге оказались где-то"},
		"factor in":                {"включить в расчеты"},
		"fall out":                 {"соритья"},
		"fill in":                  {"заполнять форму"},
		"fill out":                 {"заполнять форму"},
		"find out about":           {"получить информацию"},
		"find out":                 {"понять"},
		"finish off":               {"обессилеть чтобы продолжать"},
		"finish something off":     {"делать последнюю часть чего-либо"},
		"finish up":                {"завершить"},
		"fit in with":              {"сливаться c"},
		"fix up":                   {"запланировать"},
		"fly away":                 {"улететь"},
		"focus on something":       {"сконцентрироваться на чем-то"},
		"freak out":                {"беситься"},
		"frown on something":       {"нахмуриться на что-то"},
		"get away with something":  {"безнаказанно сделать что-то неправильно"},
		"get away":                 {"уйти"},
		"get back to somebody":     {"ответить кому-то по телефону"},
		"get by":                   {"управлять"},
		"get on something":         {"продолжать делать, то что должен"},
		"get on with":              {"ладить"},
		"get on":                   {"справиться", "прогрессировать", "садиться в транспорт"},
		"get out":                  {"избегать дела"},
		"get somebody down":        {"ощущать стресс"},
		"give back":                {"вернуть"},
		"give out":                 {"раздавать"},
		"give something away":      {"отдать что-то что больше не нужно"},
		"give something up":        {"перестать делать что-то"},
		"give up":                  {"сдаваться"},
		"go away":                  {"проводить время вне дома"},
		"go down":                  {"запомниться", "перестать работать"},
		"go off":                   {"взрываться", "зазвонить"},
		"go on and on":             {"продолжается"},
		"go on":                    {"случаться", "продолжать"},
		"go out":                   {"перестать гореть"},
		"go up":                    {"подходить"},
		"grow up":                  {"взрослеть"},
		"gun down":                 {"быть подстреленным"},
		"hand out":                 {"раздавать"},
		"hang on something":        {"дождаться чего-либо"},
		"hang on":                  {"секундочку"},
		"hang up":                  {"повесить трубку"},
		"happen on something":      {"случиться чему-то"},
		"have something off":       {"проводить время вне работы"},
		"hinge on something":       {"зависеть от чего-либо"},
		"hinge on":                 {"зависеть от"},
		"hit somebody back":        {"дать сдачу"},
		"jazz up":                  {"приукрасить (неформально)"},
		"join in":                  {"присоединиться"},
		"jut out":                  {"торчать", "выпирать"},
		"keep away (from)":         {"посторониться", "обойти"},
		"keep on":                  {"делать постоянно повторяющейся"},
		"keep up (with somebody)":  {"идти в одном темпе"},
		"keep up with":             {"не отставать"},
		"keep up":                  {"продолжать в том же духе"},
		"kick away":                {"отбрасывать"},
		"kick off":                 {"начинать"},
		"knock down":               {"сносить", "убеждать понизить цену"},
		"land up":                  {"происходить чему-то не запланированному (неформально)"},
		"lash out":                 {"отчитывать"},
		"lead-in":                  {"введение"},
		"leave on":                 {"оставлять включенным"},
		"leave out from":           {"исключать из"},
		"led on":                   {"привели к"},
		"let in":                   {"впустить"},
		"let somebody down":        {"разочаровать кого-то"},
		"lie down":                 {"прилечь"},
		"lift up":                  {"приподнять"},
		"live on":                  {"долго обсуждать одну тему"},
		"look back (on something)": {"думать о прошлом"},
		"look forward to":          {"ждать с нетерпением чего-либо"},
		"look out":                 {"будь осторожен"},
		"look up at":               {"поднять взгляд на"},
		"made up of something":     {"состоит из чего-то"},
		"make up something":        {"являться чем-то"},
		"make up":                  {"мириться"},
		"mix up":                   {"путать"},
		"move in":                  {"заезжать"},
		"move out":                 {"съезжать"},
		"nod off":                  {"засыпать"},
		"pass out":                 {"терять сознание"},
		"pay back money":           {"вернуть деньги"},
		"pay off":                  {"заплатить", "окупаться"},
		"pay somebody back":        {"вернуть кому-то деньги"},
		"pep something up":         {"сделать более интересным (неформально)"},
		"phone (somebody) back":    {"перезвонить"},
		"pick out":                 {"отделять", "перебирать"},
		"pick something up":        {"поднять что-то"},
		"pick up on":               {"замечать", "реагировать на что-то"},
		"pick up":                  {"подвозить", "подбирать"},
		"pin somebody down":        {"удерживать в горизонтальном положении"},
		"piss off":                 {"бесить (разговорный)"},
		"pitch in":                 {"вносить свой вклад", "помогать"},
		"play on":                  {"продолжать играть"},
		"plug in":                  {"подключить к электричеству"},
		"point out":                {"отметить", "подчеркнуть"},
		"press on":                 {"долго обсуждать одну тему"},
		"propped up":               {"опираться"},
		"pull out":                 {"вытягивать"},
		"put back":                 {"положить обратно"},
		"put down":                 {"убить"},
		"put off doing something":  {"причина, чтобы что-то не делать"},
		"put on":                   {"накидывать", "набирать вес", "включать (музыку)"},
		"put out":                  {"тушить"},
		"put somebody down":        {"внести в список"},
		"put something away":       {"убрать что из поля зрения"},
		"put something down":       {"положить что-то"},
		"put something off":        {"задерживать что-то"},
		"put something up":         {"повесить что-то"},
		"put up with":              {"терпеть"},
		"put up":                   {"поднимать"},
		"read on":                  {"долго обсуждать одну тему"},
		"rely on something":        {"полагаться на что-то"},
		"ride off":                 {"уехать"},
		"ring (somebody) back":     {"перезвонить"},
		"rip somebody off":         {"быть обманутым кем-то"},
		"round off":                {"завершить"},
		"rub out":                  {"стереть"},
		"run away from":            {"убегать от"},
		"run off":                  {"убежать"},
		"run out":                  {"закончиться"},
		"see somebody off":         {"провожать"},
		"set in":                   {"происходить"},
		"set off":                  {"начать"},
		"set up":                   {"организовать"},
		"shake off":                {"потеряться"},
		"share out":                {"распространять"},
		"shore up":                 {"опираться", "укрепить"},
		"shout back":               {"кричать в ответ"},
		"shout down somebody":      {"перекрикивать"},
		"shove off":                {"убирайся"},
		"show in":                  {"представить"},
		"show off":                 {"красоваться"},
		"show up":                  {"появляться"},
		"shrug off":                {"не считать важным"},
		"shut down":                {"прекратить бизнес", "отключить"},
		"sit down":                 {"присесть"},
		"slam down":                {"выбросить смяв"},
		"slow down":                {"двигаться медленнее"},
		"smile back":               {"улыбнуться"},
		"sort out":                 {"решать проблему"},
		"speak out":                {"высказаться откровенно"},
		"split off":                {"отделиться от группы"},
		"split up":                 {"расставаться", "разделить"},
		"spread out":               {"распределять"},
		"spring on":                {"вывалить на"},
		"squash in":                {"втиснуться"},
		"stand out from the crowd": {"выделяться из толпы"},
		"stand up":                 {"встать", "подтверждаться"},
		"start off":                {"начать"},
		"switch off":               {"переключать на выключенное состояние"},
		"switch on":                {"переключать на включенное состояние"},
		"take away":                {"забрать"},
		"take in":                  {"обмануть"},
		"take off":                 {"взлетать", "снимать одежду"},
		"take something down":      {"снять что-то"},
		"take up space":            {"занимать место"},
		"take up time":             {"отнимать время"},
		"take up":                  {"начать заниматься"},
		"tell off":                 {"ругать"},
		"tell somebody off":        {"сердито говорить"},
		"test out":                 {"протестировать"},
		"throw away":               {"выбрасывать"},
		"throw something away":     {"выбросить что-то в мусор"},
		"ticking off":              {"отчитывать"},
		"tidy up":                  {"чистить"},
		"tie down":                 {"привязать", "ограничивать свободу"},
		"to be boxed in":           {"ограничивать"},
		"try on":                   {"примерять одежду"},
		"try out":                  {"испробовать"},
		"turn off":                 {"выключать"},
		"turn on":                  {"включать"},
		"turn out to be":           {"оказаться"},
		"turn out":                 {"выключить", "обернуться", "оказаться"},
		"turn something down":      {"ослабить мощность чего-либо", "отказаться от чего-либо"},
		"turn something up":        {"увеличить мощность чего-либо"},
		"turn up":                  {"появляться"},
		"use something up":         {"использовать полностью"},
		"usher in":                 {"влиться в", "начинать"},
		"wake up":                  {"разбудить"},
		"walk back":                {"идти обратно"},
		"walk off":                 {"уйти"},
		"walk on":                  {"продолжать прогуливаться"},
		"walk up (to)":             {"подходить"},
		"wash up":                  {"мыть"},
		"wave back":                {"помахать"},
		"work off":                 {"отрабатывать", "выплеснуть энергию"},
		"work out":                 {"делать физические упражнения", "развиваться"},
		"work something out":       {"производить какие-то расчеты"},
		"wrap up":                  {"завершить (неформально)"},
		"write back":               {"написать обратно"},
		"write something down":     {"записать напоминалку"},
		"yell out":                 {"вскрикивать для привлечения внимания"},
	},
)

var VocabularyEngToRus *entity.Vocabulary
var VocabularyRusToEng *entity.Vocabulary
var AllVocabularies []*entity.Vocabulary

func init() {
	VocabularyEngToRus = entity.NewEmptyVocabulary().
		Update(VocabularyVerbs)
	VocabularyRusToEng = VocabularyEngToRus.MakeInvertedVocabulary()
	AllVocabularies = []*entity.Vocabulary{VocabularyEngToRus, VocabularyRusToEng}
}
