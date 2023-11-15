package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

// GitHubZen overrides predefined host and schema with new url
func GitHubZen() {
	GithubClient := request.Client(request.Params{Hostname: "localhost:8080", Protocol: "http"})
	GithubClient.Url = "https://api.github.com/zen"

	get := GithubClient.Get()
	log.Println(get().Body)
}
