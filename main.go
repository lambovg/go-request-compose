package main

import (
	"context"
	"log"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// sync requests
	var client = new(Request)
	client.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.Protocol = "https"
	client.Path = "hello-world.json"
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	// benchamrk async requests
	start := time.Now()

	Get{*client}.Request()
	client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
	Get{*client}.Request()

	// async
	errGrp, _ := errgroup.WithContext(context.Background())

	errGrp.Go(func() error {
		return AsyncGet("http://localhost:8080/hello-world.json")
	})

	errGrp.Go(func() error {
		return AsyncGet("http://localhost:8080/ping.json")
	})

	err := errGrp.Wait()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// group multiple async requests
	GroupAsync(func() error { return AsyncGet("http://localhost:8080/hello-world.json") }, func() error { return AsyncGet("http://localhost:8080/hello-world.json") })

	// benchmark
	end := time.Now()
	log.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}
