package request_test

import (
	r "github.com/lambovg/go-request-compose/pkg/request"
	test "github.com/lambovg/go-request-compose/pkg/test"
	"testing"
)

func TestUrlBuild(t *testing.T) {
	url := r.Params{Hostname: "localhost", Protocol: "http"}.BuildUrl()
	test.Ok(t, url, "http://localhost")
}
