package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	// tpl, _ = template.ParseFiles("index.html", "home.html")
	tpl, _ = template.ParseGlob("*.html")
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func bye(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}
