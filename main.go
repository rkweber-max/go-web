package main

import (
	"html/template"
	"net/http"
)

type Products struct {
	Nome        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
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
