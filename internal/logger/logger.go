package logger

import (
	"fmt"
	"log"
	"os"
	"path"
)

type LogLevel int

type AppLogger struct {
	logFile     *os.File
	logLevel    LogLevel
	innerLogger *log.Logger
}

const (
	LevelDebug LogLevel = iota // when log level is 'debug', then all messages will be printed into the log file.
	LevelInfo                  // when log level is 'info', then only 'info' messages will be printed into the log file.
)

const logFileName = "app.log"

// New - creates a new logger with a specific log level.
// appPath - where log file will be stored.
// userSetLogLevel - user-defined log level (debug or info).
func New(appPath, userSetLogLevel string) (AppLogger, error) {
	var logLevel LogLevel
	switch userSetLogLevel {
	case "debug":
		logLevel = LevelDebug
	default:
		logLevel = LevelInfo
	}

	l := AppLogger{logLevel: logLevel}

	logFilePath := path.Join(appPath, logFileName)
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return l, err
	}

	l.innerLogger = log.New(logFile, "", log.Ldate|log.Ltime)

	return l, nil
}

func (l *AppLogger) print(prefix, format string, args ...any) {
	msg := fmt.Sprintf("[%s] %s", prefix, format)
	l.innerLogger.Printf(msg, args...)
}

func (l *AppLogger) Debug(format string, args ...any) {
	if l.logLevel <= LevelDebug {
		l.print("DEBG", format, args...)
	}
}

func (l *AppLogger) Info(format string, args ...any) {
	if l.logLevel <= LevelInfo {
		l.print("INFO", format, args...)
	}
}

func (l *AppLogger) Error(format string, args ...any) {
	l.print("ERROR", format, args...)
}

func (l *AppLogger) Close() {
	//nolint:errcheck // we don't care if log file wasn't closed properly
	l.logFile.Close()
}
