package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home of Go</h1>")
}

func callOut(w http.ResponseWriter, r *http.Request) {
	url := "http://www.google.com"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "text/html")
	io.Copy(w, res.Body)
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/out", callOut)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
