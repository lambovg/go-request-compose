package request

import (
	r "github.com/lambovg/go-request-compose/pkg/response"
	"net/http"
)

// Delete sends delete to url or build url
func (p Params) Delete() func() *r.Response {

	if p.Url != "" {
		return p.delete(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.delete(p.Url)
}

// Delete with HttpClient struct
func (c HttpClient) Delete(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.delete(p.Url)
}

// Delete by given url
func Delete(url string) func() *r.Response {
	return Params{Url: url}.delete(url)
}

// delete .
func (p Params) delete(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodDelete)
}
