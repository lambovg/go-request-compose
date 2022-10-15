package request

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	test "github.com/lambovg/go-request-compose/internal"
)

func TestGivenParams_whenPost_thenReturnRequestBody(t *testing.T) {
	server := postServer(t)
	defer server.Close()

	body := "post body"
	future := Params{Url: server.URL, Body: bytes.NewBufferString(body)}.Post()

	test.Ok(t, future().Body, body)
}

func TestGivenFormData_whenPost_thenReturnRequestBody(t *testing.T) {
	server := postServer(t)
	defer server.Close()

	formData := url.Values{
		"username": {"jonh-doe"},
	}

	future := Params{Url: server.URL, FormData: formData}.Post()

	test.Ok(t, future().Body, "username=jonh-doe")
}

func postServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		test.Ok(t, req.URL.String(), "/")

		body, _ := io.ReadAll(req.Body)

		rw.Write([]byte(body))
	}))

	return server
}
