package telegram

import (
	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
	"github.com/ravil23/lingualynda/telegrambot/dao"
)

var selectedVocabularies = map[dao.ChatID]*schema.Vocabulary{}
