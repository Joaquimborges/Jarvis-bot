package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const defaultExchangeValue = "USD-BRL"

type ExchangeResponse struct {
	Values ExchangeValue `json:"USDBRL"`
}

type exchange struct {
	purchaseQuote float64
	salesQuote    float64
	max           float64
	min           float64
}

type ExchangeValue struct {
	PurchaseQuote       string `json:"bid"`
	SalesQuote          string `json:"ask"`
	VariationPercentage string `json:"pctChange"`
	Max                 string `json:"high"`
	Min                 string `json:"low"`
	CreatedAt           string `json:"create_date"`
}

func (j *JarvisUsecase) GetDayQuote(coin string) string {
	if coin == "" || coin == "." || coin == "default" {
		coin = defaultExchangeValue
	}
	url := fmt.Sprintf("https://economia.awesomeapi.com.br/last/%s", coin)
	var response ExchangeResponse
	bytes, err := j.client.Get(context.Background(), url)
	if err != nil {
		return fmt.Errorf("request error: %v", err).Error()
	}

	if er := json.Unmarshal(bytes, &response); er != nil {
		return fmt.Errorf("unmarshal error: %v", er).Error()
	}

	flValue := parseToFloat(&response.Values)
	date := strings.Split(response.Values.CreatedAt, " ")
	return fmt.Sprintf(
		"Here are the quotes: \n\npurchase: `R$%.2f\n`sale: `R$%.2f\n`This is the most that has come so far: `R$%.2f\n`the minimum: `R$%.2f\n`and had a variation of: `%s%%\n`date: %s\n`time: %s\n`",
		flValue.purchaseQuote,
		flValue.salesQuote,
		flValue.max,
		flValue.min,
		response.Values.VariationPercentage,
		date[0],
		date[1],
	)
}

func parseToFloat(values *ExchangeValue) *exchange {
	pq, _ := strconv.ParseFloat(values.PurchaseQuote, 64)
	sq, _ := strconv.ParseFloat(values.SalesQuote, 64)
	maxValue, _ := strconv.ParseFloat(values.Max, 64)
	minValue, _ := strconv.ParseFloat(values.Min, 64)
	return &exchange{
		purchaseQuote: pq,
		salesQuote:    sq,
		max:           maxValue,
		min:           minValue,
	}
}
