// In the Warehouse new items were delivered.
// When the EAN is scanned the item will be added to database
// and the the item id will be returned

package warehouse

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/database"
)

// checks if there is an product with that ean. If not an empty product will
// be returned and also an empty ItemID
// if ean is existing the new item will be added and new itemID returned
func StockProduct(ean string) string {

	var newItemID string
	product := database.GetProductByEan(ean)

	if product.EAN == ean {
		newItemID = database.NewItem(product.ProductID)
	}
	return newItemID

}
