package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("Templ!")
	http.Handle("/", templ.Handler(component))

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on :3000 🚀")
	http.ListenAndServe(":3000", nil)
}
