//basedata package is responsible for adding and removing basedata.
//It is connected to DB and UI

package basedata

import (
	"github.com/DavidWenkemann/Masterarbeit/database"
	"github.com/DavidWenkemann/Masterarbeit/model"
)

//adds new product to DB
func AddProduct(ean string, name string, price float64) model.BProduct {
	var p model.BProduct
	//checks if there is an product with that ean. If not an empty product will
	//be returned and the new product will be added
	if database.GetProductByEan(ean) != p {
		p = database.NewProduct(ean, name, price)
	}
	return p
}

//Removes specific product from DB by EAN
func RemoveProduct(ean string) {
	database.RemoveProductByEan(ean)
}
