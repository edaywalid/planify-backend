package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// MyLogger is a wrapper around zerolog.Logger
type MyLogger struct {
	zerolog.Logger
}

// NewLogger initializes a new MyLogger instance
func NewLogger() MyLogger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}

	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	zlog := zerolog.New(output).With().Caller().Timestamp().Logger()

	return MyLogger{zlog}
}

func (l *MyLogger) LogInfo() *zerolog.Event {
	return l.Info()
}

func (l *MyLogger) LogError() *zerolog.Event {
	return l.Error()
}

func (l *MyLogger) LogDebug() *zerolog.Event {
	return l.Debug()
}

func (l *MyLogger) LogWarn() *zerolog.Event {
	return l.Warn()
}

func (l *MyLogger) LogFatal() *zerolog.Event {
	return l.Fatal()
}
