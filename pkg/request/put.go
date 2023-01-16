package request

import (
	r "github.com/lambovg/go-request-compose/pkg/response"
	"net/http"
)

// Put sends put to url or build url
func (p Params) Put() func() *r.Response {

	if p.Url != "" {
		return p.put(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.put(p.Url)
}

// Put with HttpClient struct
func (c HttpClient) Put(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.put(p.Url)
}

// Put by given url
func Put(url string) func() *r.Response {
	return Params{Url: url}.put(url)
}

// put .
func (p Params) put(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodPut)
}
