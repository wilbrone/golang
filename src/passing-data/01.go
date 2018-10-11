package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(res, "Do my search: "+v)
}

// try this
//localhost:8080/?q=something else
