package response_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lambovg/go-request-compose/pkg/logger"
	"github.com/lambovg/go-request-compose/pkg/logger_test"
	r "github.com/lambovg/go-request-compose/pkg/response"
	test "github.com/lambovg/go-request-compose/pkg/test"
)

func TestResponse(t *testing.T) {
	spy := logger_test.MockBuiltinLogger(logger.NewBuiltinLogger())

	var response r.Response
	response.Response(spy)

	if !spy.WasCalled.Printf {
		t.Error("Printf was not called")
	}

	if spy.WasCalled.PrintLn {
		t.Error("Println is called")
	}
}

func TestResponseError(t *testing.T) {
	spy := logger_test.MockBuiltinLogger(logger.NewBuiltinLogger())

	var response r.Response
	response.Err = errors.New("error")
	response.Response(spy)

	response.Response(logger.NewBuiltinLogger())

	if !spy.WasCalled.PrintLn {
		t.Error("Println is not called")
	}
}

func TestResponseStatusCode(t *testing.T) {
	spy := logger_test.MockBuiltinLogger(logger.NewBuiltinLogger())

	response := r.Response{Body: "Ok", Err: nil, StatusCode: 200}.Response(spy)
	test.Ok(t, response.Body, "Ok")
	test.Ok(t, fmt.Sprintf("%d", response.StatusCode), fmt.Sprintf("%d", 200))
}
