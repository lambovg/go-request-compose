package request

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/lambovg/go-request-compose/pkg/logger"
	cresponse "github.com/lambovg/go-request-compose/pkg/response"
)

func (p Params) Get() func() *cresponse.Response {
	return get(p.Url)
}

func Get(url string) func() *cresponse.Response {
	return get(url)
}

func get(url string) func() *cresponse.Response {
	var body []byte
	var err error
	
	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		response, err := http.Get(url)
		if err == nil {
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			log.Println("async body", string(body))
		}
	}()

	return func() *cresponse.Response {
		<-rc
		return cresponse.Response{Body: string(body), Err: err}.Response(logger.NewBuiltinLogger())
	}
}
