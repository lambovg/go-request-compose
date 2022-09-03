package logger

import (
	"log"
	"os"
)

// WasCalled FIXME this should be part from the unit test
type WasCalled struct {
	PrintLn bool
	Printf  bool
	Fatalln bool
}

type BuiltinLogger struct {
	logger    *log.Logger
	WasCalled WasCalled
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *BuiltinLogger) Println(args ...interface{}) {
	//FIXME this should be part from the unit test
	l.WasCalled.PrintLn = true
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Printf(format string, args ...interface{}) {
	//FIXME this should be part from the unit test
	l.WasCalled.Printf = true
	l.logger.Printf(format, args...)
}

func (l *BuiltinLogger) Fatalln(error error) {
	//FIXME this should be part from the unit test
	l.WasCalled.Fatalln = true
	l.logger.Fatalln(error)
}
