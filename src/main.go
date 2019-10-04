package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var stdOut io.Writer = os.Stdout
var stdError io.Writer = os.Stderr

var levels = map[string]int{
	"error": 0,
	"info":  1,
	"debug": 2,
}

// Logger Simple console logger supporting log levels
type Logger struct {
	logLevel string
}

// Debug Print debug message
func (logger *Logger) Debug(message interface{}) {
	if levels[logger.logLevel] < levels["debug"] {
		return
	}
	logger.logMessage("DEBUG", message)
}

// Info Print info message
func (logger *Logger) Info(message interface{}) {
	if levels[logger.logLevel] < levels["info"] {
		return
	}
	logger.logMessage("INFO", message)
}

// Error Print error message
func (logger *Logger) Error(message interface{}) {
	logger.logMessage("ERROR", message)
}

func (logger *Logger) logMessage(level string, message interface{}) {
	out := stdOut
	if strings.ToLower(level) == "error" {
		out = stdError
	}
	_, _ = fmt.Fprintf(out, "%v - %v", level, message)
}

var loggerInstance *Logger

// GetLogger Init logger instance or get existing
// TODO: make tread-safe solution
func GetLogger(logLevel string) *Logger {
	if loggerInstance != nil {
		if loggerInstance.logLevel != logLevel {
			loggerInstance.logLevel = logLevel
		}
		return loggerInstance
	}
	loggerInstance = &Logger{logLevel}
	return loggerInstance
}
