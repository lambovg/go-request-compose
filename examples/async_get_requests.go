package examples

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func AsyncGetRequests() {
	log.Println("Multiple async")

	future1 := request.Future("http://localhost:8080/hello-world.json")
	future2 := request.Future("http://localhost:8080/ping.json")
	future3 := request.Future("http://localhost:8080/zen.json")

	zen := future3()
	log.Print("return value", zen.Body)

	hello := future2()
	log.Print("return value", hello.Body)

	ping := future1()
	log.Print("return value", ping.Body)
}
