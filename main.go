package main

import "net/http"
import "log"

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"name": "nong"}`))
  })

  log.Fatal(http.ListenAndServe(":1234", nil))
}
