package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request")
		w.Write([]byte("Deployed v7 - Working"))
	})

	http.ListenAndServe(":8080", nil)
}
