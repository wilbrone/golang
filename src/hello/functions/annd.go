package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("annd.gohtml"))
}

func main() {
	u1 := user{
		Name:  "Gandhi",
		Motto: "Belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Budhaa",
		Motto: "be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
