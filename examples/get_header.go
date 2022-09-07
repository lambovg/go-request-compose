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
		"Cache-Control":   {"must-revalidate"},
	}

	future := cr.Params{
		Url:     "http://localhost:8080/zen",
		Headers: cr.Headers{Set: setHeaders}}.Get()
	future()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}

func AppendRequestHeadersWithNetHttpHeaderFormat() {
	// benchmark
	log.Printf("Get request start")
	start := time.Now()

	setHeaders := map[string][]string{
		"Cache-Control":   {"must-revalidate"},
	}

	addHeaders := map[string][]string{
		"Cache-control": {"no-cache"},
	}

	// Resulting Cache-Control header value should be "mu-revalidate, no-cache"

	future := cr.Params{
		Url:     "http://localhost:8080/zen",
		Headers: cr.Headers{Set: setHeaders, Add: addHeaders}}.Get()
	future()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())
}

func OverrideExistingHeaderWithNetHttpHeaderFormat() {
	// benchmark
	log.Printf("Get request start")
	start := time.Now()

	setHeaders := map[string][]string{
		"User-Agent":   {"my-agent"},
		"Cache-Control":   {"must-revalidate"},
		"Accept-Encoding": {"gzip"},
	}

	addHeaders := map[string][]string{
		"User-Agent": {"v11"},
		"Cache-control": {"no-cache"},
		"Accept-Encoding": {"br"},

	}

	// Resulting Cache-Control header value is "must-revalidate, no-cache"
	// Resulting User-Agent should be "my-agent v11" but it is only my-agent, possible bug

	future := cr.Params{
		Url:     "http://localhost:8080/zen",
		Headers: cr.Headers{Set: setHeaders, Add: addHeaders}}.Get()
	future()

	// benchmark
	end := time.Now()
	log.Printf("Get request took %v seconds\n", end.Sub(start).Seconds())	
}
