package response

import (
	"errors"
	"testing"
	"github.com/lambovg/go-request-compose/pkg/logger"
)

func TestResponse(t *testing.T) {
	spy := logger.NewBuiltinLogger().MockBuiltinLogger()

	var response Response
	response.Response(spy)

	if !spy.WasCalled.Printf {
		t.Error("Printf was not called")
	}

	if spy.WasCalled.PrintLn {
		t.Error("Println is called")
	}
}

func TestResponseError(t *testing.T) {
	var response Response
	response.Err = errors.New("error")
	response.Response(logger.NewBuiltinLogger())
}

func TestResponseBody(t *testing.T) {
	var response Response
	response.Body = "body"
	response.Response(logger.NewBuiltinLogger())
}
