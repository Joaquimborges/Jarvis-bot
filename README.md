## Jarvis
### Install
```bash
go get -u github.com/Joaquimborges/jarvis-bot@latest
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
 - export TIME_LOCATION=
 - export CHAT_ID=

### Usage
```go
import "github.com/Joaquimborges/jarvis-bot/pkg/bot"


func main() {
   jarvis, err := bot.NewJarvisBot(
           bot.WithParseMode(telebot.ModeHTML),
           bot.WithDatabase("foo.db", 
		 "create db statement",
                 "other create db statement"...
	   )
   )
   if err != nil {
	//handle error
   }
   jarvis.Start()
}
```
### execute
```bash
make run
```
### then
<img width="355" alt="Screenshot 2024-03-10 at 12 17 17 AM" src="https://github.com/Joaquimborges/Jarvis-bot/assets/57245781/11ae307e-3e3d-42d2-9558-50a8bd01cd8f">

### Commands
You can write any word or start with the command below:
```sh
 /jarvis
```
<img src="https://github.com/Joaquimborges/Jarvis-bot/assets/57245781/571de568-e2c3-4615-a8d4-334b1fb12707" alt="drawing" width="300"/>

If you need to ask a direct question to the GPT chat, just write your question with the suffix /ask.
```bash
/ask Who is the best in the NBA?
```
<img src="https://github.com/Joaquimborges/Jarvis-bot/assets/57245781/27887e95-759d-4eed-a1e3-b8779b299fe7" alt="drawing" width="300"/>

By the way, there are a few names missing from this list, lol.

