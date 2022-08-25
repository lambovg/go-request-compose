package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	body string
	err  error
}

type Get struct {
	url string
}

type Post struct {
	Get
}

func (r Response) Response() {
	log.Printf(r.body)

	if r.err != nil {
		log.Fatalln(r.err)
	}
}

func (r Get) Request() {
	resp, err := http.Get(r.url)

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

func main() {
	Get{url: "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"}.Request()

	extend := Get{url: "https://d2kgi8nio2h9bn.cloudfront.net"}
	Get{url: extend.url + "/ping.json"}.Request()
	Get{url: extend.url + "/hello-world.json"}.Request()
}

//func main() {
//	resp, err := http.Get("https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json")
//
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	var response = Response{body: string(body), err: err}
//	response.Response()
//
//	//sb := string(body)
//	//log.Printf(sb)
//}
