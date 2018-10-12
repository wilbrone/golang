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
	http.ListenAndServe(":8000", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method, "\n\n")

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method, "\n\n")

	tpl.ExecuteTemplate(res, "index.gohtml", nil)

}
