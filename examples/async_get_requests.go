package examples

import (
	"log"
	"github.com/lambovg/go-request-compose/pkg/request"
)

func AsyncGetRequests() {
	log.Println("Multiple async")
	
	go request.AsyncGet("http://localhost:8080/hello-world.json")
	go request.AsyncGet("http://localhost:8080/ping.json")
}