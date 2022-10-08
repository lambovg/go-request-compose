package request_test

import (
	test "github.com/lambovg/go-request-compose/internal"
	r "github.com/lambovg/go-request-compose/pkg/request"
	"testing"
)

func TestUrlBuild(t *testing.T) {
	url := r.Params{Hostname: "localhost", Protocol: "http"}.BuildUrl()
	test.Ok(t, url, "http://localhost")
}
