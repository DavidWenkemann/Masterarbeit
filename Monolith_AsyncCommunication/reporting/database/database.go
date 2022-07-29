//Package Database emulates a Database.
//In fact all the functions are implemented, but the data is simply stored
//in RAM and will be resetted when the application is closed.

package database

import (
	"errors"
	"time"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/model"
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
Internal Functions
**
*/

/*
**
Public Functions
**
*/

//returns businessmodel of product with specific ean. If not available returns nil
func GetProductByEan(ean string) (model.DBProduct, error) {
	var p model.DBProduct
	p.ProductID = 0
	for i := range products {
		if ean == products[i].EAN {
			p = products[i]
		}
	}
	if p.ProductID == 0 {
		return p, errors.New("EAN Not Found")
	}

	return p, nil
}

//Maps all products to businessproducts and returns them
func GetAllProducts() []model.DBProduct {
	return products
}

func GetAllItems() []model.DBItem {
	return items
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

func ReceiveRemoveProduct2(id int) {
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

func ReceiveRemoveProduct(id int) {

	s := -1
	for i := range products {
		if id == products[i].ProductID {
			s = i
			break
		}
	}
	if s >= 0 {
		products = append(products[:s], products[s+1:]...)

	}
	return
}

func ReceiveEditItem(item model.DBItem) {
	for i := range items {
		if item.ItemID == items[i].ItemID {
			items[i] = item
		}
	}
}

func ReceiveNewItem(newItem model.DBItem) {
	items = append(items, newItem)
}

func ReceiveSellItem(itemid string, selledTime time.Time) {
	for i := range items {
		if itemid == items[i].ItemID {
			items[i].SellingDate = selledTime
		}
	}

}
