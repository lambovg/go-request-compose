package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

// PromiseGetRequest Run this example with: go run cmd/examples/get/promise.go
func PromiseGetRequest() {
	future := request.Params{Url: "http://localhost:8080/ping.json"}.Get()
	log.Println(future().Body)
}
