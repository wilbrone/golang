package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("ttm.gohtml"))
}

/*This is where you define the format of time the you wish to be displayed*/
func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"fdateDMY": dayMonthYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "ttm.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
