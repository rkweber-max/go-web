package routes

import (
	"net/http"

	"github.com/rkweber/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
