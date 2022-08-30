// MODULE WAREHOUSE
//
// In the Warehouse new items were delivered.
// It is possible to scan EAN of the items. The items
// will be added to the database.

package warehouse

import (
	communication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/communication"
	warehousedb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/database"

	reportingmodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/model"
	storemodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"
	warehousemodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/model"
)

// checks if there is an product with that ean. If not an empty product will
// be returned and also an empty ItemID
// if ean is existing the new item will be added and new itemID returned
func StockProduct(ean string) warehousemodel.BItem {

	var newItem warehousemodel.BItem
	product := GetProductByEan(ean)

	if product.EAN == ean {
		newItem = NewItem(product.ProductID)
	}
	return newItem

}

/*
**
//Connection to DatabaseLayer
**
*/

//Get Functions
func GetProductByEan(ean string) warehousemodel.BProduct {
	return mapWarehouseDBProductToWarehouseBProduct(warehousedb.GetProductByEan(ean))
}

//SetFunctions
func NewItem(pID int) warehousemodel.BItem {
	newItem := warehousedb.NewItem(pID)

	//Send to other DBs
	communication.SendNewItem(mapWarehouseDBItemToWarehouseBItem(newItem))

	return mapWarehouseDBItemToWarehouseBItem(newItem)
}

/*
**
Mapping DB Model -> B Model
**
*/
func mapWarehouseDBProductToWarehouseBProduct(input warehousemodel.DBProduct) warehousemodel.BProduct {
	return warehousemodel.BProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}

func mapWarehouseDBItemToWarehouseBItem(input warehousemodel.DBItem) warehousemodel.BItem {
	return warehousemodel.BItem{ProductID: input.ProductID, ItemID: input.ItemID, ReceivingDate: input.ReceivingDate}
}

func mapWarehouseDBItemToReportingDBItem(input warehousemodel.DBItem) reportingmodel.DBItem {
	return reportingmodel.DBItem{ProductID: input.ProductID, ItemID: input.ItemID, ReceivingDate: input.ReceivingDate}
}

func mapWarehouseDBItemToStoreDBItem(input warehousemodel.DBItem) storemodel.DBItem {
	return storemodel.DBItem{ProductID: input.ProductID, ItemID: input.ItemID}
}
