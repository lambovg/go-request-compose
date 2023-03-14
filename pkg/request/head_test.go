package request

import (
	"fmt"
	test "github.com/lambovg/go-request-compose/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGivenParams_whenHead_thenReturnRequestBody(t *testing.T) {
	server := headServer(t)
	defer server.Close()

	future := Params{Url: server.URL}.Head()

	test.Ok(t, future().Body, "")
	test.Equals(t, future().StatusCode, 200)
	test.Ok(t, future().Status, "200 OK")
	test.Ok(t, future().Header.Get("Content-Length"), fmt.Sprintf("%v", 20))
}

func TestGivenUrl_whenHead_thenReturnRequestBody(t *testing.T) {
	server := headServer(t)
	defer server.Close()

	future := Head(server.URL)

	test.Ok(t, future().Body, "")
	test.Equals(t, future().StatusCode, 200)
	test.Ok(t, future().Status, "200 OK")
	test.Ok(t, future().Header.Get("Content-Length"), fmt.Sprintf("%v", 20))
}

func headServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Length", "20")
		test.Ok(t, req.URL.String(), "/")
	}))

	return server
}
