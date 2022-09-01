package main

import (
	"log"
	"time"
	"github.com/lambovg/go-request-compose/pkg/request"
)

func main() {

	log.Println("getExamples")

	// async requets
	request.Params{Url: "http://localhost:8080/hello-world.json"}.Get()

	// promise
	future := request.Params{Url: "http://localhost:8080/ping.json"}.Get()
	log.Println(future().Body)

	// create client with basic setup and extend for erach call
	GithubClient := request.Client(request.Params{Hostname: "api.github.com", Protocol: "https"})
	GithubClient.Url = "https://api.github.com/zen"
	GithubClient.Get()

	time.Sleep(4 * time.Second)  // wait 1 sec
}
