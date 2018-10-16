package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8000", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}

// this is about creating the token.
// introduction to third part token providers.
