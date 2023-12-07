package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template
var db *sql.DB
var err error

func main() {
	tpl, _ = template.ParseGlob("*.html")
	db, err = sql.Open("mysql", "root:1234567890@tcp(localhost:3306)/vietnam")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":9990", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	row, _ := db.Query("SELECT name_en FROM vietnam.districts where code = '001'")
	var name string
	for row.Next() {
		row.Scan(&name)
	}
	tpl.ExecuteTemplate(w, "index.html", name)
}

func bye(w http.ResponseWriter, r *http.Request) {
	row, _ := db.Query("SELECT NameCate FROM shopcart.categories where idCate = '1'")
	var name string
	for row.Next() {
		row.Scan(&name)
	}
	tpl.ExecuteTemplate(w, "home.html", name)
}
