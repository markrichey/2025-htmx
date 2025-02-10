package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var c, err = r.Cookie("cookieTest")
		if err != nil {
			fmt.Println("Nope")
		} else {
			fmt.Println(c.Value)
		}
		http.ServeFile(w, r, "static/about.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    "cookieTest",
			Value:   "BooYa",
			Expires: time.Now().Add(120 * time.Second),
		})
		http.ServeFile(w, r, "static/index.html")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
