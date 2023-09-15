package routes

import (
	"net/http"

	"github.com/marcosrgrave/mvc-crud-products/controllers"
)

func LoadRoutes() {
	LoadClient()
	LoadService()
}

func LoadService() {
	http.HandleFunc("/products", controllers.Products)
}

func LoadClient() {
	LoadProductPages()
}

func LoadProductPages() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/products/list", controllers.ListProductsPage)
	http.HandleFunc("/products/new", controllers.NewProductPage)
	http.HandleFunc("/products/update", controllers.UpdateProductPage)
}
