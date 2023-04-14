package server

import (
	"net/http"

	"github.com/marcosrgrave/mvc-crud-products/routes"
)

func StartServer() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
