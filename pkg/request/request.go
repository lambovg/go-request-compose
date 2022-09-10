package request

import (
	cresponse "github.com/lambovg/go-request-compose/pkg/response"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Params
type Params struct {
	Url      string
	Hostname string
	Protocol string
	Path     string
	Headers  Headers
	Client   http.Client
}

// Request
type Request struct {
	*http.Request
}

// Headers
type Headers struct {
	Add Header
	Set Header
}

// Header
type Header map[string][]string

type requestFunc func(string) func() *cresponse.Response

// Client
func Client(p Params) Params {
	return p
}

// NewRequest
func NewRequest(method string, url string, body io.Reader) *Request {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Fatalln("Error creating request client: ", err)
	}

	return &Request{req}
}

// AttachHeaders
func (rq Request) AttachHeaders(p *Params) *Request {
	// set / override existing
	for key, val := range p.Headers.Set {
		rq.Header.Set(key, strings.Join(val, ","))
	}

	// add / extend definition of existing
	for key, val := range p.Headers.Add {
		rq.Header.Add(key, strings.Join(val, ","))
	}

	return &rq
}

// FutureGroup
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
