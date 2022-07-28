//Package Database emulates a Database.
//In fact all the functions are implemented, but the data is simply stored
//in RAM and will be resetted when the application is closed.

package database

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/model"
)

/*
**
Internal variables
**
*/

//Serial Numbers are unique keys in the database
//When a produt/item is created, they are increased by 1
var productSerial int

var products = []model.DBProduct{} //pruducts emulates a Database table with 4 columns ProductID (unique), EAN (unique) , Name , and Price

/*
**
Internal Functions
**
*/

func newProduct(ean string, name string, price float64) model.DBProduct {
	productSerial++ // change serial number, so its always unique
	p := model.DBProduct{}
	p.ProductID = productSerial
	p.EAN = ean
	p.Name = name
	p.Price = price
	//products = append(products, p)
	return p
}

func EditProduct(ean string, name string, price float64) model.DBProduct {

	p := model.DBProduct{}

	for i := range products {
		if ean == products[i].EAN {
			products[i].Name = name
			products[i].Price = price
			p = products[i]
		}
	}
	return p
}

/*
**
Public Functions
**
*/

//add an product to data table
func NewProduct(ean string, name string, price float64) model.DBProduct {
	p := newProduct(ean, name, price)
	products = append(products, p) //Insert into DB

	return p
}

//returns businessmodel of product with specific ean. If not available returns nil
func GetProductByEan(ean string) model.DBProduct {
	var p model.DBProduct
	p.ProductID = 0
	for i := range products {
		if ean == products[i].EAN {
			p = products[i]
		}
	}
	return p
}

//removes products with specific ean out of db.
func RemoveProductByEan(ean string) {
	var p model.DBProduct //emtpy product to overwrite the last element
	for i := range products {
		if ean == products[i].EAN {

			// Remove the element at index i from product.
			copy(products[i:], products[i+1:])    // Shift a[i+1:] left one index.
			products[len(products)-1] = p         // Erase last element (write zero value).
			products = products[:len(products)-1] // Truncate slice.
		}

	}
	p.EAN = ""
}

//Maps all products to businessproducts and returns them
func GetAllProducts() []model.DBProduct {
	return products
}

/*
**
Helper Functions
**
*/

/*
**
Receive Data from other Modul
**
*/
