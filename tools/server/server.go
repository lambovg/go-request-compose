package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		result, err := json.Marshal(response)

		if err != nil {
			log.Println(err)
		}

		time.Sleep(2 * time.Second)

		log.Printf(string(result))
		fmt.Fprintf(w, string(result))
	})

	http.HandleFunc("/ping.json", func(w http.ResponseWriter, r *http.Request) {
		response := Ping{"pong"}
		result, err := json.Marshal(response)

		if err != nil {
			log.Println(err)
		}

		time.Sleep(2 * time.Second)

		log.Printf(string(result))
		fmt.Fprintf(w, string(result))
	})

	http.HandleFunc("/zen.json", func(w http.ResponseWriter, r *http.Request) {
		response := Zen{"Keep it logically awesome"}
		result, err := json.Marshal(response)

		if err != nil {
			log.Println(err)
		}

		time.Sleep(12 * time.Millisecond)

		log.Printf(string(result))
		fmt.Fprintf(w, string(result))
	})

	log.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
