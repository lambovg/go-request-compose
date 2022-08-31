package request

import (
	"github.com/lambovg/go-request-compose/pkg/logger"
	cresponse "github.com/lambovg/go-request-compose/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
)

// Future Get Request
func Future(url string) func() *cresponse.Response {
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

// AsyncGet TODO: should work with Request()
func AsyncGet(url string) error {
	response, err := http.Get(url)

	rc := make(chan *http.Response, 1)
	if err == nil {
		rc <- response
		msg := <-rc

		defer msg.Body.Close()
		body, err := ioutil.ReadAll(msg.Body)

		cresponse.Response{Body: string(body), Err: err}.Response(logger.NewBuiltinLogger())
	}

	return err
}

func (p Params) Get() *cresponse.Response {
	resp, err := http.Get(p.Url)

	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var response = &cresponse.Response{Body: string(body), Err: err}
	return response.Response(logger.NewBuiltinLogger())
}
