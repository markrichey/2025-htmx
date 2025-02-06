package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
