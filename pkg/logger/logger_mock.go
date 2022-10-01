package logger

import (
	"log"
	"os"
)

func MockBuiltinLogger(l *BuiltinLogger) *BuiltinLogger {
	l.Logger = log.New(os.Stdout, "", 5)
	return &BuiltinLogger{Logger: l.Logger}
}
