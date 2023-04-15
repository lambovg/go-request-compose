package request

import (
	r "github.com/lambovg/go-request-compose/pkg/response"
	"net/http"
)

// Connect sends connect to url or build url
func (p Params) Connect() func() *r.Response {

	if p.Url != "" {
		return p.connect(p.Url)
	}

	p.Url = p.BuildUrl()

	return p.connect(p.Url)
}

// Connect with HttpClient struct
func (c HttpClient) Connect(p Params) func() *r.Response {
	p.Client = *c.Client
	return p.connect(p.Url)
}

// Connect by given url
func Connect(url string) func() *r.Response {
	return Params{Url: url}.connect(url)
}

// connect .
func (p Params) connect(url string) func() *r.Response {
	return p.NewRequest(url, http.MethodConnect)
}
