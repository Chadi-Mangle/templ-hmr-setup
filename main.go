package main

import (
	"fmt"
	"net/http"

	assetsfs "github.com/Chadi-Mangle/templ-hmr-setup/assets"
	"github.com/Chadi-Mangle/templ-hmr-setup/templates"
	"github.com/a-h/templ"
)

func main() {
	component := templates.Hello("World !")
	http.Handle("/", templ.Handler(component))

	var httpRoot http.FileSystem
	if assetsfs.IsEmbedded {
		httpRoot = http.FS(assetsfs.Box)
	} else {
		httpRoot = http.Dir("assets")
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(httpRoot)))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
