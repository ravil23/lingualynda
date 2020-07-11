package entity

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PollAnswer struct {
	PollID        PollID
	ChosenOptions []int
}

func NewPollAnswer(tgPollAnswer *tgbotapi.PollAnswer) *PollAnswer {
	return &PollAnswer{
		PollID:        PollID(tgPollAnswer.PollID),
		ChosenOptions: tgPollAnswer.OptionIDs,
	}
}
