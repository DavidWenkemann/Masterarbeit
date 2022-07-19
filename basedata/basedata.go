//basedata package is responsible for adding and removing basedata.
//Its connected to DB and UI

package main

import (
	"github.com/DavidWenkemann/Masterarbeit/database"
	"github.com/DavidWenkemann/Masterarbeit/model"
)

//adds new product to DB
func addProduct(ean string, name string, price float64) model.BProduct {
	var p model.BProduct
	if database.GetProductByEan(ean) != nil {
		database.NewProduct(ean, name, price)
	}
	return p
}

//Removes specific product from DB by EAN
func removeProduct(ean string) {
	database.RemoveProductByEan(ean)
}
