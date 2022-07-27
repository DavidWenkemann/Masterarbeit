package store

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/model"
)

//var cart []model.BItem

var cart = []model.BItem{
	//{ProductID: 1, ItemID: "2", ReceivingDate: time.Now()},
	//{ProductID: 2, ItemID: "2", ReceivingDate: time.Now()},
}

//adds product to the cart
func AddToCart(itemID string) {
	cart = append(cart, database.GetItemById(itemID))
}

//Sets SellingDate for everything in on time now and clears cart
func SellCart() {
	for i := range cart {
		database.SetItemSelledDate(cart[i].ItemID)
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
		price += database.GetProductByID(cart[i].ProductID).Price
	}
	return price
}

//Converts Cart to API cart and returns it
func GetCart() []model.APIItem {
	var cartAPI []model.APIItem
	for i := range cart {
		cartAPI = append(cartAPI, mapBItemToAPIItem(cart[i]))
	}
	return cartAPI
}

func mapBItemToAPIItem(input model.BItem) model.APIItem {
	return model.APIItem{Product: mapBProductToAPIroduct(database.GetProductByID(input.ProductID)), ItemID: input.ItemID, ReceivingDate: input.ReceivingDate, SellingDate: input.SellingDate}
}

func mapBProductToAPIroduct(input model.BProduct) model.APIProduct {
	return model.APIProduct{EAN: input.EAN, Name: input.Name, Price: input.Price}
}
