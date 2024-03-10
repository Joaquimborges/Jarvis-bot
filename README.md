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
 - export BOT_TOKEN=
 - export OPEN_AI_API_KEY=
 - export OPEN_AI_MODEL=
 - export ADMIN_USERNAME=

### Usage
```go
import "github.com/Joaquimborges/jarvis-bot/pkg/bot"


func main() {
   if jarvis, err := bot.NewBotWithEnv(); err != nil {
       //handle error 
   } else {
      log.Println("Bot running")
      jarvis.Start()
   }
}
```
### execute
```bash
make run
```
### then
![Screenshot 2024-03-10 at 12.17.17 AM.png](..%2F..%2FDesktop%2FScreenshot%202024-03-10%20at%2012.17.17%E2%80%AFAM.png)

### Commands
You can write any word or start with the command below:
```sh
 /jarvis
```
<img src="" alt="drawing" width="200"/>
