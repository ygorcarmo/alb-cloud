package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates sample
var assets embed.FS

var tmpl map[string]*template.Template

func init() {
	loadTemplates()
}

func main() {
	server := http.NewServeMux()

	// only serve the front end files
	folder, fserr := fs.Sub(assets, "sample/gerow")

	if fserr != nil {
		fmt.Println(fserr)
	}

	fs := http.FileServer(http.FS(folder))
	server.Handle("/", fs)

	server.HandleFunc("/home", homeHandler)
	server.HandleFunc("/about-us", aboutHandler)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Error while starting the server")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl["home"].ExecuteTemplate(w, "base-layout.tmpl", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl["about"].ExecuteTemplate(w, "base-layout.tmpl", nil)
}

func loadTemplates() {
	tmpl = make(map[string]*template.Template)

	templateFolder, err := fs.Sub(assets, "templates")

	if err != nil {
		panic(err)
	}

	tmpl["home"] = template.Must(template.ParseFS(templateFolder, "base-layout.tmpl", "home.tmpl"))
	tmpl["custom"] = template.Must(template.ParseFS(templateFolder, "base-layout.tmpl", "custom.tmpl"))
	tmpl["about"] = template.Must(template.ParseFS(templateFolder, "base-layout.tmpl", "about.tmpl"))

}
