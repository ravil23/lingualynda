package entity

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageID int

type Message struct {
	ID        MessageID
	ChatID    ChatID
	Timestamp time.Time
	Text      string

	User *User
}

func NewMessage(tgMessage *tgbotapi.Message, user *User) *Message {
	return &Message{
		ID:        MessageID(tgMessage.MessageID),
		ChatID:    ChatID(tgMessage.Chat.ID),
		Timestamp: tgMessage.Time(),
		User:      user,
		Text:      tgMessage.Text,
	}
}
