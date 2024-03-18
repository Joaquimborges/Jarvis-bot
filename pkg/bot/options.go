package bot

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ai"
	"gopkg.in/telebot.v3"
	"time"
)

type JarvisOptions func(*Jarvis) *Jarvis

// WithParseMode use it to change the bot return parse mode.
// ModeHTML is default value.
func WithParseMode(parseMode telebot.ParseMode) JarvisOptions {
	return func(j *Jarvis) *Jarvis {
		j.parseMode = parseMode
		return j
	}
}

// WithOpenAiIntegration use it if you want to integrate with
// openai ChatGPT machine, you can find models reference here:https://openai.com/pricing
func WithOpenAiIntegration(openAIModel string) JarvisOptions {
	return func(j *Jarvis) *Jarvis {
		j.openai = open_ai.NewOpenIAClient(openAIModel)
		return j
	}
}

// WithDatabase use it if you need to store something.
// sqlite3 is the default database and will be created at the root of the project.
func WithDatabase(databaseName string, creatDbQuery ...constants.CreateDatabaseQuery) JarvisOptions {
	return func(j *Jarvis) *Jarvis {
		database, er := InitDatabase(databaseName)
		if er != nil {
			j.err = er
			return j
		}
		j.database = database
		j.creatDbQuery = creatDbQuery
		return j
	}
}

/*
WithPingServerURLs It's a really cool feature that
allows you to do a health check on your test or even prod servers.
You just need a public route that only returns a valid status code.
*/
func WithPingServerURLs(urls ...string) JarvisOptions {
	return func(j *Jarvis) *Jarvis {
		j.pingUrls = urls
		return j
	}
}

func buildBotSettings(token string, parseMode telebot.ParseMode) telebot.Settings {
	return telebot.Settings{
		Token:     token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		ParseMode: parseMode,
	}
}
