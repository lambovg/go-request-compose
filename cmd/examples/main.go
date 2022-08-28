package main

import (
	"log"
	"time"

	"github.com/lambovg/go-request-compose/examples"
)

func main() {
	// benchamrk async requests
	start := time.Now()

	examples.AsyncGetRequests()
	
	// benchmark
	end := time.Now()
	log.Printf("Requests processed after %v seconds\n", end.Sub(start).Seconds())
}
