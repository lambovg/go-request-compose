package examples

import (
	"log"
	"time"
	compose_request "github.com/lambovg/go-request-compose/pkg/request"
)

func GetRequest() {
	// prepare request object
	var request = new(compose_request.Request)
	request.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	request.Protocol = "https"
	request.Path = "hello-world.json"
	request.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	// benchamrk
	log.Printf("Get request start")
	start := time.Now()

	// sync request
	compose_request.Get{Params: *request}.Request()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
