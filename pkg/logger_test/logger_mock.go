package logger_test

import (
	"github.com/lambovg/go-request-compose/pkg/logger"
	"log"
	"os"
)

func MockBuiltinLogger(l *logger.BuiltinLogger) *logger.BuiltinLogger {
	l.Logger = log.New(os.Stdout, "", 5)
	return &logger.BuiltinLogger{Logger: l.Logger}
}
