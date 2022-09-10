package request

import (
	"bytes"
	"github.com/lambovg/go-request-compose/pkg/logger"
	r "github.com/lambovg/go-request-compose/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
)

// Get
func (p Params) Get() func() *r.Response {

	if p.Url != "" {
		return get(p.Url, &p)
	}

	var b bytes.Buffer
	b.WriteString(p.Protocol)
	b.WriteString("://")
	b.WriteString(p.Hostname)
	b.WriteString(p.Path)

	if p.QueryString != "" {
		b.WriteString("?")
		b.WriteString(p.QueryString)
	}

	p.Url = b.String()

	return get(p.Url, &p)
}

func (c HttpClient) Get(p Params) func() *r.Response {
	p.Client = c.Client
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
