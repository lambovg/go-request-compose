package examples

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func AsyncGetGroupRequests() {
	log.Println("Group async requests")

	request.FutureGroup([]string{"http://localhost:8080/ping.json",
		"http://localhost:8080/hello-world.json"}, request.AsyncGet)

}
