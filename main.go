package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Deployed v1"))
	})

	http.ListenAndServe(":8080", nil)
}
