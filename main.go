package main

import (
	"github.com/Joaquimborges/waitress/pkg/bot"
	"log"
)

func main() {
	if bt, err := bot.NewBotWithEnv(); err != nil {
		log.Panic(err)
	} else {
		log.Println("Bot running")
		bt.Start()
	}
}
