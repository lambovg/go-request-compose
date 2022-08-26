package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	var client = new(Request)
	client.hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.protocol = "https"
	client.path = "hello-world.json"
	client.url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	start := time.Now()

	Get{*client}.Request()
	client.url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
	Get{*client}.Request()

	// async
	helloWorldChan := make(chan *http.Response, 1)
	pingChain := make(chan *http.Response, 1)
	errGrp, _ := errgroup.WithContext(context.Background())

	errGrp.Go(func() error { return GetAsync("http://localhost:8080/hello-world.json", helloWorldChan) })
	errGrp.Go(func() error { return GetAsync("http://localhost:8080/ping.json", pingChain) })

	err := errGrp.Wait()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// block for response
	helloWorldResponse := <-helloWorldChan
	defer helloWorldResponse.Body.Close()
	helloWorldBytes, _ := ioutil.ReadAll(helloWorldResponse.Body)
	log.Printf(string(helloWorldBytes))

	pingResponse := <-pingChain
	defer pingResponse.Body.Close()
	pingBytes, _ := ioutil.ReadAll(pingResponse.Body)
	log.Printf(string(pingBytes))

	// benchmark
	end := time.Now()
	log.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}
