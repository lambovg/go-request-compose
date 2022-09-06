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
	Headers3 HeaderSetAdd
}

type HeaderSetAdd struct {
	Add Header2
	Set Header2
}

type Header2 map[string][]string

type requestFunc func(string) func() *cresponse.Response

func Client(p Params) Params {
	return p
}

// Headers attache headers
func Headers(rq *http.Request, p *Params) {
	for key, val := range p.Headers3.Set {
		rq.Header.Set(key, strings.Join(val, ","))
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
