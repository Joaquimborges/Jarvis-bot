package cmd

import (
	"context"
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ia"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository"
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
	menu    *telebot.ReplyMarkup
	gpt     open_ia.OpenAI
	usecase *usecase.JarvisUsecase
	dB      repository.ExpenseCalculator
}

func NewCommandsInstance(
	gpt open_ia.OpenAI,
	usecase *usecase.JarvisUsecase,
	dB repository.ExpenseCalculator) *Commands {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	return &Commands{
		menu:    menu,
		gpt:     gpt,
		usecase: usecase,
		dB:      dB,
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

	//The ask prefix is necessary to identify questions that will be redirected to GPT - OpenAI
	if strings.HasPrefix(c.Message().Text, "/ask ") {
		message := strings.TrimPrefix(c.Message().Text, "/ask ")
		gptContext, err := cmd.gpt.GetMessageContext(context.Background(), message)
		if err != nil {
			return c.Send(fmt.Sprintf("Error mounting context: %v", err))
		}
		return c.Send(gptContext)
	}

	if strings.HasPrefix(c.Text(), "/exchange") {
		message := strings.TrimPrefix(c.Message().Text, "/exchange ")
		resp := cmd.usecase.GetDayQuote(message)
		//config.Logger.Println("exchange usecase was called")
		return c.Send(resp)
	}

	if strings.HasPrefix(c.Text(), "/save ") {
		message := strings.TrimPrefix(c.Text(), "/save ")

		data, err := entities.NewExpenseCalculatorBody(message)
		if err != nil {
			return c.Send(err.Error())
		}

		er := cmd.dB.Save(data)
		if er != nil {
			return c.Send(er.Error())
		}
		return c.Send("saved")
	}
	return c.Send("wait", cmd.menu)
}
