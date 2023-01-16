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

func TestGivenParams_whenDelete_thenReturnRequestBody(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	body := "post body"
	future := Params{Url: server.URL, Body: bytes.NewBufferString(body)}.Post()

	test.Ok(t, future().Body, body)
}

func TestGivenFormData_whenDelete_thenReturnRequestBody(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	formData := url.Values{
		"username": {"john-doe"},
	}

	future := Params{Url: server.URL, FormData: formData}.Delete()

	test.Ok(t, future().Body, "username=john-doe")
}

func TestGivenBuildParams_whenDelete_thenReturnUrl(t *testing.T) {
	server := deleteServer(t)
	defer server.Close()

	params := Params{Hostname: "localhost", Port: 8080, Protocol: "http", Path: "/hello-world.json"}

	test.Ok(t, params.BuildUrl(), "http://localhost:8080/hello-world.json")
}

// TODO: support for sending body and remove getServer
func TestGivenBuildParams_whenDelete_thenReturnRequestBody(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	params := Params{Hostname: "localhost", Port: 80, Protocol: "http", Path: "/"}
	params.Url = server.URL

	future := params.Delete()

	test.Ok(t, future().Body, "OK")
}

// TODO: support for sending body and remove getServer
func TestGivenUrl_whenDelete_thenReturnRequestBody(t *testing.T) {
	server := getServer(t)
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

		body, _ := io.ReadAll(req.Body)

		_, err := rw.Write([]byte(body))
		if err != nil {
			return
		}
	}))

	return server
}
