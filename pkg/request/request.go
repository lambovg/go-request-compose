package request

import (
	cresponse "github.com/lambovg/go-request-compose/pkg/response"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Params struct {
	Url      string
	Hostname string
	Protocol string
	Path     string
	Headers  Headers
	Client   http.Client
}

type Headers struct {
	Add Header
	Set Header
}

type Header map[string][]string

type requestFunc func(string) func() *cresponse.Response

func Client(p Params) Params {
	return p
}

// AttachHeaders request headers
func AttachHeaders(rq *http.Request, p *Params) {
	// set / override existing
	for key, val := range p.Headers.Set {
		rq.Header.Set(key, strings.Join(val, ","))
	}

	// add / extend definition of existing
	for key, val := range p.Headers.Add {
		rq.Header.Add(key, strings.Join(val, ","))
	}
}

func FutureGroup(fn []string, rq requestFunc) {
	errorChan := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup

	for i := range fn {
		url := fn[i]

		wg.Add(1)

		go func() {
			defer wg.Done()
			res := rq(url)()

			if res.Err != nil {
				errorChan <- res.Err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-errorChan:
		close(errorChan)
		log.Println("Error encountered: ", err)
	}
}
