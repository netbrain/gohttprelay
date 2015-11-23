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
	handler.Director = func(req *http.Request) {
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.RequestURI() + req.URL.Path[1:]
		log.Println(req.URL)
	}
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
