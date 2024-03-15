package bot

import (
	"database/sql"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/config"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/cmd"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ia"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository/expense_calculator"
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
	return newBot(
		os.Getenv("BOT_TOKEN"),
		os.Getenv("OPEN_AI_MODEL"),
		nil,
	)
}

// NewWithDatabase Use to instantiate when you need to use
// and store data in database
func NewWithDatabase(creatDbQuery ...constants.CreateDatabaseQuery) (*Waitress, error) {
	database, er := InitDatabase()
	if er != nil {
		return nil, er
	}
	errChan := make(chan error)
	for _, query := range creatDbQuery {
		go syncDatabase(database, query, errChan)
	}

	for {
		select {
		case err := <-errChan:
			if err != nil {
				return nil, err
			}
		}
		break
	}
	return newBot(
		os.Getenv("BOT_TOKEN"),
		os.Getenv("OPEN_AI_MODEL"),
		database,
	)
}

func syncDatabase(
	db *sql.DB,
	query constants.CreateDatabaseQuery,
	erChan chan<- error,
) {
	_, err := db.Prepare(query.String())
	if err != nil {
		erChan <- err
		return
	}
	return
}

func newBot(token, openAIModel string, database *sql.DB) (*Waitress, error) {
	botConf := telebot.Settings{
		Token:     token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		ParseMode: telebot.ModeMarkdown,
	}

	bot, err := telebot.NewBot(botConf)
	if err != nil {
		return nil, err
	}

	expenseCalculatorRepository := expense_calculator.NewExpenseCalculatorRepository(database)
	instance := Waitress{
		bot: bot,
		commands: cmd.NewCommandsInstance(
			open_ia.NewOpenIAClient(openAIModel),
			usecase.NewJarvisUsecase(),
			expenseCalculatorRepository,
		),
	}
	instance.setupRoutes()
	return &instance, nil
}

func (instance *Waitress) Start() {
	config.Logger.Println("Jarvis is alive...")
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
