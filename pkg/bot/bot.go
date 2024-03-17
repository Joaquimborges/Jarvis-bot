package bot

import (
	"database/sql"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/config"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/cmd"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ai"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
	"gopkg.in/telebot.v3"
	"os"
)

type Jarvis struct {
	bot          *telebot.Bot
	commands     *cmd.Commands
	database     *sql.DB
	creatDbQuery []constants.CreateDatabaseQuery
	err          error
	openai       open_ai.OpenAI
	parseMode    telebot.ParseMode
	pingUrls     []string
}

func NewJarvisBot(options ...JarvisOptions) (*Jarvis, error) {
	var params Jarvis
	for _, opt := range options {
		opt(&params)
	}

	if params.err != nil {
		return nil, params.err
	}

	if params.parseMode == "" {
		params.parseMode = telebot.ModeHTML
	}

	bot, err := telebot.NewBot(
		buildBotSettings(
			os.Getenv("BOT_TOKEN"),
			params.parseMode,
		),
	)
	if err != nil {
		return nil, err
	}

	if params.database != nil {
		if er := params.syncDatabase(); er != nil {
			return nil, er
		}
	}

	uc := usecase.NewJarvisUsecase(
		params.openai,
		params.database,
		rest.NewRestClient(),
		params.pingUrls...,
	)
	params.bot = bot
	params.commands = cmd.NewCommandsInstance(uc)
	params.setupRoutes()
	return &params, nil
}

func (instance *Jarvis) Start() {
	config.Logger.Println("Jarvis is alive...")
	instance.bot.Start()
}

func (instance *Jarvis) syncDatabase() error {
	for _, query := range instance.creatDbQuery {
		if _, er := instance.database.Exec(string(query)); er != nil {
			return er
		}
	}
	return nil
}

func (instance *Jarvis) setupRoutes() {
	instance.bot.Handle("/jarvis", instance.commands.Start)
	instance.bot.Handle(telebot.OnText, instance.commands.OnTextMessage)
}
