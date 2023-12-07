package main

import (
	m "Shopping/model"
	"database/sql"
	"html/template"
	"log"
	"math"
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
	rows, _ := db.Query("SELECT idPro, NamePro, Price, Discount, Star, Image FROM shopcart.products;")
	defer rows.Close()
	var p1 m.Product
	var ListPro []m.Product
	for rows.Next() {
		rows.Scan(&p1.Id, &p1.Name, &p1.Price, &p1.Discount, &p1.Star, &p1.Image)
		p1.Discount = math.Round(p1.Price*(1-p1.Discount)*100) / 100
		ListPro = append(ListPro, p1)
	}
	// ListCate := listCategories()
	rows, _ = db.Query("SELECT idBanner, Image FROM shopcart.banner;")
	var b m.Banner
	var ListBanner []m.Banner
	for rows.Next() {
		rows.Scan(&b.IdBannerImg, &b.Image)
		ListBanner = append(ListBanner, b)
	}
	rows, _ = db.Query("SELECT idFB, Title, Content FROM shopcart.footer_banner;")
	var fb m.FooterBanner
	var ListFB []m.FooterBanner
	for rows.Next() {
		rows.Scan(&fb.IdBanner, &fb.Title, &fb.Content)
		ListFB = append(ListFB, fb)
	}
	rows, _ = db.Query("SELECT idPro, NamePro, Price, Discount, Star, Image FROM shopcart.products order by Discount desc limit 5")
	var listDeals []m.Product
	for rows.Next() {
		rows.Scan(&p1.Id, &p1.Name, &p1.Price, &p1.Discount, &p1.Star, &p1.Image)
		p1.Discount = math.Round(p1.Price*(1-p1.Discount)*100) / 100
		listDeals = append(listDeals, p1)
	}

	rows, _ = db.Query("SELECT idPro, NamePro, Price, Discount, Star, Image FROM shopcart.products order by Date desc limit 7")
	var listNews []m.Product
	for rows.Next() {
		rows.Scan(&p1.Id, &p1.Name, &p1.Price, &p1.Discount, &p1.Star, &p1.Image)
		p1.Discount = math.Round(p1.Price*(1-p1.Discount)*100) / 100
		listNews = append(listNews, p1)
	}

	rows, _ = db.Query("SELECT idPro, NamePro, Price, Discount, Star, Image FROM shopcart.products order by Date asc limit 7")
	var listLatest []m.Product
	for rows.Next() {
		rows.Scan(&p1.Id, &p1.Name, &p1.Price, &p1.Discount, &p1.Star, &p1.Image)
		p1.Discount = math.Round(p1.Price*(1-p1.Discount)*100) / 100
		listLatest = append(listLatest, p1)
	}

	rows, _ = db.Query("SELECT idPro, NamePro, Price, Discount, Star, Image FROM shopcart.products order by Sold desc limit 5")
	var listSpecials []m.Product
	for rows.Next() {
		rows.Scan(&p1.Id, &p1.Name, &p1.Price, &p1.Discount, &p1.Star, &p1.Image)
		p1.Discount = math.Round(p1.Price*(1-p1.Discount)*100) / 100
		listSpecials = append(listSpecials, p1)
	}

	type list struct {
		ListPro    []m.Product
		ListCate   []m.Category
		ListBanner []m.Banner
		ListFB     []m.FooterBanner
		Deals      []m.Product
		News       []m.Product
		Specials   []m.Product
		Latests    []m.Product
		S          []string
	}
	// var list1 = list{ListPro: ListPro, ListCate: ListCate, ListBanner: ListBanner, ListFB: ListFB, Deals: listDeals, News: listNews, Specials: listSpecials, Latests: listLatest}
	var list2 list
	list2.ListPro = append(list2.ListPro, ListPro...)
	// list2.ListCate = append(list2.ListCate, ListCate...)
	list2.ListBanner = append(list2.ListBanner, ListBanner...)
	list2.ListFB = append(list2.ListFB, ListFB...)
	list2.Deals = append(list2.Deals, listDeals...)
	list2.News = append(list2.News, listNews...)
	list2.Specials = append(list2.Specials, listSpecials...)
	list2.Latests = append(list2.Latests, listLatest...)
	data := []string{"Chuỗi 1", "Chuỗi 2", "Chuỗi 3", "Chuỗi 4", "Chuỗi 5"}
	list2.S = append(list2.S, data...)
	tpl.ExecuteTemplate(w, "home.html", list2)
}
