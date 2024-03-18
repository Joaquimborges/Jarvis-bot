package exchange

import (
	"encoding/json"
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"strconv"
	"strings"
)

type exchangeResponse struct {
	Values Value `json:"USDBRL"`
}

type exchange struct {
	purchaseQuote float64
	salesQuote    float64
	max           float64
	min           float64
}

type Value struct {
	PurchaseQuote       string `json:"bid"`
	SalesQuote          string `json:"ask"`
	VariationPercentage string `json:"pctChange"`
	Max                 string `json:"high"`
	Min                 string `json:"low"`
	CreatedAt           string `json:"create_date"`
}

type Exchange struct {
	client rest.Waitress
}

func NewExchangeUsecase(client rest.Waitress) *Exchange {
	return &Exchange{
		client: client,
	}
}

func (*Exchange) IsValid(message string) bool {
	return util.ContainsValue(
		message,
		[]string{"moeda hoje", "cotacao"},
	)
}

func (e *Exchange) BuildResponse(message, _ string) string {
	coin := e.getMessageCoin(message)
	url := fmt.Sprintf("https://economia.awesomeapi.com.br/last/%s", coin)

	bytes, err := e.client.Get(url)
	if err != nil {
		return fmt.Errorf("request error: %v", err).Error()
	}

	var response exchangeResponse
	if er := json.Unmarshal(bytes, &response); er != nil {
		return fmt.Errorf("unmarshal error: %v", er).Error()
	}

	flValue := e.parseToFloat(&response.Values)
	date := strings.Split(response.Values.CreatedAt, " ")
	logger.Usecase("Exchange")
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

func (*Exchange) getMessageCoin(message string) string {
	if message == "" || message == "." || contains(message, "default") {
		return constants.DefaultExchangeValue
	}

	for _, coin := range constants.ExchangeCoins {
		if contains(strings.ToUpper(message), coin) {
			return coin
		}
	}
	return constants.DefaultExchangeValue
}

func contains(message, value string) bool {
	return strings.Contains(message, value)
}

func (*Exchange) parseToFloat(values *Value) *exchange {
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
