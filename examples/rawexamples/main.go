package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
	"time"
)

func main() {
	// sync requests
	var params = new(request.Params)
	params.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	params.Protocol = "https"
	params.Path = "hello-world.json"
	params.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	// benchmark async requests
	start := time.Now()

	request.Client(*params).Get()

	//request to new url
	params.Url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
	request.Client(*params).Get()

	// benchmark
	end := time.Now()
	log.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}
