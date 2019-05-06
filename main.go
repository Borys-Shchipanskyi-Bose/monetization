package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there. %v", r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Example page. Try /hello api")
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello/", hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
