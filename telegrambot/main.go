package main

import (
	"log"

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
	"github.com/ravil23/lingualynda/telegrambot/telegram"
)

type Bot struct {
	api telegram.API
}

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) Init() {
	log.Printf("Bot is initializing...")
	conn := postgres.NewConnection()
	if api, err := telegram.NewAPI(conn); err != nil {
		log.Panic(err)
	} else {
		b.api = api
	}
}

func (b *Bot) Run() {
	log.Print("Bot is running...")
	handlerFunc := func(message *dao.Message) error {
		return b.api.Reply(message, message.Text)
	}
	if err := b.api.ListenMessages(handlerFunc); err != nil {
		log.Panic(err)
	}
}

func main() {
	bot := NewBot()
	bot.Init()
	bot.Run()
}
