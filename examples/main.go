package examples

import (
	"log"
	"time"
)

func main() {

	// benchmark async requests
	start := time.Now()

	//AsyncGetRequests()
	//AsyncGetGroupRequests()
	GetExtendRequest()
	//SetRequestHeadersWithNetHttpHeaderFormat()
	//OverrideExistingHeaderWithNetHttpHeaderFormat()
	//OverrideTimeout()
	//CompatibleHttpClient()

	// benchmark
	end := time.Now()
	log.Printf("Requests processed after %v seconds\n", end.Sub(start).Seconds())
}
