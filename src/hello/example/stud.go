package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("stud.gohtml"))
}

func main() {
	p1 := struct {
		Name string
		Age  int
	}{
		"James Willow",
		42,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
