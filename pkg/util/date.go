package util

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/config"
	"strings"
	"time"
)

func CreateNewStringLocalDate(locale string) string {
	location, err := time.LoadLocation(locale)
	if err != nil {
		config.Logger.Println(
			fmt.Sprintf("[util.CreateNewStringLocalDate()]: %v",
				err,
			),
		)
	}

	now := time.Now().In(location)
	return fmt.Sprintf("%d-%s-%d",
		now.Day(),
		now.Month().String(),
		now.Year(),
	)
}

func ContainsValue(message string, values []string) bool {
	for _, v := range values {
		if strings.Contains(toUpper(message), toUpper(v)) {
			return true
		}
	}
	return false
}

func toUpper(char string) string {
	return strings.ToUpper(char)
}
