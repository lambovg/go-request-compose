package request

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/lambovg/go-request-compose/pkg/logger"
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
func newRequest(method string, url string, body io.Reader) *Request {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Fatalln("Error creating request client: ", err)
	}

	return &Request{req}
}

// NewRequest .
func (p Params) NewRequest(url string, requestMethod string) func() *r.Response {
	var body []byte
	var err error
	var response *http.Response
	var statusCode int
	var status string
	var header http.Header

	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		if p.FormData != nil {
			response, err = http.PostForm(url, p.FormData)
		} else {
			response, err = p.Client.Do(newRequest(requestMethod, url, p.Body).AttachHeaders(&p).Request)
		}

		if err == nil {
			defer response.Body.Close()
			body, _ = io.ReadAll(response.Body)
			statusCode = response.StatusCode
			status = response.Status
			header = response.Header

			log.Println("async body", string(body))
		}
	}()

	return func() *r.Response {
		<-rc
		return r.Response{
			Body:       string(body),
			Err:        err,
			StatusCode: statusCode,
			Header:     header,
			Status:     status}.Response(logger.NewBuiltinLogger())
	}
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
