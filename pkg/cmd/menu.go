package cmd

import "gopkg.in/telebot.v3"

func (cmd *Commands) Menu(c telebot.Context) error {
	menu := cmd.menu
	menu.Reply(
		menu.Row(cmd.UsecaseBtn()),
		menu.Row(cmd.pingServer()),
	)
	return c.Send(menu)
}
