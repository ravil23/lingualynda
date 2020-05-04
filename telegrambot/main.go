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
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Panic(err)
		}
	}()
}

func (b *Bot) Run() {
	b.api.SendAlert("Bot is running...")
	defer func() {
		if r := recover(); r != nil {
			b.api.SendAlert(fmt.Sprintf("Recovered from panic: %s", r))
		}
	}()
	b.api.SetMessagesHandler(func(message *dao.Message) error {
		return b.api.SendNextPoll(message.User)
	})
	b.api.SetPollAnswersHandler(func(pollAnswer *dao.PollAnswer) error {
		return b.api.SendNextPoll(pollAnswer.User)
	})
	if err := b.api.ListenUpdates(); err != nil {
		log.Panic(err)
	}
}

func main() {
	bot := NewBot()
	bot.Init()
	bot.HealthCheck()
	for {
		bot.Run()
	}
}
