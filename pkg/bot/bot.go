package bot

import (
	"github.com/Joaquimborges/waitress/pkg/cmd"
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

type Waitress struct {
	bot      *telebot.Bot
	commands cmd.WaitressCommands
}

func NewBot(token string) *Waitress {
	botConf := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(botConf)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	instance := Waitress{
		bot:      bot,
		commands: cmd.NewCommandsInstance(),
	}
	instance.setupRoutes()
	return &instance
}

func (instance *Waitress) Start() {
	instance.bot.Start()
}

func (instance *Waitress) setupRoutes() {
	instance.bot.Handle("/jarvis", instance.commands.Start)
	//instance.Bot.Handle("/menu", instance.Commands.Menu)
	//instance.Bot.Handle(telebot.OnText, server.Commands.OnTextMessage)
}
