package telegram

import (
	"github.com/ravil23/lingualynda/telegrambot/collection"
	"github.com/ravil23/lingualynda/telegrambot/collection/body"
	"github.com/ravil23/lingualynda/telegrambot/collection/idioms"
	"github.com/ravil23/lingualynda/telegrambot/collection/lesson"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/collection/superlativeadjectives"
	"github.com/ravil23/lingualynda/telegrambot/entity"
)

type ChatManager struct {
	chats map[entity.ChatID]*entity.Chat
}

func NewChatManager() *ChatManager {
	return &ChatManager{
		chats: make(map[entity.ChatID]*entity.Chat),
	}
}

func (m *ChatManager) GetChatOrCreate(chatID entity.ChatID) *entity.Chat {
	if chat, found := m.chats[chatID]; found {
		return chat
	} else {
		chat := entity.NewChat(chatID)
		m.chats[chatID] = chat
		return chat
	}
}

func (m *ChatManager) UpdateChatConfigurations(chatID entity.ChatID, text string) *entity.Chat {
	chat := m.GetChatOrCreate(chatID)
	chat.ConfigureFromText(text)
	m.SetupChatConfiguration(chat, chat.GetMode(), chat.GetVocabularyType())
	return chat
}

func (m *ChatManager) SetupChatConfiguration(chat *entity.Chat, mode entity.ChatMode, vocabularyType entity.ChatVocabularyType) {
	chat.Configure(chat.IsDebuggingEnabled(), mode, vocabularyType)
	switch vocabularyType {
	case entity.ChatVocabularyTypeAllTerms:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(collection.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(collection.VocabularyRusToEng)
		default:
			chat.SetVocabularies(collection.AllVocabularies...)
		}
	case entity.ChatVocabularyTypePauline:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(pauline.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(pauline.VocabularyRusToEng)
		default:
			chat.SetVocabularies(pauline.AllVocabularies...)
		}
	case entity.ChatVocabularyTypePhrasalVerbs:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(phrasalverbs.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(phrasalverbs.VocabularyRusToEng)
		default:
			chat.SetVocabularies(phrasalverbs.AllVocabularies...)
		}
	case entity.ChatVocabularyTypeSuperlativeAdjectives:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(superlativeadjectives.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(superlativeadjectives.VocabularyRusToEng)
		default:
			chat.SetVocabularies(superlativeadjectives.AllVocabularies...)
		}
	case entity.ChatVocabularyTypeBody:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(body.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(body.VocabularyRusToEng)
		default:
			chat.SetVocabularies(body.AllVocabularies...)
		}
	case entity.ChatVocabularyTypeIdioms:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(idioms.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(idioms.VocabularyRusToEng)
		default:
			chat.SetVocabularies(idioms.AllVocabularies...)
		}
	case entity.ChatVocabularyTypeLesson:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(lesson.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(lesson.VocabularyRusToEng)
		default:
			chat.SetVocabularies(lesson.AllVocabularies...)
		}
	default:
		chat.SetVocabularies(collection.AllVocabularies...)
	}
}
