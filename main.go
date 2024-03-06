package main

import (
	"github.com/Joaquimborges/waitress/pkg/bot"
	"log"
	"os"
)

func main() {
	if bt, err := bot.NewBot(os.Getenv("BOT_TOKEN")); err != nil {
		log.Panic(err)
	} else {
		bt.Start()
	}
}
