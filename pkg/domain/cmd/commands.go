package cmd

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase"
	"gopkg.in/telebot.v3"
	"os"
)

type WaitressCommands interface {
	Start(c telebot.Context) error
	OnTextMessage(c telebot.Context) error
	Menu(c telebot.Context) error
}

type Commands struct {
	usecase usecase.UCBuilder
}

func NewCommandsInstance(
	usecase usecase.UCBuilder,
) *Commands {
	return &Commands{
		usecase: usecase,
	}
}

func (cmd *Commands) Start(c telebot.Context) error {
	if c.Sender().Username == os.Getenv("ADMIN_USERNAME") {
		return c.Send(
			"It's always good to have you here",
		)
	}
	return c.Send(
		fmt.Sprintf(
			"Hi, %s",
			c.Sender().
				Username,
		))
}

func (cmd *Commands) OnTextMessage(c telebot.Context) error {
	if c.Message().Text == "" ||
		c.Message().Text == " " {
		return c.Send("If you want to talk, write something more complete and starting with /ask")
	}
	context := cmd.usecase.BuildResponseContext(c.Text())
	return c.Send(context)
}
