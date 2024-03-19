package constants

const (
	DefaultExchangeValue = "USD-BRL"
)

const (
	ImportForgotMessage              = "You forgot to import %s dependency, \nuse the %s option"
	InvalidExpenseUsecaseCharMessage = "Invalid format! \nTo request this action, you must provide at this format: \nkey-words, amount number"
	ExpenseSavedMessage              = "I just saved the amount: %s \nwith the name: %s \nin the external expenses list"
)

var (
	ExchangeCoins = []string{
		"USD-BRL",
		"EUR-BRL",
		"BTC-BRL",
	}
)
