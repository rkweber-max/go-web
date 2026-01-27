package main

import (
	"net/http"

	"github.com/rkweber/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
