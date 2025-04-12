package util

import (
	"log"
	"os"
)

func LogError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func LogErrorExit(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
		os.Exit(1)
	}
}
