//Package Database emulates a Database.
//In fact all the functions are implemented, but the data is simply stored
//in RAM and will be resetted when the application is closed.

package database

import (
	"strconv"
	"time"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/model"
)

/*
**
Internal variables
**
*/

//Serial Numbers are unique keys in the database
//When a produt/item is created, they are increased by 1
var itemSerial int

var items = []model.DBItem{}
var products = []model.DBProduct{} //pruducts emulates a Database table with 4 columns ProductID (unique), EAN (unique) , Name , and Price

/*
**
Internal Functions
**
*/

//internal function. adds item to db (with oldItem-Function) and returns it. Not possible to use a timeSelled
func newItem(pID int) model.DBItem {
	return oldItem(pID, time.Now(), nil)
}

//adds items to database and returns them. Also possible for older items.
func oldItem(pID int, timeReceived time.Time, timeSelled *time.Time) model.DBItem {
	itemSerial++ // change itemserial number, so its always unique
	i := model.DBItem{}
	i.ItemID = strconv.Itoa(itemSerial)
	i.ProductID = pID
	i.ReceivingDate = timeReceived
	if timeSelled != nil {
		i.SellingDate = *timeSelled
	}
	//Print item serial for bottle
	return i
}

/*
**
Public Functions
**
*/

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

//public function to add an item to data table. New Product ID will be returned
func NewItem(pID int) model.DBItem {
	i := newItem(pID)
	items = append(items, i) //Insert into DB
	return i
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
