package telegram

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
	"github.com/ravil23/lingualynda/telegrambot/dao"
)

type mode string

const (
	modeRandom   = mode("random")
	modeEngToRus = mode("eng-rus")
	modeRusToEng = mode("rus-eng")
)

var selectedVocabularies = map[dao.ChatID][]*schema.Vocabulary{}
var selectedModes = map[dao.ChatID]mode{}
