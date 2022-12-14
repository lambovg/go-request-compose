package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	test "github.com/lambovg/go-request-compose/internal"
)

func TestGetFutureWithParams(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	future := Params{Url: server.URL}.Get()
	test.Ok(t, future().Body, "OK")
}

func TestGetFutureWithUrl(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	future := Get(server.URL)
	test.Ok(t, future().Body, "OK")
}

func TestGetPromiseWithParams(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	promise := Params{Url: server.URL}.Get()
	test.Ok(t, promise().Body, "OK")
}

func TestGetAsync(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	Get(server.URL)
	Params{Url: server.URL}.Get()
	//TODO count server requests in order to make sure that both requests are async
	time.Sleep(1 * time.Second)
}

func TestGetSetClientAndOverrideTimeout(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	client := &http.Client{Timeout: 30 * time.Second}
	Params{Url: server.URL, Client: *client}.Get()
}

func TestGetOverrideTimeoutWithCompableHttpClient(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	client := http.Client{Timeout: 30 * time.Second}
	params := Params{Url: server.URL}

	HttpClient{&client}.Get(params)
}

func TestBuildUrlByParams(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	params, _ := Params{Hostname: "localhost", Port: 8080, Protocol: "http", Path: "/hello-world.json"}.Getv2()

	test.Ok(t, params.Url, "http://localhost:8080/hello-world.json")
}

func TestStatusCode(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	client := http.Client{Timeout: 30 * time.Second}
	params := Params{Url: server.URL}

	future := HttpClient{&client}.Get(params)

	test.Ok(t, fmt.Sprintf("%d", future().StatusCode), "200")
}

func TestStatus(t *testing.T) {
	server := getServer(t)
	defer server.Close()

	client := http.Client{Timeout: 30 * time.Second}
	params := Params{Url: server.URL}

	future := HttpClient{&client}.Get(params)

	test.Ok(t, future().Status, "200 OK")
}

func TestWithHttpTransport_returnsOkBody(t *testing.T) {
	client := client(t)

	future := HttpClient{client}.Get(Params{Url: "/"})
	test.Ok(t, future().Body, "OK")
}

func TestWithHttpTransport_returnsOkStatus(t *testing.T) {
	client := client(t)

	future := HttpClient{client}.Get(Params{Url: "/"})
	test.Ok(t, future().Status, "200 OK")
}

func TestWithHttpTransport_returns200StatusCode(t *testing.T) {
	client := client(t)

	future := HttpClient{client}.Get(Params{Url: "/"})
	test.Ok(t, fmt.Sprintf("%d", future().StatusCode), "200")
}

func getServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		test.Ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	return server
}

func client(t *testing.T) *http.Client {
	client := NewTestClient(func(req *http.Request) *http.Response {
		test.Ok(t, req.URL.String(), "/")
		return &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	return client
}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
