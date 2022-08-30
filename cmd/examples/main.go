package main

import (
	"log"
	"time"

	"github.com/lambovg/go-request-compose/examples"
)

func main() {
	// benchmark async requests
	start := time.Now()

	examples.AsyncGetRequests()
	//examples.AsyncGetGroupRequests()

	// benchmark
	end := time.Now()
	log.Printf("Requests processed after %v seconds\n", end.Sub(start).Seconds())
}
