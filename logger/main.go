package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const dateFormat = "2006-01-02T15:04:05Z"

var stdOut io.Writer = os.Stdout
var stdError io.Writer = os.Stderr
var now = time.Now
var getPid = os.Getpid
var loggerInstance *Logger

// Logger Simple console logger supporting log levels
type Logger struct {
	options LoggingOptions
}

// LoggingOptions options for logger initialization
type LoggingOptions struct {
	LogLevel   string
	ServiceName string
}

// GetLogger Init logger instance or get existing
func GetLogger(options LoggingOptions) *Logger {
	if loggerInstance != nil {
		return loggerInstance
	}
	loggerInstance = &Logger{options}
	return loggerInstance
}

// Fatal Print fatal message
func (logger *Logger) Fatal(message interface{}) {
	logger.logMessage(fatal, message)
}

// Error Print error message
func (logger *Logger) Error(message interface{}) {
	logger.logMessage(errorLevel, message)
}

// Warn Print warning message
func (logger *Logger) Warn(message interface{}) {
	logger.logMessage(warn, message)
}

// Info Print info message
func (logger *Logger) Info(message interface{}) {
	logger.logMessage(info, message)
}

// Debug Print debug message
func (logger *Logger) Debug(message interface{}) {
	logger.logMessage(debug, message)
}

func (logger *Logger) logMessage(level string, message interface{}) {
	messageSeverity := levels[level]
	minSeverity := levels[logger.options.LogLevel]
	if messageSeverity < minSeverity {
		return
	}
	out := stdOut
	if messageSeverity >= levels[warn] {
		out = stdError
	}
	formattedMessage, err :=logger.createLogMessage(level, message)
	if err != nil || formattedMessage == "" {
		fmt.Printf("error occurred while trying to create log message: %v", err)
		return
	}
	_, _ = fmt.Fprintln(out, formattedMessage)
}

func (logger *Logger) createLogMessage(level string, message interface{}) (string, error) {
	log := logMessage{
		Message:     message,
		Type:        level,
		ProcessId:   getPid(),
		Date:        now().Format(dateFormat),
		ServiceName: logger.options.ServiceName,
	}
	messageBytes, err := json.Marshal(log)
	if err != nil {
		return "", err
	}
	return string(messageBytes), nil
}
