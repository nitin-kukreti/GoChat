package utils

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
)

func Info(msg string, args ...any) {
	infoLogger.Printf(msg, args...)
}

func Error(msg string, args ...any) {
	errorLogger.Printf(msg, args...)
}
