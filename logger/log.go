package logger

import (
	"log"
	"os"
)

// LogLevel represents the level of logging.
type LogLevel int

const (
	// DEBUG level logs everything.
	DEBUG LogLevel = iota
	// INFO level logs Info, Warnings and Errors.
	INFO
	// WARNING level logs Warning and Errors.
	WARNING
	// ERROR level logs only Errors.
	ERROR
)

var (
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger

	currentLogLevel LogLevel
)

func init() {
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	SetLogLevel(INFO)
}

// SetLogLevel sets the current logging level.
func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

// Debug logs debug messages when the current logging level is DEBUG.
func Debug(format string, v ...interface{}) {
	if currentLogLevel <= DEBUG {
		debugLogger.Printf(format, v...)
	}
}

// Info logs info messages when the current logging level is DEBUG or INFO.
func Info(format string, v ...interface{}) {
	if currentLogLevel <= INFO {
		infoLogger.Printf(format, v...)
	}
}

// Warning logs warning messages when the current logging level is DEBUG, INFO or WARNING.
func Warning(format string, v ...interface{}) {
	if currentLogLevel <= WARNING {
		warningLogger.Printf(format, v...)
	}
}

// Error logs error messages in all logging levels.
func Error(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}
