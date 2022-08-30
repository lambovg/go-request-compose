package main

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func main() {
	// first solution
	request.Get{Url: "http://localhost:8080/hello-world.json"}.Request()
	request.Post{Url: "http://localhost:8080/hello-world.json", Body: "1=1"}.Request()

	// return value of Response
	rq := request.Get{Url: "http://localhost:8080/zen.json"}.Response()
	log.Println("return value", rq.Body)
	log.Println("return err", rq.Err)

}
