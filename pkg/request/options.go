package request

import (
	r "github.com/lambovg/go-request-compose/pkg/response"
	"net/http"
)

// Options sends options to url or build url
func (p Params) Options() func() *r.Response {

	if p.Url != "" {
		return p.options(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.options(p.Url)
}

// Options with HttpClient struct
func (c HttpClient) Options(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.options(p.Url)
}

// Options by given url
func Options(url string) func() *r.Response {
	return Params{Url: url}.options(url)
}

// options .
func (p Params) options(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodOptions)
}
