package request

import (
	"github.com/lambovg/go-request-compose/pkg/logger"
	cresponse "github.com/lambovg/go-request-compose/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
)

func (p Params) Get() *cresponse.Response {
	resp, err := http.Get(p.Url)

	if err != nil {
		log.Println(err)
	}

	if (err == nil) {
		body, err := ioutil.ReadAll(resp.Body)
		var response = &cresponse.Response{Body: string(body), Err: err}
		return response.Response(logger.NewBuiltinLogger())
	}
	
	var response = &cresponse.Response{Body: "", Err: err}
	return response.Response(logger.NewBuiltinLogger())
	
}
