package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const timeout = 10

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) Init() {
	log.Printf("Bot is initializing...")
	var err error
	b.api, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", b.api.Self.UserName)
}

func (b *Bot) Run() {
	log.Print("Bot is running...")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err = b.api.Send(msg)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
	}
}

func main() {
	bot := NewBot()
	bot.Init()
	bot.Run()
}
