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

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

const (
	retryPeriod     = time.Second
	maxRetriesCount = 30
)

var helpText = strings.Join([]string{
	"<b>Vocabularies</b>",
	"/all - All words",
	fmt.Sprintf("/%s - Only words from <i>Vocabulary for IELTS Advanced - Pauline Cullen</i>", vocabularyPauline),
	fmt.Sprintf("/%s - Only phrasal verbs", vocabularyPhrasalVerbs),
	fmt.Sprintf("/%s - Only superlative adjectives", vocabularySuperlativeAdjectives),
	fmt.Sprintf("/%s - Only words about body", vocabularyBody),
	fmt.Sprintf("/%s - Only idioms", vocabularyIdioms),
	"",
	"<b>Modes</b>",
	fmt.Sprintf("/%s - Only Russian to English tasks", modeRusToEng),
	fmt.Sprintf("/%s - Only English to Russian tasks", modeEngToRus),
	fmt.Sprintf("/%s - Select random task for both side", modeRandom),
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
	b.api.SetMessagesHandler(func(message *dao.Message) error {
		var chatsState *ChatsState
		if state, found := chatsStates[message.ChatID]; found {
			chatsState = state
		} else {
			chatsState = &ChatsState{}
			chatsStates[message.ChatID] = chatsState
		}
		switch message.Text {
		case "/help", "/start":
			b.api.SendHTMLMessage(message.ChatID, helpText)
			return nil
		case "/debug":
			chatsState.debug = true
		case fmt.Sprintf("/%s", modeRandom):
			chatsState.SetMode(modeRandom)
		case fmt.Sprintf("/%s", modeEngToRus):
			chatsState.SetMode(modeEngToRus)
		case fmt.Sprintf("/%s", modeRusToEng):
			chatsState.SetMode(modeRusToEng)
		case fmt.Sprintf("/%s", vocabularyAll):
			chatsState.SetVocabulary(vocabularyAll)
		case fmt.Sprintf("/%s", vocabularyPauline):
			chatsState.SetVocabulary(vocabularyPauline)
		case fmt.Sprintf("/%s", vocabularyPhrasalVerbs):
			chatsState.SetVocabulary(vocabularyPhrasalVerbs)
		case fmt.Sprintf("/%s", vocabularySuperlativeAdjectives):
			chatsState.SetVocabulary(vocabularySuperlativeAdjectives)
		case fmt.Sprintf("/%s", vocabularyBody):
			chatsState.SetVocabulary(vocabularyBody)
		case fmt.Sprintf("/%s", vocabularyIdioms):
			chatsState.SetVocabulary(vocabularyIdioms)
		}
		b.debug(chatsState)
		return b.api.SendNextPoll(message.User)
	})
	b.api.SetPollAnswersHandler(func(pollAnswer *dao.PollAnswer) error {
		return b.api.SendNextPoll(pollAnswer.User)
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

func (b *Bot) debug(chatsState *ChatsState) {
	if chatsState.debug {
		debugMessage := fmt.Sprintf("\nSelected mode: '%s'", chatsState.mode)
		debugMessage += fmt.Sprintf("\nSelected vocabulary size: %d", len(chatsState.vocabularies))
		debugMessage += fmt.Sprintf("\nExample term and translations:")
		for _, vocabulary := range chatsState.vocabularies {
			term := vocabulary.GetRandomTerm()
			debugMessage += fmt.Sprintf("\n %s - %s", term, vocabulary.GetTranslations(term))
		}
		b.api.SendAlert(debugMessage)
	}
}
