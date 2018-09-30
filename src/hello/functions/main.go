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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// type items struct {
// 	Wisdom    []sages
// 	Transport []car
// }

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("main.gohtml"))
}

func fisrtThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

// func init() {
// 	tpl = template.Must(template.ParseFiles("main.gohtml"))
// }

func main() {
	a := sages{
		Name:  "Budha",
		Motto: "belief of no beliefs",
	}
	b := sages{
		Name:  "Luthur",
		Motto: "Hatred does not heal hatred",
	}
	x := sages{
		Name:  "Jesus",
		Motto: "love all",
	}
	y := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	z := car{
		Manufacturer: "Toyota",
		Model:        "Corrolla",
		Doors:        4,
	}

	try := []sages{a, b, x}
	tryial := []car{y, z}

	datta := struct {
		Wisdom    []sages
		Transport []car
	}{
		try,
		tryial,
	}

	err := tpl.Execute(os.Stdout, datta)
	if err != nil {
		log.Fatalln(err)
	}
}
