package examples

import (
	cr "github.com/lambovg/go-request-compose/pkg/request"
	"log"
	"time"
)

func SetRequestHeadersWithNetHttpHeaderFormat() {
	// benchmark
	log.Printf("Get request start")
	start := time.Now()

	setHeaders := map[string][]string{
		"Accept":          {"application/json"},
		"Accept-Language": {"en-us"},
	}

	addHeaders := map[string][]string{
		"Cache-control": {"no-cache"},
	}

	future := cr.Params{
		Url:      "http://localhost:8080/zen",
		Headers: cr.Headers{Set: setHeaders, Add: addHeaders}}.Get()
	future()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
