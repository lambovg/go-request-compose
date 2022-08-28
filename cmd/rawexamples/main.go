package main

import (
	"log"
	"time"
	"github.com/lambovg/go-request-compose/pkg/request"
)

func main() {
	// sync requests
	var client = new(request.Request)
	client.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.Protocol = "https"
	client.Path = "hello-world.json"
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	// benchamrk async requests
	start := time.Now()

	request.Get{Params: *client}.Request()
	// request to new url
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
	request.Get{Params: *client}.Request()

	// async multiple requests without transaction
	log.Println("Multiple async")
	go request.AsyncGet("http://localhost:8080/hello-world.json")
	go request.AsyncGet("http://localhost:8080/ping.json")

	// group multiple async requests into
	log.Println("Group async requests")
	helloWorld := func() error { return request.AsyncGet("http://localhost:8080/hello-world.json") }
	zen := func() error { return request.AsyncGet("http://localhost:8080/zen.json") }
	request.GroupAsync([]func() error{helloWorld, zen})

	// benchmark
	end := time.Now()
	log.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}