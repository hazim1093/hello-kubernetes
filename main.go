package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseFiles("public/index.html"))

	http.HandleFunc("/", rootHandler)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
