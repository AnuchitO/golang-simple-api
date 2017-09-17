package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"name": "anuchit"}`))
	})

	log.Fatal(http.ListenAndServe(":1234", nil))
}
