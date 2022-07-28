//basedata package is responsible for adding and removing basedata.
//It is connected to DB and UI Layer

package basedata

import (
	basedatadb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/database"
	reportingdb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/database"
	storedb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/database"
	warehousedb "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/database"

	basedatamodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/model"
	reportingmodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/model"
	storemodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"
	warehousemodel "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/model"
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

	//change in other DB / send to publisher
	reportingdb.ReceiveNewProduct(mapBasedataDBProductToReportingDBProduct(newProduct))
	storedb.ReceiveNewProduct(mapBasedataDBProductToStoreDBProduct(newProduct))
	warehousedb.ReceiveNewProduct(mapBasedataDBProductToWarehouseDBProduct(newProduct))

	return mapBasedataDBProductToBasedataBProduct(newProduct)
}

func RemoveProductByEan(ean string) {
	basedatadb.RemoveProductByEan(ean)

	//TODO: Connection to other DBs
	reportingdb.ReceiveRemoveProduct(ean)
	//store
	//Warehouse
}

func EditProduct(ean string, name string, price float64) basedatamodel.BProduct {
	editedProduct := basedatadb.EditProduct(ean, name, price)

	//TODO: Connection to other DBs
	reportingdb.ReceiveEditProduct(mapBasedataDBProductToReportingDBProduct(editedProduct))
	storedb.ReceiveEditProduct(mapBasedataDBProductToStoreDBProduct(editedProduct))
	warehousedb.ReceiveEditProduct(mapBasedataDBProductToWarehouseDBProduct(editedProduct))

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
