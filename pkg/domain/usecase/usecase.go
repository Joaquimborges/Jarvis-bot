package usecase

import (
	"context"
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ia"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"strings"
)

type JarvisUsecase struct {
	client *rest.Client
	gpt    open_ia.OpenAI
	dB     repository.ExpenseCalculator
}

func NewJarvisUsecase(
	gpt open_ia.OpenAI,
	dB repository.ExpenseCalculator,
) *JarvisUsecase {
	return &JarvisUsecase{
		client: rest.NewRestClient(),
		gpt:    gpt,
		dB:     dB,
	}
}

func (j *JarvisUsecase) FindAndBuildUsecase(message string) string {
	switch message != "" {
	//
	case strings.HasPrefix(message, "/ask "):
		mgs := strings.TrimPrefix(message, "/ask ")
		gptContext, err := j.gpt.GetMessageContext(context.Background(), mgs)
		if err != nil {
			return fmt.Sprintf("Error mounting context: %v", err)
		}
		return gptContext
		//
	case util.ContainsValue(message, []string{"moeda hoje", "cotacao"}):
		return j.GetDayQuote("")
		//
	case util.ContainsValue(message, []string{"gastei", "comprei", "anota nos gastos", "gastos externos", "acabei de gastar"}):
		data := strings.Split(message, "/ ")
		amount := data[1]
		name := data[2]
		payload, err := entities.NewExpenseCalculatorBody(name, amount)
		if err != nil {
			return fmt.Sprintf("[usecase.NewExpenseCalculatorBody]Error was fount: %v", err)
		}

		if er := j.dB.Save(payload); er != nil {
			return fmt.Sprintf("[usecase.SaveExpense()]Error was fount: %v", err)
		}
		return fmt.Sprintf("I just saved the amount %s with the name %s in the external expenses list",
			amount,
			name,
		)
	//
	case util.ContainsValue(message, []string{"ver gastos"}):
		resp, err := j.dB.Select(constants.GetAllExpense.String())
		if err != nil {
			return fmt.Sprintf("[usecase.Select.all]Error was fount: %v", err)
		}
		finalMsg := ""
		for _, expense := range resp {
			finalMsg += fmt.Sprintf(
				"Description: %s\nAmount: R$%.2f\nDate: %s\n\n",
				expense.Name,
				expense.Amount,
				expense.Date,
			)
		}
		return finalMsg
	}
	return ""
}
