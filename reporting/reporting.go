//package reporting contains the module reporting
//this module is responsible to create an overview of all products
//stored in the warehouse at the moment
//It has its own UI and handles the database of the softwaresystem

package reporting

import (
	"github.com/DavidWenkemann/Masterarbeit/reporting/database"
	"github.com/DavidWenkemann/Masterarbeit/reporting/model"
)

//loading data from DB
var productsInStock = database.GetAllProducts()

//Maps Product from Datebase-Model to Business-Model
func mapDBProductToBProduct(input model.DBProduct) model.BProduct {

	return model.BProduct{EAN: input.EAN, Name: input.Name, Price: input.Price, Quantity: input.Quantity}

}

//Maps Product from Business-Model to Database-Model
func mapBProductToDBProduct(input model.BProduct) model.DBProduct {

	return model.DBProduct{EAN: input.EAN, Name: input.Name, Price: input.Price, Quantity: input.Quantity}

}

//checks if productname is in database and adds quantity by one
func StockProductByOne(ean string) bool {

	for i := range productsInStock {

		if productsInStock[i].EAN == ean {
			productsInStock[i].Quantity += 1
			return true
		}
	}
	return false
}

//checks if productname is in database and adds quantity by one
func RemoveProductByOne(ean string) bool {

	for i := range productsInStock {

		if productsInStock[i].EAN == ean {
			productsInStock[i].Quantity -= 1
			return true
		}
	}
	return false
}
