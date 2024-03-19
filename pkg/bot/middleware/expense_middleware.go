package middleware

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"gopkg.in/telebot.v3"
	"os"
	"strings"
)

func ExpenseValidator(next telebot.HandlerFunc) telebot.HandlerFunc {
	adminSlice := strings.Split(
		os.Getenv("ADMIN_WHITE_LIST"),
		"|",
	)
	return func(c telebot.Context) error {
		if util.ContainsValue(c.Text(), []string{"gastos", "gastar", "expense"}) {
			for _, id := range adminSlice {
				if c.Sender().Username == id {
					return next(c)
				}
			}
			return c.Reply("You don't have permission to access this feature")
		}
		return next(c)
	}
}
