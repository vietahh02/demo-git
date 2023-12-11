package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tpl *template.Template

type Result struct {
	ListCate []Category
	F        string
}

type Category struct {
	IdCate   string `json:"id`
	NameCate string `json:"name"`
}

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	var err error
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// db, err = sql.Open("mysql", "root:1234567890@tcp(localhost:3306)/shopcart")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe("localhost:9990", nil))
}

func home(w http.ResponseWriter, r *http.Request) {

	// var send Result
	// send.ListCate = listCategories()
	// send.F = "had"
	// tpl.ExecuteTemplate(w, "header.html", send)
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// func listCategories() []Category {
// 	rows, _ := db.Query("SELECT idCate, NameCate FROM shopcart.categories")
// 	var c1 Category
// 	var ListCate []Category
// 	for rows.Next() {
// 		rows.Scan(&c1.IdCate, &c1.NameCate)
// 		ListCate = append(ListCate, c1)
// 	}
// 	return ListCate
// }
