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

func main() {

	http.HandleFunc("/hello-world.json", func(w http.ResponseWriter, r *http.Request) {
		helloWorld := HellWorld{"world"}
		result, err := json.Marshal(helloWorld)
		
		if err != nil {
			log.Println(err)
		}

		time.Sleep(2 * time.Second)
		fmt.Fprintf(w, string(result))
	})

	log.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
