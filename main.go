package main

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/bot"
	"log"
)

func main() {
	if bt, err := bot.NewBotWithEnv(); err != nil {
		log.Panic(err)
	} else {
		bt.Start()
	}
}
