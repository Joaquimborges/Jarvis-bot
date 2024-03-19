package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "[Jarvis-bot] - ", log.LstdFlags)
)

func Info(message string, tags ...interface{}) {
	logger.Printf("[INFO] - %s", format(message, tags...))
}

func Warn(message string, tags ...interface{}) {
	logger.Printf("[WARN] - %s", format(message, tags...))
}

func Usecase(usecase string) {
	Info("SUCCESSFUL CALLED USECASE: %s", usecase)
}

func format(message string, tags ...interface{}) string {
	return fmt.Sprintf(message, tags...)
}
