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
		"Accept": {"application/json"},
	}

	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}

	//
	//map[string][]string
	//

	future := cr.Params{
		Url:      "http://localhost:8080/zen",
		Headers2: cfg,
		Headers:  []cr.Header{jsonHeader, xmlHeader}}.Get()
	future()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}
