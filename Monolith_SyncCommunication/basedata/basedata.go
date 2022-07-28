//basedata package is responsible for adding and removing basedata.
//It is connected to DB and UI Layer

package basedata

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/model"
)

//checks if product with ean is in db. If yes -> edit that and return. If no -> new Product
func AddProduct(ean string, name string, price float64) model.BProduct {
	var p model.BProduct
	if database.GetProductByEan(ean).ProductID == 0 {
		p = NewProduct(ean, name, price)
	} else {
		EditProduct(ean, name, price)
	}
	return p
}

/*
**
//Connection to DatabaseLayer
**
*/
func GetAllProducts() []model.BProduct {
	return mapDBProductSliceToBProductSlice(database.GetAllProducts())
}

func GetProductByEan(ean string) model.BProduct {
	return mapDBProductToBProduct(database.GetProductByEan(ean))
}

func NewProduct(ean string, name string, price float64) model.BProduct {
	return mapDBProductToBProduct(database.NewProduct(ean, name, price))
}

func RemoveProductByEan(ean string) {
	database.RemoveProductByEan(ean)
}

func EditProduct(ean string, name string, price float64) model.BProduct {
	return mapDBProductToBProduct(database.EditProduct(ean, name, price))
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
