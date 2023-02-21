package request

import (
	r "github.com/lambovg/go-request-compose/pkg/response"
	"net/http"
)

// Head sends head to url or build url
func (p Params) Head() func() *r.Response {

	if p.Url != "" {
		return p.head(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.head(p.Url)
}

// Head with HttpClient struct
func (c HttpClient) Head(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.head(p.Url)
}

// Head by given url
func Head(url string) func() *r.Response {
	return Params{Url: url}.head(url)
}

// head .
func (p Params) head(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodHead)
}
