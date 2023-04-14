package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/marcosrgrave/mvc-crud-products/models"
	"github.com/marcosrgrave/mvc-crud-products/utils"
)

type ProductsController struct {
	db *sql.DB
	w  http.ResponseWriter
	r  *http.Request
}

func (pc *ProductsController) Read() {

	params := pc.r.URL.Query()
	productID := params.Get("id")

	var result []*models.Product

	if productID != "" {
		product, err := models.ReadProduct(pc.db, productID)
		if err != nil {
			utils.JSONResponseError(pc.w, http.StatusInternalServerError, "Product read failed.", err)
			return
		}

		result = append(result, product)

	} else {
		products, err := models.ReadAllProducts(pc.db)
		if err != nil {
			utils.JSONResponseError(pc.w, http.StatusInternalServerError, "Products read failed.", err)
			return
		}

		result = products
	}

	data := map[string]interface{}{
		"products": result,
	}

	utils.JSONResponseData(pc.w, http.StatusNotFound, data)
}

func (pc *ProductsController) Create() {

	var productDTO models.ProductDTOInput

	err := json.NewDecoder(pc.r.Body).Decode(&productDTO)
	if err != nil {
		utils.JSONResponseMessage(pc.w, http.StatusInternalServerError, "JSON decode failed.")
		return
	}

	product := models.NewProduct(productDTO.Name, productDTO.Description, productDTO.Price, productDTO.Quantity, productDTO.CreatorID)

	err = models.CreateProduct(pc.db, product)
	if err != nil {
		utils.JSONResponseMessage(pc.w, http.StatusInternalServerError, "Create product failed.")
		return
	}

	utils.JSONResponseMessage(pc.w, http.StatusCreated, "Product created successfully.")
}

func (pc *ProductsController) Update() {
	var product *models.Product

	err := json.NewDecoder(pc.r.Body).Decode(&product)
	if err != nil {
		utils.JSONResponseMessage(pc.w, http.StatusInternalServerError, "JSON decoder failed.")
		return
	}

	models.UpdateProduct(pc.db, product)

	utils.JSONResponseMessage(pc.w, http.StatusOK, "Product updated successfully.")
}

func (pc *ProductsController) Delete() {
	utils.JSONResponseMessage(pc.w, http.StatusNotFound, "Delete method not implemented yet.")
}
