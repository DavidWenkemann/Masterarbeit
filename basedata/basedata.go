package main

import (
	"github.com/DavidWenkemann/Masterarbeit/database"
	"github.com/DavidWenkemann/Masterarbeit/model"
)

func addProduct(ean string, name string, price float64) model.BProduct {
	var p model.BProduct
	if database.GetProductByEan(ean) != nil {
		database.NewProduct(ean, name, price)
	}
	return p
}

func removeProduct(ean string) {
	database.RemoveProductByEan(ean)
}
