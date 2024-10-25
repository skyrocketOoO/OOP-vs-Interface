package main

import (
	"fmt"
)

// Logger interface defines a logging behavior
type Logger interface {
	Log(message string)
}

// ConsoleLogger implements Logger for console output
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

// FileLogger implements Logger for file output
type FileLogger struct{}

func (f FileLogger) Log(message string) {
	// Write to a file (for demonstration, we'll just print)
	fmt.Println("File:", message)
}

// Application struct uses Logger interface
type Application struct {
	logger Logger
}

// NewApplication creates a new Application with a given logger
func NewApplication(logger Logger) *Application {
	return &Application{logger: logger}
}

func (app *Application) DoSomething() {
	app.logger.Log("Doing something!")
}

func main() {
	consoleLogger := ConsoleLogger{}
	app := NewApplication(consoleLogger)
	app.DoSomething()

	fileLogger := FileLogger{}
	app = NewApplication(fileLogger) // Reusing Application with different logger
	app.DoSomething()
}
