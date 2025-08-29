package main

import (
	"fmt"
	"net/http"

	"embed"

	"github.com/a-h/templ"

	"github.com/Chadi-Mangle/templ-hmr-setup/templates"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	component := templates.Hello("Chadi")
	http.Handle("/", templ.Handler(component))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assetsFS))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
