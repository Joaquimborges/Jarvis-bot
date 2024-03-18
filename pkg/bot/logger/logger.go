package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "[Jarvis-bot] - ", log.LstdFlags)
)

func Info(message string, tags ...string) {
	msg := fmt.Sprintf(message, tags)
	logger.Printf("[INFO] - %s", msg)
}

func Warn(message string, tags ...string) {
	msg := fmt.Sprintf(message, tags)
	logger.Printf("[WARN] - %s", msg)
}

func Usecase(usecase string) {
	message := fmt.Sprintf("SUCESSFULL CALLED USECASE: %s", usecase)
	Info(message)
}
