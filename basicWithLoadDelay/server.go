package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/page/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pageData/test.html")
	})

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "fakeApi/myData.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pageData/index.html")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
