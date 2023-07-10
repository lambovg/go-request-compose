package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
)

func main() {

	request.Params{Url: "http://localhost:8080/ping.json"}.Post()
}
