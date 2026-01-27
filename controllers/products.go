package controllers

import (
	"net/http"
	"text/template"

	"github.com/rkweber/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}
