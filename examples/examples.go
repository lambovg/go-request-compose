package examples

import (
	"log"
	"time"
    "github.com/lambovg/go-request-compose/pkg/request"
)

func SyncRequest() {
	// sync requests
	var client = new(request.Request)
	client.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.Protocol = "https"
	client.Path = "hello-world.json"
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	// benchamrk async requests
	start := time.Now()
    
    request.Get{*client}.Request()

	// benchmark
	end := time.Now()
	log.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}
