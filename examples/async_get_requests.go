package examples

import (
	"log"

	"github.com/lambovg/go-request-compose/pkg/request"
)

func AsyncGetRequests() {
	log.Println("Multiple async")

	future1 := request.GetAsync("http://localhost:8080/hello-world.json")
	future2 := request.GetAsync("http://localhost:8080/ping.json")
	futuer3 := request.GetAsync("http://localhost:8080/zen.json")

	futuer3()
	future2()
	future1()

}
