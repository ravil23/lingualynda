package telegram

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ravil23/lingualynda/telegrambot/collection"
	"github.com/ravil23/lingualynda/telegrambot/collection/pauline"
	"github.com/ravil23/lingualynda/telegrambot/collection/phrasalverbs"
	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const (
	retryPeriod     = time.Second
	maxRetriesCount = 30
)

const helpText = `Vocabularies:
/all - Use total vocabulary
/phrasalverbs - Use only phrasal verbs
/pauline - Use only words from book "Vocabulary for IELTS Advanced - Pauline Cullen"

Modes:
/rus2eng - Use only tasks for translation from Russian to English
/eng2rus - Use only tasks for translation from English to Russian
/random - Select random task for both side`

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
	b.api.SendAlert(botMention + " is running...")
	defer func() {
		if r := recover(); r != nil {
			b.api.SendAlert(fmt.Sprintf("Recovered from panic: %s", r))
		}
	}()
	b.api.SetMessagesHandler(func(message *dao.Message) error {
		if message.Text == "/help" {
			b.api.SendMessage(message.ChatID, helpText)
			return nil
		}
		switch message.Text {
		case "/random":
			selectedModes[message.ChatID] = modeRandom
		case "/eng2rus":
			selectedModes[message.ChatID] = modeEngToRus
		case "/rus2eng":
			selectedModes[message.ChatID] = modeRusToEng
		}
		selectedMode := selectedModes[message.ChatID]
		switch message.Text {
		case "/all":
			switch selectedMode {
			case modeEngToRus:
				selectedVocabularies[message.ChatID] = collection.VocabularyEngToRus
			case modeRusToEng:
				selectedVocabularies[message.ChatID] = collection.VocabularyRusToEng
			default:
				selectedVocabularies[message.ChatID] = collection.VocabularyTotal
			}
		case "/pauline":
			switch selectedMode {
			case modeEngToRus:
				selectedVocabularies[message.ChatID] = pauline.VocabularyEngToRus
			case modeRusToEng:
				selectedVocabularies[message.ChatID] = pauline.VocabularyRusToEng
			default:
				selectedVocabularies[message.ChatID] = pauline.VocabularyTotal
			}
		case "/phrasalverbs":
			switch selectedMode {
			case modeEngToRus:
				selectedVocabularies[message.ChatID] = phrasalverbs.VocabularyEngToRus
			case modeRusToEng:
				selectedVocabularies[message.ChatID] = phrasalverbs.VocabularyRusToEng
			default:
				selectedVocabularies[message.ChatID] = phrasalverbs.VocabularyTotal
			}
		}
		return b.api.SendNextPoll(message.User)
	})
	b.api.SetPollAnswersHandler(func(pollAnswer *dao.PollAnswer) error {
		return b.api.SendNextPoll(pollAnswer.User)
	})
	go func() {
		if err := b.api.ListenUpdates(); err != nil {
			log.Panic(err)
		}
	}()
	b.serve()
}

func (b *Bot) serve() {
	b.api.SendAlert(fmt.Sprintf("%s started", botMention))
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	b.api.SendAlert(fmt.Sprintf("%s stopped", botMention))
}
