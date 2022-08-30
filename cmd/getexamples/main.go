package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func main() {
	// first solution
	request.Get{Url: "http://localhost:8080/hello-world.json"}.Request()

	// return value of Response
	rq := request.Get{Url: "http://localhost:8080/zen.json"}.Response()
	log.Println("return value", rq.Body)
	log.Println("return err", rq.Err)

	// alternative, better way
	//TODO: request.Get{Url: "url"}.Response()
	ping := request.Params{Url: "http://localhost:8080/ping.json"}.Get()
	log.Println("return value", ping.Body)
	log.Println("return err", ping.Err)

	// create client with basic setup and extend for erach call
	GithubClient := request.Client(request.Params{Hostname: "api.github.com", Protocol: "https"})
	GithubClient.Url = "https://api.github.com/zen"
	GithubClient.Get()
}
