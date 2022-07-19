package database

import "github.com/DavidWenkemann/Masterarbeit/model"

//pruducts emulates a Database table with 4 columns EAN (unique) , Name , Price ans Quantity
var products = []model.DBProduct{
	{ProductID: 1, EAN: "4011803092174", Name: "Spezi", Price: 0.75},
	{ProductID: 2, EAN: "4066600641919", Name: "Paulaner Weissbier alk.frei", Price: 1.2},
	{ProductID: 3, EAN: "4029764001807", Name: "Clubmate", Price: 2.2},
}

var items = []model.DBItem{
	{ProductID: 1, ItemID: "1"},
}

//maps DB-Products to B-Products and returns them
func GetAllProducts() []model.BProduct {

	var businessProducts []model.BProduct

	for i := range products {
		businessProducts = append(businessProducts, mapDBProductToBProduct(products[i]))
	}

	return businessProducts
}

//maps DB-Products to B-Products
func mapDBProductToBProduct(input model.DBProduct) model.BProduct {

	return model.BProduct{EAN: input.EAN, Name: input.Name, Price: input.Price}

}
