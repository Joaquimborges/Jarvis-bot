package util

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"strings"
	"time"
)

func CreateNewStringLocalDate(locale string) string {
	location, err := time.LoadLocation(locale)
	if err != nil {
		logger.Info(
			"[util.CreateNewStringLocalDate()]: %v",
			err.Error(),
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
