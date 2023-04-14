package routes

import (
	"net/http"

	"github.com/marcosrgrave/mvc-crud-products/controllers"
)

func LoadRoutes() {
	LoadPages()
	LoadCRUD()
}

func LoadCRUD() {
	http.HandleFunc("/products", controllers.Products)
}

func LoadPages() {
	LoadProductPages()
}

func LoadProductPages() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/products/list", controllers.ListProductsPage)
	http.HandleFunc("/products/new", controllers.NewProductPage)
	http.HandleFunc("/products/update", controllers.UpdateProductPage)
}
