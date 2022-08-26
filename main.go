package main

import (
	//"io/ioutil"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	var client = new(Request)
	client.hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.protocol = "https"
	client.path = "hello-world.json"
	client.url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	start := time.Now()

	Get{*client}.Request()

	// async
	paymentChan := make(chan *http.Response)
	client.url = "http://localhost:8080/hello-world.json"
	go GetAsync(client.url, paymentChan)

	// request in between
	client.url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
	Get{*client}.Request()

	// block for response
	paymentResponse := <-paymentChan
	defer paymentResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(paymentResponse.Body)
	log.Printf(string(bytes))

	// benchmark
	end := time.Now()
	log.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}
