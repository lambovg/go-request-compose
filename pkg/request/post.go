package request

import (
	"io"
	"log"
	"net/http"

	"github.com/lambovg/go-request-compose/pkg/logger"
	r "github.com/lambovg/go-request-compose/pkg/response"
)

func (p Params) Post() func() *r.Response {
	return p.post(p.Url)
}

// post .
func (p Params) post(url string) func() *r.Response {
	var body []byte
	var err error
	var statusCode int
	var status string

	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		response, err := p.Client.Do(NewRequest(http.MethodPost, url, p.Body).AttachHeaders(&p).Request)

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
