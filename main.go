package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("it come here")
		w.Write([]byte(`{"name": "anuchit"}`))
	})
	log.Println("start server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
