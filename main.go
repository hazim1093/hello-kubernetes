package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

type TemplateData struct {
	Hostname   string
	TestEnvVar string
}

func main() {
	tmpl = template.Must(template.ParseFiles("public/index.html"))

	http.HandleFunc("/", rootHandler)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	data := TemplateData{
		Hostname: os.Getenv("HOSTNAME"),
		TestEnvVar: os.Getenv("TEST_VAR"),
	}

	tmpl.Execute(w, data)
}
