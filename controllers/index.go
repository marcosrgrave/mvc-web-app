package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/marcosrgrave/mvc-crud-products/db"
	"github.com/marcosrgrave/mvc-crud-products/models"
	"github.com/marcosrgrave/mvc-crud-products/utils"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {

	err := temp.ExecuteTemplate(w, "HomePage", nil)
	if err != nil {
		utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error executing Home Page template.")
		return
	}
}

func ListProductsPage(w http.ResponseWriter, r *http.Request) {

	db, err := db.DBConn()
	if err != nil {
		utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error establishing database connection.")
		return
	}
	defer db.Close()

	products, err := models.ReadAllProducts(db)
	if err != nil {
		utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error reading product: "+err.Error())
		return
	}

	err = temp.ExecuteTemplate(w, "ListProductsPage", products)
	if err != nil {
		utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error executing List Products Page template.")
		return
	}
}

func NewProductPage(w http.ResponseWriter, r *http.Request) {

	db, err := db.DBConn()
	if err != nil {
		utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error establishing database connection.")
		return
	}
	defer db.Close()

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error converting price to float.")
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error converting quantity to integer.")
		}

		product := models.NewProduct(name, description, priceFloat, quantityInt, "nil")

		err = models.CreateProduct(db, product)
		if err != nil {
			utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error creating new product.")
		}

	}

	err = temp.ExecuteTemplate(w, "NewProductPage", nil)
	if err != nil {
		utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error executing New Product Page template.")
		return
	}
}

func UpdateProductPage(w http.ResponseWriter, r *http.Request) {

	err := temp.ExecuteTemplate(w, "UpdateProductPage", nil)
	if err != nil {
		fmt.Println(err)
		utils.JSONResponseError(w, http.StatusInternalServerError, "Error executing Update Product Page template.", err)
		return
	}

	// if r.Method == http.MethodPut {
	// }

	// err = models.UpdateProduct(db, product)
	// if err != nil {
	// 	utils.JSONResponseMessage(w, http.StatusInternalServerError, "Error updating product.")
	// 	return
	// }

	// utils.JSONResponseMessage(w, http.StatusInternalServerError, "Product updated successfully.")

}
