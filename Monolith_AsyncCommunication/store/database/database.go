//Package Database emulates a Database.
//In fact all the functions are implemented, but the data is simply stored
//in RAM and will be resetted when the application is closed.

package database

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"
)

/*
**
Internal variables
**
*/

var items = []model.DBItem{}
var products = []model.DBProduct{} //pruducts emulates a Database table with 4 columns ProductID (unique), EAN (unique) , Name , and Price

/*
**
Public Functions
**
*/

//returns businessmodel of product with specific id. If not available returns nil
func GetProductByID(id int) model.DBProduct {
	var p model.DBProduct
	for i := range products {
		if id == products[i].ProductID {
			p = products[i]
		}
	}
	return p
}

func GetItemById(itemID string) model.DBItem {
	var item model.DBItem
	itemID = ""
	for i := range items {
		if itemID == items[i].ItemID {
			item = items[i]
		}
	}
	return item
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

/*
func SetItemSelledDate(itemID string) model.DBItem {

	var item model.DBItem
	for i := range items {
		if itemID == items[i].ItemID {
			items[i].SellingDate = time.Now()
			item = items[i]
		}
	}

	return item
}
*/

/*
**
Receive Data from other Modul
**
*/
func ReceiveNewProduct(newProduct model.DBProduct) {
	products = append(products, newProduct)
}

func ReceiveEditProduct(editedProduct model.DBProduct) {
	for i := range products {
		if editedProduct.EAN == products[i].EAN {
			products[i] = editedProduct
		}
	}
}

func ReceiveRemoveProduct(id int) {
	var p model.DBProduct //emtpy product to overwrite the last element
	for i := range products {
		if id == products[i].ProductID {

			copy(products[i:], products[i+1:])    // shift valuesafter the indexwith a factor of 1
			products[len(products)-1] = p         // remove element
			products = products[:len(products)-1] // truncateslice

			return
		}

	}
}

func ReceiveNewItem(newItem model.DBItem) {
	items = append(items, newItem)
}
