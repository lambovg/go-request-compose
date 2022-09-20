package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetFutureWithParams(t *testing.T) {
	server := server(t)
	defer server.Close()

	future := Params{Url: server.URL}.Get()
	ok(t, future().Body, "OK")
}

func TestGetFutureWithUrl(t *testing.T) {
	server := server(t)
	defer server.Close()

	future := Get(server.URL)
	ok(t, future().Body, "OK")
}

func TestGetPromiseWithParams(t *testing.T) {
	server := server(t)
	defer server.Close()

	promise := Params{Url: server.URL}.Get()
	ok(t, promise().Body, "OK")
}

func TestGetAsync(t *testing.T) {
	server := server(t)
	defer server.Close()

	Get(server.URL)
	Params{Url: server.URL}.Get()
	//TODO count server requests in order to make sure that both requests are async
	time.Sleep(1 * time.Second)
}

func TestGetSetClientAndOverrideTimeout(t *testing.T) {
	server := server(t)
	defer server.Close()

	client := &http.Client{Timeout: 30 * time.Second}
	Params{Url: server.URL, Client: *client}.Get()
}

func TestGetOverrideTimeoutWithCompableHttpClient(t *testing.T) {
	server := server(t)
	defer server.Close()

	client := http.Client{Timeout: 30 * time.Second}
	params := Params{Url: server.URL}

	HttpClient{client}.Get(params)
}

func TestBuildUrlByParams(t *testing.T) {
	server := server(t)
	defer server.Close()

	params, _ := Params{Hostname: "localhost", Port: 8080, Protocol: "http", Path: "/hello-world.json"}.Getv2()

	ok(t, params.Url, "http://localhost:8080/hello-world.json")
}

func ok(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func server(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	return server
}
