package main

import (
	"fmt"
	"net/http"

	"github.com/StevenYeu/templ-play/templates"
	"github.com/a-h/templ"
)

func main() {
	comp := templates.Hello("Steven")
	http.Handle("/", templ.Handler(comp))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
