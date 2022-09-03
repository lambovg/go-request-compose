package examples

import (
	"github.com/lambovg/go-request-compose/pkg/request"
	"log"
)

func AsyncGetRequests() {
	log.Println("Multiple async")

	future1 := request.Params{Url: "http://localhost:8080/hello-world.json"}.Get()
	future2 := request.Params{Url: "http://localhost:8080/ping.json"}.Get()
	future3 := request.Params{Url: "http://localhost:8080/zen"}.Get()

	zen := future3()
	log.Print("return value", zen.Body)

	hello := future2()
	log.Print("return value", hello.Body)

	ping := future1()
	log.Print("return value", ping.Body)
}
