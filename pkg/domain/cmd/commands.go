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
	menu    *telebot.ReplyMarkup
	usecase *usecase.JarvisUsecase
}

func NewCommandsInstance(
	usecase *usecase.JarvisUsecase,
) *Commands {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	return &Commands{
		menu:    menu,
		usecase: usecase,
	}
}

func (cmd *Commands) Start(c telebot.Context) error {
	menu := cmd.menu
	menu.Reply(
		menu.Row(cmd.UsecaseBtn()),
		menu.Row(cmd.PingServer()),
	)
	if c.Sender().Username == os.Getenv("ADMIN_USERNAME") {
		//config.Logger.Println("start talking with JB")
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
	return c.Send(cmd.usecase.FindAndBuildUsecase(c.Text()))
}
