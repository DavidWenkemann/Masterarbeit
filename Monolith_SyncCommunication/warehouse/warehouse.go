// In the Warehouse new items were delivered.
// When the EAN is scanned the item will be added to database
// and the the item id will be returned

package warehouse

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/model"
)

// checks if there is an product with that ean. If not an empty product will
// be returned and also an empty ItemID
// if ean is existing the new item will be added and new itemID returned
func StockProduct(ean string) string {

	var newItemID string
	product := GetProductByEan(ean)

	if product.EAN == ean {
		newItemID = NewItem(product.ProductID)
	}
	return newItemID

}

/*
**
//Connection to DatabaseLayer
**
*/
func GetProductByEan(ean string) model.BProduct {
	return mapDBProductToBProduct(database.GetProductByEan(ean))
}

func NewItem(pID int) string {
	return database.NewItem(pID)
}

/*
**
Mapping DB Model -> B Model
**
*/
func mapDBProductToBProduct(input model.DBProduct) model.BProduct {
	return model.BProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}
