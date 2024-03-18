package util

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"strings"
	"time"
)

func CreateLocalTime(locale string) time.Time {
	location, err := time.LoadLocation(locale)
	if err != nil {
		logger.Warn(
			"[util.CreateNewStringLocalDate()]: %v",
			err.Error(),
		)
		return time.Now()
	}
	return time.Now().In(location)
}

func ParseDate(date time.Time) string {
	return fmt.Sprintf("%d-%s-%d",
		date.Day(),
		date.Month().String(),
		date.Year(),
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
