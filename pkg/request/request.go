package request

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"

	r "github.com/lambovg/go-request-compose/pkg/response"
)

// Params .
type Params struct {
	Url         string
	Hostname    string
	Port        int
	Protocol    string
	Path        string
	QueryString string
	Body        io.Reader
	FormData    url.Values
	Headers     Headers
	Client      http.Client
}

// Request .
type Request struct {
	*http.Request
}

// HttpClient .
type HttpClient struct {
	*http.Client
}

// requestFunc .
type requestFunc func(string) func() *r.Response

// Client .
func Client(p Params) *Params {
	return &p
}

// NewRequest .
func NewRequest(method string, url string, body io.Reader) *Request {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Fatalln("Error creating request client: ", err)
	}

	return &Request{req}
}

// FutureGroup .
func FutureGroup(fn []string, rq requestFunc) {
	errorChan := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup

	for i := range fn {
		url := fn[i]

		wg.Add(1)

		go func() {
			defer wg.Done()
			res := rq(url)()

			if res.Err != nil {
				errorChan <- res.Err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-errorChan:
		close(errorChan)
		log.Println("Error encountered: ", err)
	}
}
