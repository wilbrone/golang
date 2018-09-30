package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	home := Page{
		Title:   " Nothing Escaped",
		Heading: "Nothing is escaped with template/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := tpl.Execute(os.Stdout, "main.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
}
