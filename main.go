package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*.tmpl
var views embed.FS

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./sample/gerow"))

	server.Handle("/", fs)
	server.HandleFunc("/home", homeHandler)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Error while starting the server")
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFS(views, "**/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(w, "home.tmpl", nil)
}
