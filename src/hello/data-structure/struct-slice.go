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
	tpl = template.Must(template.ParseFiles("struct-slice.gohtml"))
}

func main() {
	a := sages{
		Name:  "Budha",
		Motto: "belief of no beliefs",
	}
	b := sages{
		Name:  "Martin Luthur",
		Motto: "Hatred does not heal hatred",
	}
	x := sages{
		Name:  "Jesus",
		Motto: "love all",
	}
	y := sages{
		Name:  "Mohammed",
		Motto: "To evil with good is good",
	}
	z := sages{
		Name:  "None",
		Motto: "say nothhing",
	}

	try := []sages{a, b, x, y, z}

	err := tpl.Execute(os.Stdout, try)
	if err != nil {
		log.Fatalln(err)
	}
}
