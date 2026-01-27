package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/rkweber/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error to convert price to float", err)
		}

		convertQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error to convert quantity to int", err)
		}

		models.CreateNewProducts(name, description, convertPrice, convertQuantity)
	}
	http.Redirect(w, r, "/", 301)
}
