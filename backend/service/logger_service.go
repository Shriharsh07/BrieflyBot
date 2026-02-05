package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() {

	// Ensure logs directory exists
	logDir := "logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		fmt.Println("Failed to create logs directory:", err)
		os.Exit(1)
	}

	// Create log file name based on date
	date := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("%s.log", date)
	logPath := filepath.Join(logDir, logFileName)

	// Open log file in append mode
	file, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}

	// Configure logger
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("========== BrieflyBot started ==========")
}
