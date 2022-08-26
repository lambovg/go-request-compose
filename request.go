package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	url      string
	hostname string
	protocol string
	path     string
}

type Get struct {
	params Request
}

type Post struct {
	params Request
}

func (r Get) Request() {
	resp, err := http.Get(r.params.url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var response = Response{body: string(body), err: err}
	response.Response()
}

func (r Post) Request() {
	//TODO implementation
}