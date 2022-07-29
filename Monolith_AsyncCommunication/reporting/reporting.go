//package reporting contains the module reporting
//this module is responsible to create an overview of all products
//stored in the warehouse at the moment
//It has its own UI
package reporting

import (
	"fmt"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/model"
)

func GetItemsInStockByEan(ean string) int {
	var count int
	items := GetAllItems()
	pID := GetProductByEan(ean).ProductID

	//TODO: funktioniert wohl ncch nicht mit time

	for i := range items {
		if pID == items[i].ProductID && items[i].SellingDate.IsZero() {
			count++
		}
	}

	return count
}

/*
**
Connection to Database Layer
**
*/
func GetAllProducts() []model.BProduct {
	return mapDBProductSliceToBProductSlice(database.GetAllProducts())
}

func GetAllItems() []model.BItem {
	return mapDBItemSliceToItemSlice(database.GetAllItems())
}

func GetProductByEan(ean string) model.DBProduct {
	product, err := database.GetProductByEan(ean)
	if err != nil {
		fmt.Println(err)
	}
	return product

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

func mapDBItemSliceToItemSlice(input []model.DBItem) []model.BItem {

	var output []model.BItem
	for i := range input {
		output = append(output, mapDBItemToBItem(input[i]))
	}
	return output
}
