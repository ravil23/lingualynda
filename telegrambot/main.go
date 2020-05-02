package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
	"github.com/ravil23/lingualynda/telegrambot/telegram"
)

const (
	retryPeriod     = time.Second
	maxRetriesCount = 30
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
	botToken := telegram.GetBotTokenOrPanic()
	for i := 1; i <= maxRetriesCount; i++ {
		if api, err := telegram.NewAPI(botToken, conn); err != nil {
			log.Printf("Attempt %d failed: %v", i, err)
			time.Sleep(retryPeriod)
		} else {
			b.api = api
			log.Printf("Bot successfully initialized")
			return
		}
	}
	log.Panic("max retries count exceeded")
}

func (b *Bot) HealthCheck() {
	go func() {
		address := ":8080"
		path := "/healthcheck"
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s request to %s%s with User-Agent: %s", r.Method, r.Host, r.URL, r.UserAgent())
			_, _ = fmt.Fprint(w, `{"status": "ok"}`)
		})
		log.Printf("Listening health check on address %s%s", address, path)
		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Panic(err)
		}
	}()
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
	bot.HealthCheck()
	bot.Run()
}
