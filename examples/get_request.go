package examples

import (
	cr "github.com/lambovg/go-request-compose/pkg/request"
	"log"
	"net/http"
	"time"
)

func GetRequest() {
	// prepare request object
	var request = new(cr.Params)
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

// OverrideTimeout client doesn't wait 30 secs request to finish
func OverrideTimeout() {
	client := http.Client{Timeout: 10 * time.Second}
	future := cr.Params{Url: "http://localhost:8080/timeout", Client: client}.Get()
	future()
}

func CompatibleHttpClient() {
	client := http.Client{Timeout: 10 * time.Second}
	params := cr.Params{Url: "http://localhost:8080/timeout"}

	future := cr.HttpClient{Client: &client}.Get(params)
	future()
}
