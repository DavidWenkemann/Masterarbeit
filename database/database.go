//Package Database emulates a Database.
//In fact all the functions are implemented, but the data is simply stored
//in RAM and will be resetted when the application is closed.

package database

import (
	"strconv"
	"time"

	"github.com/DavidWenkemann/Masterarbeit/model"
)

/*
**
Internal variables
**
*/

//Serial Numbers are unique keys in the database
//When a produt/item is created, they are increased by 1
var productSerial int
var itemSerial int

var items = []model.DBItem{}
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

//SpinupDB is a public function that fills up the emulated database.
//It is called from main.go when the application starts up.
func SpinupDB() {

	//todo products hinzufügen
	NewProduct("4011803092174", "Spezi", 0.75)
	NewProduct("4066600641919", "Paulaner Hefeweizen", 1.39)
	NewProduct("4029764001807", "Clubmate", 2.50)
	NewProduct("4102560080068", "Alasia Medium Water", 1)

	//emtpyTime := time.Now().IsZero()

	//fill up item database with several items in the past.
	items = append(items, oldItem(1, time.Now().Add(-24*time.Hour), nil))
	items = append(items, oldItem(1, time.Now().Add(-54*time.Hour), nil))
	items = append(items, oldItem(3, time.Now().Add(-827*time.Hour), nil))
	items = append(items, oldItem(2, time.Now().Add(-46*time.Hour), timePtr(time.Now().Add(-24*time.Hour))))

	//items[0].ItemID = "4005906003427"

	//fmt.Printf("%v", products)
	//fmt.Printf("%v", items)

}

/*
**
Public Functions
**
*/

//add an product to data table
func NewProduct(ean string, name string, price float64) model.BProduct {
	p := newProduct(ean, name, price)
	products = append(products, p) //Insert into DB

	return mapDBProductToBProduct(p)
}

//returns businessmodel of product with specific ean. If not available returns nil
func GetProductByEan(ean string) model.BProduct {
	var p model.BProduct
	p.ProductID = 0
	for i := range products {
		if ean == products[i].EAN {
			p = mapDBProductToBProduct(products[i])
		}
	}
	return p
}

//returns businessmodel of product with specific id. If not available returns nil
func GetProductByID(id int) model.BProduct {
	var p model.BProduct
	for i := range products {
		if id == products[i].ProductID {
			p = mapDBProductToBProduct(products[i])
		}
	}
	return p
}

//removes products with specific ean out of db.
func RemoveProductByEan(ean string) {
	var p model.DBProduct //emtpy product to overwrite the last element
	for i := range products {
		if ean == products[i].EAN {
			products[i] = products[len(products)-1] // Copy last element to index i.
			products[len(products)-1] = p           // Erase last element (write zero value).
			products = products[:len(products)-1]   // Truncate slice.
		}
	}
	p.EAN = ""
}

//public function to add an item to data table. New Product ID will be returned
func NewItem(pID int) string {
	i := newItem(pID)
	items = append(items, i) //Insert into DB
	return i.ItemID
}

func GetItemsInStockByEan(ean string) int {

	var count int
	pID := GetProductByEan(ean).ProductID

	//TODO: funktioniert wohl ncch nicht mit time

	for i := range items {
		if pID == items[i].ProductID && items[i].SellingDate.IsZero() {
			count++
		}
	}

	return count
}

func GetItemById(itemID string) model.BItem {
	var item model.BItem
	for i := range items {
		if itemID == items[i].ItemID {
			item = mapDBItemToBItem(items[i])
		}
	}
	return item
}

//Maps all products to businessproducts and returns them
func GetAllProducts() []model.BProduct {

	var businessProducts []model.BProduct

	for i := range products {
		businessProducts = append(businessProducts, mapDBProductToBProduct(products[i]))
	}

	return businessProducts
}

func SetItemSelledDate(itemID string) {
	for i := range items {
		if itemID == items[i].ItemID {
			items[i].SellingDate = time.Now()
		}
	}
}

/*
**
Helper and Mapping Functions
**
*/

//maps DB-Products to B-Products
func mapDBProductToBProduct(input model.DBProduct) model.BProduct {
	return model.BProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}

func mapDBItemToBItem(input model.DBItem) model.BItem {
	return model.BItem{ProductID: input.ProductID, ItemID: input.ItemID, ReceivingDate: input.ReceivingDate, SellingDate: input.SellingDate}
}

//helper function to convert a time into a pointer
func timePtr(t time.Time) *time.Time {
	return &t
}
