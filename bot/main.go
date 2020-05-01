package main

import (
	"github.com/ravil23/lingualynda/bot/telegrambot"
)

func main() {
	bot := telegrambot.NewBot()
	bot.Init()
	bot.Run()
}
