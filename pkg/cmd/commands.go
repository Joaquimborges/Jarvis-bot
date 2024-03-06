package cmd

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"os"
)

type WaitressCommands interface {
	Start(c telebot.Context) error
	//Menu(c telebot.Context) error
	//HelpBtnResponse(c telebot.Context) error
	//OnTextMessage(c telebot.Context) error
	//HelpBtnInstance() *telebot.Btn
}

type commands struct {
	menu       *telebot.ReplyMarkup
	inlineMenu *telebot.ReplyMarkup
	helpBtn    telebot.Btn
}

func NewCommandsInstance() WaitressCommands {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	inlineMenu := &telebot.ReplyMarkup{}
	return &commands{
		menu:       menu,
		inlineMenu: inlineMenu,
		helpBtn:    menu.Text("Quais são os casos de uso?"),
	}
}

func (command *commands) Start(c telebot.Context) error {
	menu := command.menu
	menu.Reply(
		menu.Row(command.helpBtn),
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
