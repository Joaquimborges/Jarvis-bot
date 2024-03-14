package bot

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/cmd"
	"github.com/Joaquimborges/jarvis-bot/pkg/open_ia"
	"github.com/Joaquimborges/jarvis-bot/pkg/usecase"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

type Waitress struct {
	bot      *telebot.Bot
	commands *cmd.Commands
	logger   *log.Logger
}

// NewBotWithEnv Use to instantiate when you have
// the BOT_TOKEN variable accessible in the application.
func NewBotWithEnv() (*Waitress, error) {
	return NewBot(
		os.Getenv("BOT_TOKEN"),
		os.Getenv("OPEN_AI_MODEL"),
	)
}

func NewBot(token, openAIModel string) (*Waitress, error) {
	botConf := telebot.Settings{
		Token:     token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		ParseMode: telebot.ModeMarkdown,
	}

	bot, err := telebot.NewBot(botConf)
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "[Jarvis-bot] - ", log.LstdFlags)
	instance := Waitress{
		bot: bot,
		commands: cmd.NewCommandsInstance(
			open_ia.NewOpenIAClient(openAIModel),
			usecase.NewJarvisUsecase(),
			logger,
		),
		logger: logger,
	}
	instance.setupRoutes()
	return &instance, nil
}

func (instance *Waitress) Start() {
	instance.logger.Println("Jarvis is alive...")
	instance.bot.Start()
}

func (instance *Waitress) setupRoutes() {
	usecaseBtn := instance.commands.UsecaseBtn()
	wakeServerBtn := instance.commands.PingServer()

	instance.bot.Handle("/jarvis", instance.commands.Start)
	instance.bot.Handle(&usecaseBtn, instance.commands.UsecaseResponse)
	instance.bot.Handle(&wakeServerBtn, instance.commands.PingServersResponse)
	instance.bot.Handle(telebot.OnText, instance.commands.OnTextMessage)
}
