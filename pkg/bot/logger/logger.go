package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "[Jarvis-bot] - ", log.LstdFlags)
)

func Info(message string) {
	logger.Printf("[INFO] - %s", message)
}

func Usecase(usecase string) {
	message := fmt.Sprintf("SUCESSFULL CALLED USECASE: %s", usecase)
	Info(message)
}
