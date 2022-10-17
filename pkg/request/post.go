package request

import (
	"io"
	"log"
	"net/http"

	"github.com/lambovg/go-request-compose/pkg/logger"
	r "github.com/lambovg/go-request-compose/pkg/response"
)

// Post send post to url or build url
func (p Params) Post() func() *r.Response {

	if p.Url != "" {
		return p.post(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.post(p.Url)
}

// Get by given url
func Post(url string) func() *r.Response {
	return Params{Url: url}.post(url)
}

// post .
func (p Params) post(url string) func() *r.Response {
	var body []byte
	var err error
	var response *http.Response
	var statusCode int
	var status string

	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		if p.FormData != nil {
			response, err = http.PostForm(url, p.FormData)
		} else {
			response, err = p.Client.Do(NewRequest(http.MethodPost, url, p.Body).AttachHeaders(&p).Request)
		}

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
