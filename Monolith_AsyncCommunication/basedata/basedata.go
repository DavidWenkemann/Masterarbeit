// MODULE BASEDATA
//
// Basedata handels all products. Products can be created,
// edited or deleted.

package basedata

import (
	basedatadb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/database"
	storedb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/database"
	warehousedb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/database"

	basedatamodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/model"
	reportingmodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/model"
	storemodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"
	warehousemodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/model"

	communication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/communication"
)

//checks if product with ean is in db. If yes -> edit that and return. If no -> new Product
func AddProduct(ean string, name string, price float64) basedatamodel.BProduct {
	var p basedatamodel.BProduct
	if GetProductByEan(ean).ProductID == 0 {
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

// Get Functions
func GetAllProducts() []basedatamodel.BProduct {
	return mapBasedataDBProductSliceToBasedataBProductSlice(basedatadb.GetAllProducts())
}

func GetProductByEan(ean string) basedatamodel.BProduct {
	return mapBasedataDBProductToBasedataBProduct(basedatadb.GetProductByEan(ean))
}

//SetFunctions
func NewProduct(ean string, name string, price float64) basedatamodel.BProduct {

	newProduct := basedatadb.NewProduct(ean, name, price)

	//send to publish
	communication.SendEditProduct(mapBasedataDBProductToBasedataBProduct(newProduct))
	storedb.ReceiveNewProduct(mapBasedataDBProductToStoreDBProduct(newProduct))
	warehousedb.ReceiveNewProduct(mapBasedataDBProductToWarehouseDBProduct(newProduct))

	return mapBasedataDBProductToBasedataBProduct(newProduct)
}

func RemoveProductByEan(id int) {

	if len(GetAllProducts()) == 0 {
		return
	}

	basedatadb.RemoveProductByEan(id)

	//send to publish
	communication.SendRemovedProduct(id) //reportingdb.ReceiveRemoveProduct(id)

}

func EditProduct(ean string, name string, price float64) basedatamodel.BProduct {
	editedProduct := basedatadb.EditProduct(ean, name, price)

	//send to publish
	communication.SendEditProduct(mapBasedataDBProductToBasedataBProduct(editedProduct))

	return mapBasedataDBProductToBasedataBProduct(editedProduct)
}

/*
**
Mapping DB Model -> B Model
**
*/
func mapBasedataDBProductToBasedataBProduct(input basedatamodel.DBProduct) basedatamodel.BProduct {
	return basedatamodel.BProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}

func mapBasedataDBProductToReportingDBProduct(input basedatamodel.DBProduct) reportingmodel.DBProduct {
	return reportingmodel.DBProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}

func mapBasedataDBProductToStoreDBProduct(input basedatamodel.DBProduct) storemodel.DBProduct {
	return storemodel.DBProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}

func mapBasedataDBProductToWarehouseDBProduct(input basedatamodel.DBProduct) warehousemodel.DBProduct {
	return warehousemodel.DBProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}

func mapBasedataDBProductSliceToBasedataBProductSlice(input []basedatamodel.DBProduct) []basedatamodel.BProduct {

	var output []basedatamodel.BProduct
	for i := range input {
		output = append(output, mapBasedataDBProductToBasedataBProduct(input[i]))
	}
	return output
}
