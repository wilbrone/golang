package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your reqest method at foo: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your reqest method at bar: ", req.Method, "\n\n")

	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusSeeOther)

}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your reqest method at barred: ", req.Method, "\n\n")

	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
