package main

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/bot"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/telebot.v3"
	"log"
)

func main() {
	bt, err := bot.NewJarvisBot(
		bot.WithParseMode(telebot.ModeHTML),
		bot.WithDatabase(
			"jarvis.db",
			constants.ExpenseCalculatorCreateDatabaseQuery,
		),
	)
	if err != nil {
		log.Panic(err)
		return
	}
	bt.Start()
}
