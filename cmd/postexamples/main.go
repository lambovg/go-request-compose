package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
)

func main() {
	// first solution
	request.Post{Url: "http://localhost:8080/hello-world.json"}.Request()

	request.Params{Url: "http://localhost:8080/ping.json"}.Post()
}
