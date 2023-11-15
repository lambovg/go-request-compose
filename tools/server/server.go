package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type HellWorld struct {
	Hello string
}

type Ping struct {
	Ping string
}

type Zen struct {
	Zen string
}

func main() {

	http.HandleFunc("/hello-world.json", func(w http.ResponseWriter, r *http.Request) {
		response := HellWorld{"world"}
		body, err := json.Marshal(response)

		if err != nil {
			log.Println(err)
		}

		time.Sleep(2 * time.Second)

		responseWriter(w, r, string(body))
	})

	http.HandleFunc("/ping.json", func(w http.ResponseWriter, r *http.Request) {
		response := Ping{"pong"}
		body, err := json.Marshal(response)

		if err != nil {
			log.Println(err)
		}

		time.Sleep(2 * time.Second)

		responseWriter(w, r, string(body))
	})

	http.HandleFunc("/zen", func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		var err error

		response := Zen{"Keep it logically awesome"}
		body = []byte(response.Zen)

		if r.Header.Get("Accept") == "application/json" {
			body, err = json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
		}

		if r.Header.Get("Accept") == "application/xml" {
			body, err = xml.Marshal(response)
			w.Header().Set("Content-Type", "application/xml")
		}

		if err != nil {
			log.Println(err)
		}

		time.Sleep(12 * time.Millisecond)
		responseWriter(w, r, string(body))
	})

	http.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(30 * time.Second)
		responseWriter(w, r, "timeout")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/ping") {
			responseWriter(w, r, "ping: "+strings.Trim(r.URL.Path, "/ping"))
		} else {
			http.NotFound(w, r)
		}
	})

	log.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func responseWriter(w http.ResponseWriter, r *http.Request, body string) {
	allHeaders, _ := json.Marshal(r.Header)

	log.Println(r.RequestURI, string(allHeaders), "\n", body)

	fmt.Fprintf(w, body)
}
