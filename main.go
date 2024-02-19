package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates assets
var assets embed.FS

var tmpl map[string]*template.Template

func init() {
	loadTemplates()
}

func main() {
	server := http.NewServeMux()

	// only serve the front end files
	folder, fserr := fs.Sub(assets, "assets")

	if fserr != nil {
		fmt.Println(fserr)
	}

	fs := http.FileServer(http.FS(folder))
	server.Handle("/assets/", http.StripPrefix("/assets/", fs))

	server.HandleFunc("/", homeHandler)
	server.HandleFunc("/about-us", aboutHandler)
	server.HandleFunc("/services", servicesHandler)
	server.HandleFunc("/contact", contactHandler)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Error while starting the server")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		tmpl["not-found"].ExecuteTemplate(w, "base-layout.tmpl", nil)
		return
	}

	tmpl["home"].ExecuteTemplate(w, "base-layout.tmpl", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl["about"].ExecuteTemplate(w, "base-layout.tmpl", nil)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl["services"].ExecuteTemplate(w, "base-layout.tmpl", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl["contact"].ExecuteTemplate(w, "base-layout.tmpl", nil)
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
	tmpl["services"] = template.Must(template.ParseFS(templateFolder, "base-layout.tmpl", "services.tmpl"))
	tmpl["contact"] = template.Must(template.ParseFS(templateFolder, "base-layout.tmpl", "contact.tmpl"))
	tmpl["not-found"] = template.Must(template.ParseFS(templateFolder, "base-layout.tmpl", "not-found.tmpl"))

}
