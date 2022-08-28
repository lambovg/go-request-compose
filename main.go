package main

import (
	"log"
	"time"
)

func main() {
	// sync requests
	var client = new(Request)
	client.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.Protocol = "https"
	client.Path = "hello-world.json"
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	// benchamrk async requests
	start := time.Now()

	Get{*client}.Request()
	// request to new url
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
	Get{*client}.Request()

	// async multiple requests without transaction
	log.Println("Multiple async")
	go AsyncGet("http://localhost:8080/hello-world.json")
	go AsyncGet("http://localhost:8080/ping.json")

	// group multiple async requests into
	log.Println("Group async requests")
	helloWorld := func() error { return AsyncGet("http://localhost:8080/hello-world.json") }
	zen := func() error { return AsyncGet("http://localhost:8080/zen.json") }
	GroupAsync([]func() error{helloWorld, zen})

	// benchmark
	end := time.Now()
	log.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}
