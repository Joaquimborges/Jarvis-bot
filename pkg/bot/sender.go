package bot

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
	"gopkg.in/telebot.v3"
	"os"
)

const sendMessageURL = "https://api.telegram.org/bot%s/sendMessage"

type sendMessageBody struct {
	ChatID              string `json:"chat_id"`
	Text                string `json:"text"`
	ParseMode           string `json:"parse_mode"`
	DisableNotification bool   `json:"disable_notification"`
}

func SendMessage(message string, disableNotification bool) error {
	token := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")
	url := fmt.Sprintf(sendMessageURL, token)
	body := sendMessageBody{
		ChatID:              chatID,
		Text:                message,
		ParseMode:           telebot.ModeHTML,
		DisableNotification: disableNotification,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := rest.NewRestClient()
	return client.Post(url, &body, headers)
}
