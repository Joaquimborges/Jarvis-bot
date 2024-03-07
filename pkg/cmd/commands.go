package cmd

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"os"
)

type WaitressCommands interface {
	Start(c telebot.Context) error
	OnTextMessage(c telebot.Context) error
	Menu(c telebot.Context) error
	//HelpBtnResponse(c telebot.Context) error
	//HelpBtnInstance() *telebot.Btn
}

type Commands struct {
	menu       *telebot.ReplyMarkup
	inlineMenu *telebot.ReplyMarkup
}

func NewCommandsInstance() *Commands {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	inlineMenu := &telebot.ReplyMarkup{}
	return &Commands{
		menu:       menu,
		inlineMenu: inlineMenu,
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
	return c.Reply("thanks")
}
