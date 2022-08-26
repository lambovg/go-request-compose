package main

import (
	"log"
)

type Response struct {
	Body string
	Err  error
}

func (r Response) Response() {
	log.Printf(r.Body)

	if r.Err != nil {
		log.Fatalln(r.Err)
	}
}
