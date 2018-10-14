package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundunce", abundunce)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8000", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #1:", c)
	}

	b, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #2:", b)
	}

	a, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #3:", a)
	}
}

func abundunce(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(res, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}
