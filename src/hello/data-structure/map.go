package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("map.gohtml"))
}

func main() {
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MK",
		"Meditate": "Budha",
		"Love":     "Jesus",
		"Prophet":  "Mohammed",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
