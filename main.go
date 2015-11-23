package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Expected 1 argument with a target url, i.e http://www.example.org")
	}

	target, err := url.Parse(args[0])
	if err != nil {
		log.Fatal(err)
	}

	handler := httputil.NewSingleHostReverseProxy(target)
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
