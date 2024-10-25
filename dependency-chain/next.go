package main

import (
	"fmt"
	"time"
)

// Logger interface
type Logger interface {
	Log(message string)
}

// ConsoleLogger implements Logger
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console:", formatMessage(message))
}

// FileLogger implements Logger
type FileLogger struct{}

func (f FileLogger) Log(message string) {
	fmt.Println("File:", formatMessage(message))
}

// Helper function for formatting messages
func formatMessage(message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] %s", timestamp, message)
}

func main() {
	consoleLogger := ConsoleLogger{}
	consoleLogger.Log("Doing something!")

	fileLogger := FileLogger{}
	fileLogger.Log("Doing something!")
}
