package main

import (
	"errors"
	"log"
	"os"
	"testing"
)

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

	if spy.wasCalled.Fatalln {
		t.Error("Fatalln is called")
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
