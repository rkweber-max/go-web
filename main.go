package main

import (
	"html/template"
	"net/http"

	"github.com/rkweber/routes"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
