package request

import (
	"io"
	"log"
	"net/http"

	"github.com/lambovg/go-request-compose/pkg/logger"
	r "github.com/lambovg/go-request-compose/pkg/response"
)

// Get request with params builds the @link{Params} struct
func (p Params) Get() func() *r.Response {

	if p.Url != "" {
		return p.get(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.get(p.Url)
}

func (p Params) Getv2() (Params, func() *r.Response) {
	if p.Url != "" {
		return p, p.get(p.Url)
	}

	p.Url = p.BuildUrl()

	return p, p.get(p.Url)
}

// Get with HttpClient struct
func (c HttpClient) Get(p Params) func() *r.Response {
	p.Client = c.Client
	return p.get(p.Url)
}

// Get By url
func Get(url string) func() *r.Response {
	return Params{Url: url}.get(url)
}

// get .
func (p Params) get(url string) func() *r.Response {
	var body []byte
	var err error
	var statusCode int
	var status string

	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		response, err := p.Client.Do(NewRequest(http.MethodGet, url, nil).AttachHeaders(&p).Request)

		if err == nil {
			defer response.Body.Close()
			body, _ = io.ReadAll(response.Body)
			statusCode = response.StatusCode
			status = response.Status

			log.Println("async body", string(body))
		}
	}()

	return func() *r.Response {
		<-rc
		return r.Response{
			Body:       string(body),
			Err:        err,
			StatusCode: statusCode,
			Status:     status}.Response(logger.NewBuiltinLogger())
	}
}
