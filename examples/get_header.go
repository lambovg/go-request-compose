package examples

import (
	cr "github.com/lambovg/go-request-compose/pkg/request"
	"log"
	"time"
)

func SetRequestHeadersInGetRequest() {
	// benchmark
	log.Printf("Get request start")
	start := time.Now()

	jsonHeader := cr.Header{Set: "Accept: application/json"}
	xmlHeader := cr.Header{Set: "Accept: application/xml"}

	cfg := map[string][]string{
		"Accept":          {"application/json"},
		"Accept-Language": {"en-us"},
	}

	future := cr.Params{
		Url:     "http://localhost:8080/zen",
		Headers: []cr.Header{jsonHeader, xmlHeader}}.Get()
	future()

	future2 := cr.Params{Url: "http://localhost:8080/zen", Headers2: cfg}.Get()
	future2()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}

func SetRequestHeadersWithNetHttpHeaderFormat() {
	setHeaders := map[string][]string{
		"Accept":          {"application/json"},
		"Accept-Language": {"en-us"},
	}

	addHeaders := map[string][]string{
		"Cache-control": {"no-cache"},
	}

	headers := cr.HeaderSetAdd{Set: setHeaders, Add: addHeaders}

	future := cr.Params{
		Url:     "http://localhost:8080/zen",
		Headers3: headers}.Get()
	future()
}