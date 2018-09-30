package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err = template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")

	err = tpl.ExecuteTemplates(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplates(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplates(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplates(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
