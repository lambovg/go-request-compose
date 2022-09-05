package examples

import (
	crequest "github.com/lambovg/go-request-compose/pkg/request"
	"log"
	"time"
)

func SetRequestHeadersInGetRequest() {
	// benchmark
	log.Printf("Get request start")
	start := time.Now()

	jsonHeader := crequest.Header{Set: "Accept: application/json"}
	xmlHeader := crequest.Header{Set: "Accept: application/xml"}

	future := crequest.Params{Url: "http://localhost:8080/zen",
		Headers: []crequest.Header{jsonHeader, xmlHeader}}.Get()
	future()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
