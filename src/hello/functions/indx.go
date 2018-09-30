package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("indx.gohtml"))
}

func main() {
	xs := []string{"ZERO", "ONE", "TWO", "THREE", "FOUR", "FIVE"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"McLeod",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
