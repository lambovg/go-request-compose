package examples

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func AsyncGetGroupRequests() {
	log.Println("Group async requests")
	helloWorld := func() error { return request.AsyncGet("http://localhost:8080/hello-world.json") }
	zen := func() error { return request.AsyncGet("http://localhost:8080/zen.json") }
	request.GroupAsync([]func() error{helloWorld, zen})

	// new func for group requests
	request.GroupAsync2([]string{"http://localhost:8080/ping.json",
		"http://localhost:8080/hello-world.json"})

}
