package request

import (
	"context"
	"log"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Params struct {
	Url      string
	Hostname string
	Protocol string
	Path     string
}

func Client(p Params) Params {
	return p
}

// GroupAsync deprecated
func GroupAsync(fn []func() error) bool {
	errGrp, _ := errgroup.WithContext(context.Background())

	for i := range fn {
		request := fn[i]
		errGrp.Go(func() error {
			return request()
		})
	}

	err := errGrp.Wait()
	if err != nil {
		log.Println(err)
	}

	return err == nil
}

// GroupAsync2 New interface for creating async group requests
func GroupAsync2(fn []string) {
	errGrp, _ := errgroup.WithContext(context.Background())

	for i := range fn {
		url := fn[i]
		errGrp.Go(func() error {
			return AsyncGet(url)
		})
	}

	err := errGrp.Wait()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func GroupAsync3(fn []string) {
	errorChan := make(chan error)
	wgDone := make(chan bool)
	
	var wg sync.WaitGroup

	for i := range fn {
		url := fn[i]

		wg.Add(1)

		go func() {
			defer wg.Done()
			error := AsyncGet(url)

			if error != nil {
				errorChan <- error
			}
		}()
	}

	go func () {
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
