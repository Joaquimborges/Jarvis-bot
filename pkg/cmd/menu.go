package cmd

import "gopkg.in/telebot.v3"

func (command *commands) Menu(c telebot.Context) error {
	menu := command.menu
	menu.Reply(
		menu.Row(command.helpBtn),
	)
	return c.Send(menu)
}
