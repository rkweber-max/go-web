package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Products struct {
	Nome        string
	Description string
	Price       float64
	Quantity    int
}

func connectionDB() *sql.DB {
	connection := "user=maxter dbname=store password=admin host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectionDB()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Products{
		{Nome: "T-Shirt", Description: "Blue", Price: 50, Quantity: 5},
		{"Shoes", "Red", 100, 2},
		{"New Product", "Promotion", 57, 1},
	}

	templates.ExecuteTemplate(w, "Index", products)
}
