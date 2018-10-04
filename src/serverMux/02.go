package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	//will run even with /dog/something/else
	mux.Handle("/dog/", d)
	//will return 404 when you run /cat/something/else
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
