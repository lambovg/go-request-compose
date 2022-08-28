package main

import (
	"errors"
	"log"
	"os"
	"testing"
)

// TODO wasCallled should be defined here
func (l *BuiltinLogger) MockBuiltinLogger() *BuiltinLogger {
	l.logger = log.New(os.Stdout, "", 5)
	return &BuiltinLogger{logger: l.logger}
}

func TestResponse(t *testing.T) {
	spy := NewBuiltinLogger().MockBuiltinLogger()

	var response Response
	response.Response(spy)

	if !spy.wasCalled.Printf {
		t.Error("Printf was not called")
	}

	if spy.wasCalled.PrintLn {
		t.Error("Println is called")
	}
}

func TestResponseError(t *testing.T) {
	var response Response
	response.Err = errors.New("error")
	response.Response(NewBuiltinLogger())
}

func TestResponseBody(t *testing.T) {
	var response Response
	response.Body = "body"
	response.Response(NewBuiltinLogger())
}
