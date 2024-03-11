package cmd

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"os"
)

func (cmd *Commands) UsecaseBtn() telebot.Btn {
	return cmd.menu.Text("What are the use cases?")
}

func (cmd *Commands) UsecaseResponse(c telebot.Context) error {
	return c.Reply("If you want to talk, write something more complete and starting with /ask ...")
}

func (cmd *Commands) PingServer() telebot.Btn {
	return cmd.menu.Text("Wake up the test servers")
}

func (cmd *Commands) PingServersResponse(c telebot.Context) error {
	err := cmd.usecase.WakeAllTestServers(
		os.Getenv("MACHINE_SOCKET_SERVER_URL"),
		os.Getenv("MACHINE_API_URL"),
		os.Getenv("BOSS_YM_API_URL"),
	)

	if err != nil {
		return c.Send(fmt.Sprintf("Request error: %v", err))
	}
	return c.Reply("The servers are ready")
}
