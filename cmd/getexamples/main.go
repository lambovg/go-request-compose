package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func main() {

	request.Get{Url: "http://localhost:8080/hello-world.json"}.Request()
	request.Post{Url: "http://localhost:8080/hello-world.json", Body: "1=1"}.Request()

	//	execute get request with response
	response := request.Get{Url: "http://localhost:8080/zen.json"}.Response()
	log.Println("return value", response.Body)
	log.Println("return err", response.Err.Error())
}
