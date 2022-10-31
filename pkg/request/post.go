package request

import (
	r "github.com/lambovg/go-request-compose/pkg/response"
	"net/http"
)

// Post send post to url or build url
func (p Params) Post() func() *r.Response {

	if p.Url != "" {
		return p.post(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.post(p.Url)
}

// Post with HttpClient struct
func (c HttpClient) Post(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.post(p.Url)
}

// Post by given url
func Post(url string) func() *r.Response {
	return Params{Url: url}.post(url)
}

// post .
func (p Params) post(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodPost)
}
