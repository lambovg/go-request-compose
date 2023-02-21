package request

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	test "github.com/lambovg/go-request-compose/internal"
)

func TestGivenParams_whenDelete_thenReturnEmptyRequestBody(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	future := Params{Url: server.URL}.Delete()

	test.Ok(t, future().Body, "OK")
}

func TestGivenFormData_whenDelete_thenReturnRequestBody(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	future := Params{Url: server.URL}.Delete()

	test.Ok(t, future().Body, "OK")
}

func TestGivenBuildParams_whenDelete_thenReturnUrl(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	params := Params{Hostname: "localhost", Port: 8080, Protocol: "http", Path: "/hello-world.json"}

	test.Ok(t, params.BuildUrl(), "http://localhost:8080/hello-world.json")
}

func TestGivenBuildParams_whenDelete_thenReturnResponseBody(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	params := Params{Hostname: "localhost", Port: 80, Protocol: "http", Path: "/"}
	params.Url = server.URL

	future := params.Delete()

	test.Ok(t, future().Body, "OK")
}

func TestGivenUrl_whenDelete_thenReturnResponseBody(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	future := Delete(server.URL)

	test.Ok(t, future().Body, "OK")
}

func TestGivenClient_whenDelete_thenReturnStatusCode(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	client := &http.Client{Timeout: 30 * time.Second}
	future := Params{Url: server.URL, Client: *client}.Delete()

	test.Ok(t, fmt.Sprintf("%d", future().StatusCode), "200")
}

func deleteServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		test.Ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	return server
}
