package main

import (
	"database/sql"

	db2 "github.com/cassioleodegario/go-hexagonal/adapters/db"
	"github.com/cassioleodegario/go-hexagonal/application"
)

func main() {
	Db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(Db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}
