package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/sessions"
)

var db *sql.DB
var tpl *template.Template
var tplMail *template.Template

type Result struct {
	T string
	F string
}

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	tplMail, _ = tpl.ParseGlob("mail/*.html")
	var err error
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	db, err = sql.Open("mysql", "root:1234567890@tcp(localhost:3306)/shopcart")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe("localhost:9990", nil))
}

func home(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "home.html", nil)
}
