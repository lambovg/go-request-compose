package get

import (
	"github.com/lambovg/go-request-compose/pkg/logger"
	cresponse "github.com/lambovg/go-request-compose/pkg/response"
	"io/ioutil"
	"net/http"
)

func Promise(url string) func() *cresponse.Response {
	var body []byte
	var err error

	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		response, err := http.Get(url)
		if err == nil {
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
		}
	}()

	return func() *cresponse.Response {
		<-rc
		return cresponse.Response{Body: string(body), Err: err}.Response(logger.NewBuiltinLogger())
	}
}
