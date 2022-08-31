package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func main() {

	ping := request.Params{Url: "http://localhost:8080/ping.json"}.Get()
	log.Println("return value", ping.Body)
	log.Println("return err", ping.Err)

	// create client with basic setup and extend for erach call
	GithubClient := request.Client(request.Params{Hostname: "api.github.com", Protocol: "https"})
	GithubClient.Url = "https://api.github.com/zen"
	GithubClient.Get()
}
