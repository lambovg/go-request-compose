package request

import (
	"github.com/lambovg/go-request-compose/pkg/logger"
	r "github.com/lambovg/go-request-compose/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (p Params) Get() func() *r.Response {
	return get(p.Url)
}

func Get(url string) func() *r.Response {
	return get(url)
}

func get(url string) func() *r.Response {
	var body []byte
	var err error

	rc := make(chan *http.Response, 1)

	go func() {
		defer close(rc)

		client := &http.Client{
			Timeout: time.Second * 10,
		}

		response, err := client.Get(url)

		if err == nil {
			defer response.Body.Close()
			body, _ = ioutil.ReadAll(response.Body)
			log.Println("async body", string(body))
		}
	}()

	return func() *r.Response {
		<-rc
		return r.Response{Body: string(body), Err: err}.Response(logger.NewBuiltinLogger())
	}
}
