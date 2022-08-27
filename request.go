package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

type Request struct {
	Url      string
	Hostname string
	Protocol string
	Path     string
}

type Get struct {
	params Request
}

type Post struct {
	params Request
}

func (r Get) Request() {
	resp, err := http.Get(r.params.Url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var response = Response{Body: string(body), Err: err}
	response.Response(NewBuiltinLogger())
}

func (r Post) Request() {
	//TODO implementation
}

// Deprecated
// Very first impementation
func GetAsync(url string, rc chan *http.Response) error {
	response, err := http.Get(url)

	if err == nil {
		rc <- response
	}

	return err
}

// TODO: parsing response should be optional
// TODO: should work with Request()
func AsyncGet(url string) error {
	response, err := http.Get(url)

	rc := make(chan *http.Response, 1)
	if err == nil {
		rc <- response
		msg := <-rc

		defer msg.Body.Close()
		body, err := ioutil.ReadAll(msg.Body)

		Response{Body: string(body), Err: err}.Response(NewBuiltinLogger())
	}

	return err
}

type requestAsync func() error

func GroupAsync(fn []func() error) {
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
		os.Exit(1)
	}

}
