package main

import (
	"html/template"
	"net/http"

	"github.com/rkweber/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}
