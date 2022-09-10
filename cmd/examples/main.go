package main

import (
	"github.com/lambovg/go-request-compose/examples"
	"log"
	"time"
)

func main() {
	// benchmark async requests
	start := time.Now()

	//examples.AsyncGetRequests()
	//examples.AsyncGetGroupRequests()
	examples.GetExtendRequest()
	//examples.SetRequestHeadersWithNetHttpHeaderFormat()
	//examples.OverrideExistingHeaderWithNetHttpHeaderFormat()
	//examples.OverrideTimeout()

	// benchmark
	end := time.Now()
	log.Printf("Requests processed after %v seconds\n", end.Sub(start).Seconds())
}
