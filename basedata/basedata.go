package main

/*

import (
	"github.com/DavidWenkemann/Masterarbeit/basedata/model"
)

/*
type Product struct {
	EAN   string  `json:"ean"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
*/

/*
func mapDBProductToBProduct(input model.DBProduct) model.BProduct {

	return model.BProduct{EAN: input.EAN, Name: input.Name, Price: input.Price}

}

func mapBProductToDBProduct(input model.BProduct) model.DBProduct {

	return model.DBProduct{EAN: input.EAN, Name: input.Name, Price: input.Price}

}

//var products = mapDBProductToBProduct(database.GetAllProducts())

type baseDataBusiness struct{}

func GetProductByEAN(key string) Product {
	return bdb.GetProductByEAN(key)
}

func (bdb *baseDataBusiness) GetProductByEAN(key string) Product {

	var p Product

	//searches product with same ean in slice and returns it
	for i := range products {

		if products[i].EAN == key {
			return products[i]
		}
	}

	//returns an empty product if ean wasnt found
	return p

}

//returns all products in the slice
func (bdb *baseDataBusiness) GetAllProducts() []Product {
	return products
}

//Adds a new product to the slice, quanity is 0
func (bdb *baseDataBusiness) AddProduct(ean string, name string, price float64) Product {

	//TODO: wenn ean schon vorhanden -> Fehler

	//Initialize new Product
	p := Product{
		EAN:   ean,
		Name:  name,
		Price: price,
	}

	//Add new Product to the slice
	products = append(products, p)

	//return the new product
	return p
}

//removes specific product, changes sorting of slice
func (bdb *baseDataBusiness) RemoveProductByEAN(key string) bool {

	//searchs product with same ean in slice
	for i := range products {
		if products[i].EAN == key {

			//replacing found product with last product in slice
			products[i] = products[len(products)-1]
			//shorten up slice by 1
			products = products[:len(products)-1]
			return true
		}
	}
	return false
}

*/
