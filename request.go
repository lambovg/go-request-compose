package main

import (
	"io/ioutil"
	"log"
	"net/http"
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
	response.Response()
}

func (r Post) Request() {
	//TODO implementation
}

func GetAsync(url string, rc chan *http.Response) error {
	response, err := http.Get(url)

	if err == nil {
		rc <- response
	}

	return err
}
