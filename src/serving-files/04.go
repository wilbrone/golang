package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/mine", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
		<!--does not serve-->
		<img src="/home/aphya5/Pictures/mine.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "/home/aphya5/Pictures/mine.jpg")
}
