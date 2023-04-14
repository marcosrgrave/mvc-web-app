package controllers

import (
	"log"
	"net/http"

	"github.com/marcosrgrave/mvc-crud-products/db"
)

func Products(w http.ResponseWriter, r *http.Request) {

	db, err := db.DBConn()
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	pc := ProductsController{db, w, r}

	switch r.Method {
	case http.MethodGet:
		pc.Read()
	case http.MethodPost:
		pc.Create()
	case http.MethodPut:
		pc.Update()
	case http.MethodDelete:
		pc.Delete()
	}

}
