package request

import (
	test "github.com/lambovg/go-request-compose/pkg/test"
	"testing"
)

func TestUrlBuild(t *testing.T) {
	url := Params{Hostname: "localhost", Protocol: "http"}.BuildUrl()
	test.Ok(t, url, "http://localhost")
}
