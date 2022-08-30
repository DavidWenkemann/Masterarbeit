// MODULE STORE
//
// In the Store its possible to add items to the cart.
// After that the whole cart can be deleted or checked out.

package store

import (
	//"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/database"
	//"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"

	"time"

	storedb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/database"

	communication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/communication"

	storemodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"
)

var cart = []storemodel.BItem{}

//adds product to the cart
func AddToCart(itemID string) {

	/*
		item := GetItemById(itemID)

		if item.ProductID > 0 {
			cart = append(cart, item)
		}
	*/

	if GetProductByID(GetItemById(itemID).ProductID).Name != "" {
		if GetItemById(itemID).SellingDate.IsZero() {
			cart = append(cart, GetItemById(itemID))
		}
	}

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
	var clear []storemodel.BItem
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
func GetCart() []storemodel.BItem {
	return cart
}

/*
**
//Connection to DatabaseLayer
**
*/

//Get functions
func GetProductByID(id int) storemodel.BProduct {
	return mapStoreDBProductToStoreBProduct(storedb.GetProductByID(id))
}

func GetItemById(itemID string) storemodel.BItem {
	return mapStoreDBItemToStoreBItem(storedb.GetItemById(itemID))
}

//Set functions
func SetItemSelledDate(itemID string) {
	sendItem := GetItemById(itemID)
	sendItem.SellingDate = time.Now()
	communication.SendEditItem(sendItem)
}

/*
**
Mapping DB Model -> B Model
**
*/
func mapStoreDBProductToStoreBProduct(input storemodel.DBProduct) storemodel.BProduct {
	return storemodel.BProduct{
		ProductID: input.ProductID,
		EAN:       input.EAN,
		Name:      input.Name,
		Price:     input.Price}
}

func mapStoreDBProductSliceToStoreBProductSlice(input []storemodel.DBProduct) []storemodel.BProduct {

	var output []storemodel.BProduct
	for i := range input {
		output = append(output, mapStoreDBProductToStoreBProduct(input[i]))
	}
	return output
}

func mapStoreDBItemToStoreBItem(input storemodel.DBItem) storemodel.BItem {
	return storemodel.BItem{
		ProductID: input.ProductID,
		ItemID:    input.ItemID}
}

/*
func mapStoreDBItemToReportingDBItem(input storemodel.DBItem) reportingmodel.DBItem {
	return reportingmodel.DBItem{ProductID: input.ProductID, ItemID: input.ItemID}
}
*/
