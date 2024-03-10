## Jarvis
### Install
```bash
go get -u github.com/Joaquimborges/jarvis-bot
```

### Overview
Jarvis is a simple Telegram bot integrated with GPT from openAI.
This project arose from a personal need to automate some tasks, 
one of the main motivations was to health check some test servers and 
listen to the webhooks of the deployments that happen weekly.

### Requirements
- Go version 1.21...
- [Telegram bot token](https://core.telegram.org/bots/api)
- [OpenAI API KEY](https://platform.openai.com/api-keys)

### Envs
 - BOT_TOKEN=
 - OPEN_AI_API_KEY=
 - OPEN_AI_MODEL=

### Usage
```go
import "github.com/Joaquimborges/jarvis-bot/pkg/bot"


func main() {
   if bt, err := bot.NewBotWithEnv(); err != nil {
       //handle error 
   } else {
      bt.Start()
   }
}
```