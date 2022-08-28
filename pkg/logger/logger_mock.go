package logger

import (
	"log"
	"os"
)

func (l *BuiltinLogger) MockBuiltinLogger() *BuiltinLogger {
	l.logger = log.New(os.Stdout, "", 5)
	return &BuiltinLogger{logger: l.logger}
}

func (l *BuiltinLogger) WasCalled() WasCalled {
	return l.wasCalled;
}