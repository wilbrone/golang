package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {

	//will run even with /dog/something/else
	// http.HandleFunc("/dog/", d)
	http.Handle("/dog/", http.HandlerFunc(d))

	//will return 404 when you run /cat/something/else
	// http.HandleFunc("/cat", c)
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
