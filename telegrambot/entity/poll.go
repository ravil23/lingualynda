package entity

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/collection/schema"
)

type PollID string
type PollType string
type PollOptionID int64

const (
	PollTypeQuiz = "quiz"
)

func (t PollType) String() string {
	return string(t)
}

type PollOption struct {
	Translation schema.Translation
	IsCorrect   bool
}

type Poll struct {
	ID       PollID
	Type     PollType
	IsPublic bool

	Term    schema.Term
	Options []*PollOption
}

func (p *Poll) AllIsCorrect(chosenOptionIndexes []int) bool {
	for _, index := range chosenOptionIndexes {
		if !p.Options[index].IsCorrect {
			return false
		}
	}
	return true
}

func (p *Poll) ToChatable(chatID ChatID) *tgbotapi.SendPollConfig {
	correctOptionID := -1
	correctAnswer := ""
	tgOptions := make([]string, 0, len(p.Options))
	for i, option := range p.Options {
		tgOptions = append(tgOptions, option.Translation.String())
		if option.IsCorrect {
			correctOptionID = i
			correctAnswer = option.Translation.String()
		}
	}
	tgPoll := tgbotapi.NewPoll(int64(chatID), p.Term.String(), tgOptions...)
	tgPoll.CorrectOptionID = int64(correctOptionID)
	tgPoll.Type = p.Type.String()
	tgPoll.IsAnonymous = !p.IsPublic
	tgPoll.Explanation = correctAnswer
	return &tgPoll
}
