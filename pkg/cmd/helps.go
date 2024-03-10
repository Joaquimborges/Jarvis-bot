package cmd

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func (cmd *Commands) UsecaseBtn() telebot.Btn {
	return cmd.menu.Text("What are the use cases?")
}

func (cmd *Commands) UsecaseResponse(c telebot.Context) error {
	return c.Reply("I'm still being improved")
}

func (cmd *Commands) PingServer() telebot.Btn {
	return cmd.menu.Text("Wake up the test servers")
}

func (cmd *Commands) PingServersResponse(c telebot.Context) error {
	err := cmd.usecase.WakeAllTestServers()
	if err != nil {
		return c.Send(fmt.Sprintf("Request error: %v", err))
	}
	return c.Reply("The servers are ready")
}
