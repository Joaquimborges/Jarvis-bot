package cmd

import (
	"context"
	"fmt"
	"github.com/Joaquimborges/waitress/pkg/open_ia"
	"gopkg.in/telebot.v3"
	"os"
	"strings"
)

type WaitressCommands interface {
	Start(c telebot.Context) error
	OnTextMessage(c telebot.Context) error
	Menu(c telebot.Context) error
}

type Commands struct {
	menu       *telebot.ReplyMarkup
	inlineMenu *telebot.ReplyMarkup
	gpt        open_ia.OpenAI
}

func NewCommandsInstance(gpt open_ia.OpenAI) *Commands {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	inlineMenu := &telebot.ReplyMarkup{}
	return &Commands{
		menu:       menu,
		inlineMenu: inlineMenu,
		gpt:        gpt,
	}
}

func (cmd *Commands) Start(c telebot.Context) error {
	menu := cmd.menu
	menu.Reply(
		menu.Row(cmd.UsecaseBtn()),
		menu.Row(cmd.pingServer()),
	)
	if c.Sender().Username == os.Getenv("ADMIN_USERNAME") {
		return c.Send(
			"It's always good to have you here",
			menu,
		)
	}
	return c.Send(
		fmt.Sprintf(
			"Hi, %s",
			c.Sender().
				Username,
		), menu)
}

func (cmd *Commands) OnTextMessage(c telebot.Context) error {
	if c.Message().Text == "" ||
		c.Message().Text == " " {
		return c.Send("If you want to talk, write something more complete and starting with /ask")
	}

	//The ask prefix is necessary to identify questions that will be redirected to GPT - OpenAI
	if strings.HasPrefix(c.Message().Text, "ask ") {
		message := strings.TrimPrefix(c.Message().Text, "ask ")
		gptContext, err := cmd.gpt.GetMessageContext(context.Background(), message)
		if err != nil {
			return c.Send(fmt.Sprintf("Error mounting context: %v", err))
		}
		return c.Send(gptContext)
	}
	return c.Send("wait")
}
