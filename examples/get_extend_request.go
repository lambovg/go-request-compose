package examples

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
	"time"
)

func GetExtendRequest() {
	// sync requests
	var client = new(request.Params)
	client.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.Protocol = "https"
	client.Path = "/hello-world.json"

	// benchmark async requests
	start := time.Now()

	// errors not block further execution
	client.Get()()

	// request to new path
	client.Path = "/ping.json"
	client.Get()()

	// request to new url
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/zen.json"
	client.Get()()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
