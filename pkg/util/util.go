package util

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	dateTimeRegexCompile = regexp.MustCompile("[0-9]+")
	ExpenseRegexCompiler = regexp.MustCompile(`(?i)gastos|gastar`)
)

func CreateLocalTime(locale string) time.Time {
	location, err := time.LoadLocation(locale)
	if err != nil {
		logger.Warn(
			"[util.CreateNewStringLocalDate()]: %v",
			err.Error(),
		)
		return time.Now().UTC()
	}
	return time.Now().In(location)
}

func BuildComparableTime(days, months int) (now time.Time, then time.Time) {
	now = time.Now()
	then = now.AddDate(0, months, days)
	return
}

func GetNumberValueFromMessage(message string) int {
	strSlice := dateTimeRegexCompile.FindAllString(message, -1)
	if len(strSlice) < 1 {
		return 0
	}
	n, err := strconv.Atoi(strSlice[0])
	if err != nil {
		return 0
	}
	return n
}

func ParseDate(date time.Time) string {
	return fmt.Sprintf("%d-%s-%d",
		date.Day(),
		date.Month().String(),
		date.Year(),
	)
}

func SliceEnvs(value string) []string {
	return strings.Split(value, "|")
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
