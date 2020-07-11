package telegram

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ravil23/lingualynda/telegrambot/entity"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const (
	retryPeriod     = time.Second
	maxRetriesCount = 30
)

var helpText = strings.Join([]string{
	"<b>Vocabularies</b>",
	"/all - All words",
	fmt.Sprintf("/%s - Only words from <i>Vocabulary for IELTS Advanced - Pauline Cullen</i>", entity.ChatVocabularyPauline),
	fmt.Sprintf("/%s - Only phrasal verbs", entity.ChatVocabularyPhrasalVerbs),
	fmt.Sprintf("/%s - Only superlative adjectives", entity.ChatVocabularySuperlativeAdjectives),
	fmt.Sprintf("/%s - Only words about body", entity.ChatVocabularyBody),
	fmt.Sprintf("/%s - Only idioms", entity.ChatVocabularyIdioms),
	fmt.Sprintf("/%s - Only words from lesson", entity.ChatVocabularyLesson),
	"",
	"<b>Modes</b>",
	fmt.Sprintf("/%s - Only Russian to English tasks", entity.ChatModeRusToEng),
	fmt.Sprintf("/%s - Only English to Russian tasks", entity.ChatModeEngToRus),
	fmt.Sprintf("/%s - Select random task for both side", entity.ChatModeRandom),
	"",
	"<b>Tap to /next for getting new poll</b>",
}, "\n")

type Bot struct {
	api API
}

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) Init() {
	log.Printf("Bot is initializing...")
	conn := postgres.NewConnection()
	botToken := GetBotTokenOrPanic()
	for i := 1; i <= maxRetriesCount; i++ {
		if api, err := NewAPI(botToken, conn); err != nil {
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
	defer func() {
		if r := recover(); r != nil {
			b.api.SendAlert(fmt.Sprintf("Recovered from panic: %s", r))
		}
	}()
	b.api.SetMessagesHandler(func(message *entity.Message) error {
		if _, found := chatsStates[message.ChatID]; !found {
			chatsStates[message.ChatID] = entity.NewChat(message.ChatID)
		}
		chat := chatsStates[message.ChatID]
		switch message.Text {
		case "/help", "/start":
			b.api.SendHTMLMessage(message.ChatID, helpText)
			return nil
		case "/debug true":
			chat.SetDebug(true)
		case "/debug false":
			chat.SetDebug(false)
		case fmt.Sprintf("/%s", entity.ChatModeRandom):
			chat.SetMode(entity.ChatModeRandom)
		case fmt.Sprintf("/%s", entity.ChatModeEngToRus):
			chat.SetMode(entity.ChatModeEngToRus)
		case fmt.Sprintf("/%s", entity.ChatModeRusToEng):
			chat.SetMode(entity.ChatModeRusToEng)
		case fmt.Sprintf("/%s", entity.ChatVocabularyAll):
			chat.SetVocabulary(entity.ChatVocabularyAll)
		case fmt.Sprintf("/%s", entity.ChatVocabularyPauline):
			chat.SetVocabulary(entity.ChatVocabularyPauline)
		case fmt.Sprintf("/%s", entity.ChatVocabularyPhrasalVerbs):
			chat.SetVocabulary(entity.ChatVocabularyPhrasalVerbs)
		case fmt.Sprintf("/%s", entity.ChatVocabularySuperlativeAdjectives):
			chat.SetVocabulary(entity.ChatVocabularySuperlativeAdjectives)
		case fmt.Sprintf("/%s", entity.ChatVocabularyBody):
			chat.SetVocabulary(entity.ChatVocabularyBody)
		case fmt.Sprintf("/%s", entity.ChatVocabularyIdioms):
			chat.SetVocabulary(entity.ChatVocabularyIdioms)
		case fmt.Sprintf("/%s", entity.ChatVocabularyLesson):
			chat.SetVocabulary(entity.ChatVocabularyLesson)
		}
		b.debug(chat)
		return b.api.SendNextPoll(message.User)
	})
	b.api.SetPollAnswersHandler(func(user *entity.User, _ *entity.PollAnswer) error {
		return b.api.SendNextPoll(user)
	})
	b.serve()
}

func (b *Bot) serve() {
	go func() {
		b.api.SendAlert(fmt.Sprintf("%s started", botMention))
		if err := b.api.ListenUpdates(); err != nil {
			log.Panic(err)
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	b.api.SendAlert(fmt.Sprintf("%s stopped", botMention))
}

func (b *Bot) debug(chat *entity.Chat) {
	if chat.IsDebuggingEnabled() {
		debugMessage := fmt.Sprintf("\nChat ID: %d", chat.GetID())
		debugMessage += fmt.Sprintf("\nSelected mode: %s", chat.GetMode())
		debugMessage += fmt.Sprintf("\nSelected vocabulary type: %s", chat.GetVocabulary())
		debugMessage += fmt.Sprintf("\nSelected vocabularies count: %d", len(chat.GetVocabularies()))
		b.api.SendAlert(debugMessage)
	}
}
