package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	test "github.com/lambovg/go-request-compose/internal"
)

func TestGivenParams_whenPut_thenReturnRequestBody(t *testing.T) {
	server := putServer(t)
	defer server.Close()

	body := "put body"
	future := Params{Url: server.URL, Body: bytes.NewBufferString(body)}.Put()

	test.Ok(t, future().Body, body)
}

func TestGivenFormData_whenPut_thenReturnRequestBody(t *testing.T) {
	server := putServer(t)
	defer server.Close()

	formData := url.Values{
		"username": {"john-doe"},
	}

	future := Params{Url: server.URL, FormData: formData}.Put()

	test.Ok(t, future().Body, "username=john-doe")
}

func TestGivenBuildParams_whenPut_thenReturnUrl(t *testing.T) {
	server := putServer(t)
	defer server.Close()

	params := Params{Hostname: "localhost", Port: 8080, Protocol: "http", Path: "/hello-world.json"}

	test.Ok(t, params.BuildUrl(), "http://localhost:8080/hello-world.json")
}

func TestGivenBuildParams_whenPut_thenReturnRequestBody(t *testing.T) {
	server := putServer(t)
	defer server.Close()

	params := Params{Hostname: "localhost", Port: 80, Protocol: "http", Path: "/"}
	params.Url = server.URL
	params.Body = bytes.NewBufferString("OK")

	future := params.Put()

	test.Ok(t, future().Body, "OK")
}

func TestGivenUrlWithBody_whenPut_thenReturnRequestBody(t *testing.T) {
	server := putServer(t)
	defer server.Close()

	future := Put(server.URL, bytes.NewBufferString("OK"))

	test.Ok(t, future().Body, "OK")
}

func TestGivenUrl_whenPut_thenReturnRequestBody(t *testing.T) {
	server := putServer(t)
	defer server.Close()

	future := Put(server.URL, nil)

	test.Ok(t, future().Body, "")
}

func TestGivenClient_whenPut_thenReturnStatusCode(t *testing.T) {
	server := putServer(t)
	defer server.Close()

	client := &http.Client{Timeout: 30 * time.Second}
	future := Params{Url: server.URL, Client: *client}.Put()

	test.Ok(t, fmt.Sprintf("%d", future().StatusCode), "200")
}

func putServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		test.Ok(t, req.URL.String(), "/")

		body, _ := io.ReadAll(req.Body)

		rw.Write([]byte(body))
	}))

	return server
}
