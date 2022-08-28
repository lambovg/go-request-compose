package examples

import (
	"log"
	"time"
	"github.com/lambovg/go-request-compose/pkg/request"
)

func GetExtendRequest() {
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

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
