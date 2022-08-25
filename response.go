package main

import (
	"log"
)

type Response struct {
	body string
	err  error
}

func (r Response) Response() {
	log.Printf(r.body)

	if r.err != nil {
		log.Fatalln(r.err)
	}
}
