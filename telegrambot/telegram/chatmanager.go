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

func (m *ChatManager) GetChat(chatID entity.ChatID) (*entity.Chat, bool) {
	chat, found := m.chats[chatID]
	return chat, found
}

func (m *ChatManager) UpdateChatConfigurations(chatID entity.ChatID, text string) *entity.Chat {
	if _, found := m.chats[chatID]; !found {
		m.chats[chatID] = entity.NewChat(chatID)
	}
	chat := m.chats[chatID]
	chat.Configure(text)
	switch chat.GetVocabulary() {
	case entity.ChatVocabularyAll:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(collection.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(collection.VocabularyRusToEng)
		default:
			chat.SetVocabularies(collection.AllVocabularies...)
		}
	case entity.ChatVocabularyPauline:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(pauline.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(pauline.VocabularyRusToEng)
		default:
			chat.SetVocabularies(pauline.AllVocabularies...)
		}
	case entity.ChatVocabularyPhrasalVerbs:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(phrasalverbs.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(phrasalverbs.VocabularyRusToEng)
		default:
			chat.SetVocabularies(phrasalverbs.AllVocabularies...)
		}
	case entity.ChatVocabularySuperlativeAdjectives:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(superlativeadjectives.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(superlativeadjectives.VocabularyRusToEng)
		default:
			chat.SetVocabularies(superlativeadjectives.AllVocabularies...)
		}
	case entity.ChatVocabularyBody:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(body.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(body.VocabularyRusToEng)
		default:
			chat.SetVocabularies(body.AllVocabularies...)
		}
	case entity.ChatVocabularyIdioms:
		switch chat.GetMode() {
		case entity.ChatModeEngToRus:
			chat.SetVocabularies(idioms.VocabularyEngToRus)
		case entity.ChatModeRusToEng:
			chat.SetVocabularies(idioms.VocabularyRusToEng)
		default:
			chat.SetVocabularies(idioms.AllVocabularies...)
		}
	case entity.ChatVocabularyLesson:
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
	return chat
}
