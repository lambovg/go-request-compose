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
	Headers  []Header
	Headers2 Header2
}

type Header struct {
	Add string
	Set string
}

type Header2 map[string][]string

type requestFunc func(string) func() *cresponse.Response

func Client(p Params) Params {
	return p
}

// Headers attache headers
func Headers(rq *http.Request, p *Params) {
	for i := range p.Headers {
		h := strings.Split(p.Headers[i].Set, ":")
		rq.Header.Set(h[0], h[1])
	}

	for key, val := range p.Headers2 {
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
