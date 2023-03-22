package main

import (
	"github.com/CyclopsV/check-sub-tgbot/configs"
	"github.com/CyclopsV/check-sub-tgbot/internal/bot"
)

func main() {
	config := configs.New()
	bot := bot.New(config)
	bot.Run()
}
