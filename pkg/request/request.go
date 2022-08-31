package request

import (
	"log"
	"sync"
	cresponse "github.com/lambovg/go-request-compose/pkg/response"
)

type Params struct {
	Url      string
	Hostname string
	Protocol string
	Path     string
}

type requestFunc func(string) func() *cresponse.Response

func Client(p Params) Params {
	return p
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
