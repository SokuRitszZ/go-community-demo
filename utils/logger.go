package utils

import (
	"log"
	"os"
	"time"
)

var Logger *log.Logger

func InitLogger() error {
	timeStr := time.Now().Format("2006-01-01")
	filePath := "logs/log-" + timeStr
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		return err
	}
	log.SetOutput(file)
	return nil
}
