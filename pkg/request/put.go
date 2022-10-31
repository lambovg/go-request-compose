package request

import (
	r "github.com/lambovg/go-request-compose/pkg/response"
	"net/http"
)

// Put send post to url or build url
func (p Params) Put() func() *r.Response {

	if p.Url != "" {
		return p.put(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.post(p.Url)
}

// Put with HttpClient struct
func (c HttpClient) Put(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.post(p.Url)
}

// Put Get by given url
func Put(url string) func() *r.Response {
	return Params{Url: url}.post(url)
}

// put .
func (p Params) put(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodPut)
}
