//package reporting contains the module reporting
//this module is responsible to create an overview of all products
//stored in the warehouse at the moment
//It has its own UI
package reporting

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/model"
)

func GetItemsInStockByEan(ean string) int {
	return database.GetItemsInStockByEan(ean)
}

func GetAllProducts() []model.BProduct {
	return database.GetAllProducts()
}
