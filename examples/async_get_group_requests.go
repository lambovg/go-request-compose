package examples

import (
	"log"

	"github.com/lambovg/go-request-compose/pkg/request"
)

func AsyncGetGroupRequests() {
	log.Println("Group async requests")

	request.FutureGroup([]string{"http://localhost:8080/ping.json",
		"http://localhost:8080/hello-world.json"}, request.Promise)

}
