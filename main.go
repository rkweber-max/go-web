package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Products struct {
	Id          int
	Name        string
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
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectionDB()

	getAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Products{}
	products := []Products{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	templates.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}
