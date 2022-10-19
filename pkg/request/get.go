package request

import (
	"net/http"

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

// Getv2 .
func (p Params) Getv2() (Params, func() *r.Response) {
	if p.Url != "" {
		return p, p.get(p.Url)
	}

	p.Url = p.BuildUrl()

	return p, p.get(p.Url)
}

// Get with HttpClient struct
func (c HttpClient) Get(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.get(p.Url)
}

// Get By url
func Get(url string) func() *r.Response {
	return Params{Url: url}.get(url)
}

// get .
func (p Params) get(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodGet)
}
