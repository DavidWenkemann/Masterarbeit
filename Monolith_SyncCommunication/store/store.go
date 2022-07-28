package store

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/model"
)

var cart = []model.BItem{}

//adds product to the cart
func AddToCart(itemID string) {
	cart = append(cart, GetItemById(itemID))
}

//Sets SellingDate for everything in on time now and clears cart
func SellCart() {
	for i := range cart {
		SetItemSelledDate(cart[i].ItemID)
	}
	ClearCart()
}

//removes everything of the cart
func ClearCart() {
	var clear []model.BItem
	cart = clear
}

func GetPriceOfCart() float64 {
	var price float64
	for i := range cart {
		price += GetProductByID(cart[i].ProductID).Price
	}
	return price
}

//Converts Cart to API cart and returns it
func GetCart() []model.BItem {
	//var cartAPI []model.APIItem
	//for i := range cart {
	//	cartAPI = append(cartAPI, mapBItemToAPIItem(cart[i]))
	//}
	return cart
}

/*
**
//Connection to DatabaseLayer
**
*/
func GetProductByID(id int) model.BProduct {
	return mapDBProductToBProduct(database.GetProductByID(id))
}

func SetItemSelledDate(itemID string) {
	database.SetItemSelledDate(itemID)
}

func GetItemById(itemID string) model.BItem {
	return mapDBItemToBItem(database.GetItemById(itemID))
}

/*
**
Mapping DB Model -> B Model
**
*/
func mapDBProductToBProduct(input model.DBProduct) model.BProduct {
	return model.BProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}

func mapDBProductSliceToBProductSlice(input []model.DBProduct) []model.BProduct {

	var output []model.BProduct
	for i := range input {
		output = append(output, mapDBProductToBProduct(input[i]))
	}
	return output
}

func mapDBItemToBItem(input model.DBItem) model.BItem {
	return model.BItem{ProductID: input.ProductID, ItemID: input.ItemID, ReceivingDate: input.ReceivingDate, SellingDate: input.SellingDate}
}
