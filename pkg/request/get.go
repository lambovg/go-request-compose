package request

import (
	"github.com/lambovg/go-request-compose/pkg/logger"
	r "github.com/lambovg/go-request-compose/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
)

// Get
func (p Params) Get() func() *r.Response {
	return get(p.Url, &p)
}

// Get
func Get(url string) func() *r.Response {
	return get(url, &Params{Url: url})
}

func get(url string, p *Params) func() *r.Response {
	var body []byte
	var err error

	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		response, err := p.Client.Do(NewRequest(http.MethodGet, url, nil).AttachHeaders(p).Request)

		if err == nil {
			defer response.Body.Close()
			body, _ = ioutil.ReadAll(response.Body)
			log.Println("async body", string(body))
		}
	}()

	return func() *r.Response {
		<-rc
		return r.Response{Body: string(body), Err: err}.Response(logger.NewBuiltinLogger())
	}
}
