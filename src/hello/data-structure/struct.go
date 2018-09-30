package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sages struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("struct.gohtml"))
}

func main() {
	x := sages{
		Name:  "Budha",
		Motto: "belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, x)
	if err != nil {
		log.Fatalln(err)
	}
}
