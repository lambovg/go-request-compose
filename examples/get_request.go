package examples

import (
	crequest "github.com/lambovg/go-request-compose/pkg/request"
	"log"
	"time"
)

func GetRequest() {
	// prepare request object
	var request = new(crequest.Params)
	request.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	request.Protocol = "https"
	request.Path = "hello-world.json"
	request.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	// benchmark
	log.Printf("Get request start")
	start := time.Now()

	// async request
	request.Get()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
