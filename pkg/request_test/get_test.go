package request

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	r "github.com/lambovg/go-request-compose/pkg/request"
	test "github.com/lambovg/go-request-compose/internal"
)

func TestGetFutureWithParams(t *testing.T) {
	server := server(t)
	defer server.Close()

	future := r.Params{Url: server.URL}.Get()
	test.Ok(t, future().Body, "OK")
}

func TestGetFutureWithUrl(t *testing.T) {
	server := server(t)
	defer server.Close()

	future := r.Get(server.URL)
	test.Ok(t, future().Body, "OK")
}

func TestGetPromiseWithParams(t *testing.T) {
	server := server(t)
	defer server.Close()

	promise := r.Params{Url: server.URL}.Get()
	test.Ok(t, promise().Body, "OK")
}

func TestGetAsync(t *testing.T) {
	server := server(t)
	defer server.Close()

	r.Get(server.URL)
	r.Params{Url: server.URL}.Get()
	//TODO count server requests in order to make sure that both requests are async
	time.Sleep(1 * time.Second)
}

func TestGetSetClientAndOverrideTimeout(t *testing.T) {
	server := server(t)
	defer server.Close()

	client := &http.Client{Timeout: 30 * time.Second}
	r.Params{Url: server.URL, Client: *client}.Get()
}

func TestGetOverrideTimeoutWithCompableHttpClient(t *testing.T) {
	server := server(t)
	defer server.Close()

	client := http.Client{Timeout: 30 * time.Second}
	params := r.Params{Url: server.URL}

	r.HttpClient{client}.Get(params)
}

func TestBuildUrlByParams(t *testing.T) {
	server := server(t)
	defer server.Close()

	params, _ := r.Params{Hostname: "localhost", Port: 8080, Protocol: "http", Path: "/hello-world.json"}.Getv2()

	test.Ok(t, params.Url, "http://localhost:8080/hello-world.json")
}

func TestStatusCode(t *testing.T) {
	server := server(t)
	defer server.Close()

	client := http.Client{Timeout: 30 * time.Second}
	params := r.Params{Url: server.URL}

	future := r.HttpClient{client}.Get(params)

	test.Ok(t, fmt.Sprintf("%d", future().StatusCode), "200")
}

func TestStatus(t *testing.T) {
	server := server(t)
	defer server.Close()

	client := http.Client{Timeout: 30 * time.Second}
	params := r.Params{Url: server.URL}

	future := r.HttpClient{client}.Get(params)

	test.Ok(t, future().Status, "200 OK")
}

func server(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		test.Ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	return server
}
