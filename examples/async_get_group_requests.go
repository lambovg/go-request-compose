package examples

import (
	"log"
	"github.com/lambovg/go-request-compose/pkg/request"
)

func AsyncGetGroupRequests() {
	log.Println("Group async requests")
	helloWorld := func() error { return request.AsyncGet("http://localhost:8080/hello-world.json") }
	zen := func() error { return request.AsyncGet("http://localhost:8080/zen.json") }
	request.GroupAsync([]func() error{helloWorld, zen})

}