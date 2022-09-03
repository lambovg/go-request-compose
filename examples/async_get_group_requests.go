package examples

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func AsyncGetGroupRequests() {
	log.Println("Group async requests")

	//TODO add option to Future group get and post requests
	request.FutureGroup([]string{"http://localhost:8080/ping.json",
		"http://localhost:8080/hello-world.json"}, request.Get)

}
