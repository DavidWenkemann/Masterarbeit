//communication of basedata is only sending informations
//there is no receiving informations

package communication

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/database"

	reportingcommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/communication"
	storecommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/communication"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/model"

	"encoding/json"
)

var WarehouseListenerProducts chan []byte

func WarehouseListener() {

	WarehouseListenerProducts = make(chan []byte, 1024)

	//listener for products
	go func() {
		for {

			select {
			case msg := <-WarehouseListenerProducts:

				receivedWarehouseBProduct := model.BProduct{}

				err := json.Unmarshal(msg, &receivedWarehouseBProduct)
				if err != nil {
					panic(err)
				}

				//SAVE TO DB
				saveReceivedProduct(receivedWarehouseBProduct)

				//case <-quit.C:
				//	break
			}
		}
	}()

}

/*
**
Save Functions
**
*/

func saveReceivedProduct(receivedProduct model.BProduct) {
	//if new Product
	if database.GetProductByEan(receivedProduct.EAN).ProductID == 0 {
		database.ReceiveNewProduct(mapBProductToDBProduct(receivedProduct))
	}

	//if NOT New Product
	if database.GetProductByEan(receivedProduct.EAN).ProductID != 0 {
		//If Not deleted
		if database.GetProductByEan(receivedProduct.EAN).EAN != "" {
			database.ReceiveEditProduct(mapBProductToDBProduct(receivedProduct))
		} else { //deleted
			database.ReceiveRemoveProduct(receivedProduct.ProductID)
		}

	}
}

/*
**
Send Functions
**
*/
func SendNewItem(sendItem model.BItem) {

	//serialize product
	serializedSendItem, err := json.Marshal(sendItem)
	if err != nil {
		panic(err)
	}

	//send products to all interested channels
	reportingcommunication.ReportingListenerItems <- serializedSendItem
	storecommunication.StoreListenerItems <- serializedSendItem

}

/*
**
Mapping B Model -> DB Model
**
*/

func mapBProductToDBProduct(input model.BProduct) model.DBProduct {
	return model.DBProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}
