package main

import (
	"github.com/marcosrgrave/mvc-crud-products/db"
	"github.com/marcosrgrave/mvc-crud-products/server"
)

func main() {
	db.Init()
	server.StartServer()
}
