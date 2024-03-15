package config

import (
	"log"
	"os"
)

var (
	Logger = log.New(os.Stdout, "[Jarvis-bot] - ", log.LstdFlags)
)
