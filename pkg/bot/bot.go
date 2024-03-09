package bot

import (
	"github.com/Joaquimborges/waitress/pkg/cmd"
	"github.com/Joaquimborges/waitress/pkg/open_ia"
	"gopkg.in/telebot.v3"
	"os"
	"time"
)

type Waitress struct {
	bot      *telebot.Bot
	commands *cmd.Commands
}

// NewBotWithEnv Use to instantiate when you have
// the BOT_TOKEN variable accessible in the application.
func NewBotWithEnv() (*Waitress, error) {
	return NewBot(os.Getenv("BOT_TOKEN"))
}

func NewBot(token string) (*Waitress, error) {
	botConf := telebot.Settings{
		Token:     token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		ParseMode: telebot.ModeMarkdown,
	}

	bot, err := telebot.NewBot(botConf)
	if err != nil {
		return nil, err
	}

	instance := Waitress{
		bot:      bot,
		commands: cmd.NewCommandsInstance(open_ia.NewOpenIAClient()),
	}
	instance.setupRoutes()
	return &instance, nil
}

func (instance *Waitress) Start() {
	instance.bot.Start()
}

func (instance *Waitress) setupRoutes() {
	usecaseBtn := instance.commands.UsecaseBtn()
	instance.bot.Handle("/jarvis", instance.commands.Start)
	instance.bot.Handle(&usecaseBtn, instance.commands.UsecaseResponse)
	instance.bot.Handle(telebot.OnText, instance.commands.OnTextMessage)
}
