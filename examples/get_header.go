package examples

import (
	"log"
	"time"

	crequest "github.com/lambovg/go-request-compose/pkg/request"
)

func GetRequestSetHeader() {
	// benchmark
	log.Printf("Get request start")
	start := time.Now()

	future := crequest.Params{Url: "http://localhost:8080/zen", Headers: []string{"Accept:application/json"}}.Get()
	future()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
