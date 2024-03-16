package main

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/bot"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	bt, err := bot.NewWithDatabase(constants.ExpenseCalculatorCreateDatabaseQuery)
	if err != nil {
		log.Panic(err)
		return
	}

	if er := bt.SyncDatabase(); er != nil {
		log.Panic(er)
		return
	}
	bt.Start()
}
