package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetFutureWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	future := Params{Url: server.URL}.Get()
	ok(t, future().Body, "OK")
}

func TestGetFutureWithUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	future := Get(server.URL)
	ok(t, future().Body, "OK")
}

func TestGetPromiseWithUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	promise := Get(server.URL)
	promise()
}

func TestGetPromiseWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	promise := Params{Url: server.URL}.Get()
	promise()
}

func TestGetAsync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	Get(server.URL)
	Params{Url: server.URL}.Get()
	//TODO count server requests in order to make sure that both requests are async
	time.Sleep(1 * time.Second)
}

func TestGetSetClientAndOverrideTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ok(t, req.URL.String(), "/")
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	client := &http.Client{Timeout: 30 * time.Second}
	Params{Url: server.URL, Client: *client}.Get()
}

func ok(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
