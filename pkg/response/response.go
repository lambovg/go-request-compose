package response

import (
	"github.com/lambovg/go-request-compose/pkg/logger"
	"net/http"
)

type Response struct {
	Body       string
	Err        error
	StatusCode int
	Header     http.Header
	Status     string
}

func (r Response) Response(log *logger.BuiltinLogger) *Response {
	log.Printf(r.Body)

	if r.Err != nil {
		log.Println(r.Err)
	}

	return &r
}
