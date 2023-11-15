package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"time"
)

// AsyncGetRequest Run this example with: go run cmd/examples/get/async.go
// Start the server with: go run tools/server/server.go
func AsyncGetRequest() {
	// async requests
	request.Params{Url: "http://localhost:8080/ping/async"}.Get()

	time.Sleep(1 * time.Second)
}
