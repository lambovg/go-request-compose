package request

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
)

type Request struct {
	Url      string
	Hostname string
	Protocol string
	Path     string
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
