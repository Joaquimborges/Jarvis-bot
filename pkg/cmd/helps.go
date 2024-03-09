package cmd

import "gopkg.in/telebot.v3"

func (cmd *Commands) UsecaseBtn() telebot.Btn {
	return cmd.menu.Text("What are the use cases?")
}

func (cmd *Commands) UsecaseResponse(c telebot.Context) error {
	return c.Reply("I'm still being improved")
}

func (cmd *Commands) pingServer() telebot.Btn {
	return cmd.menu.Text("Wake up the test servers")
}
