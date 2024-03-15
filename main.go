package main

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/bot"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	if bt, err := bot.NewBotWithEnv(); err != nil {
		log.Panic(err)
	} else {
		bt.Start()
	}
}
