package response

import (
	"errors"
	"github.com/lambovg/go-request-compose/pkg/logger"
	"testing"
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
	spy := logger.NewBuiltinLogger().MockBuiltinLogger()

	var response Response
	response.Err = errors.New("error")
	response.Response(spy)

	response.Response(logger.NewBuiltinLogger())

	if !spy.WasCalled.PrintLn {
		t.Error("Println is not called")
	}
}
