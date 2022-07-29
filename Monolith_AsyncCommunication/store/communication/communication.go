package communication

import (
	reportingcommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/communication"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"

	"encoding/json"
)

var StoreListenerProducts chan []byte
var StoreListenerItems chan []byte

func StoreListener() {

	//Create channels
	StoreListenerProducts = make(chan []byte, 1024)
	StoreListenerItems = make(chan []byte, 1024)

	//listener for products
	go func() {
		for {

			select {
			case msg := <-StoreListenerProducts:

				receivedStoreBProduct := model.BProduct{}

				err := json.Unmarshal(msg, &receivedStoreBProduct)
				if err != nil {
					panic(err)
				}

				//SAVE TO DB
				saveReceivedProduct(receivedStoreBProduct)

				//case <-quit.C:
				//break
			}
		}
	}()

	//listener for items
	go func() {
		for {

			select {
			case msg := <-StoreListenerItems:

				receivedStoreBItem := model.BItem{}

				err := json.Unmarshal(msg, &receivedStoreBItem)
				if err != nil {
					panic(err)
				}

				saveReceivedItem(receivedStoreBItem)

				//case <-quit.C:
				//break
			}
		}
	}()
}

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

//Only new Items possible
func saveReceivedItem(receivedItem model.BItem) {
	database.ReceiveNewItem(mapBItemToDBItem(receivedItem))
}

func SendEditItem(sendItem model.BItem) {

	//serialize product
	serializedSendItem, err := json.Marshal(sendItem)
	if err != nil {
		panic(err)
	}

	//send products to all interested channels
	reportingcommunication.ReportingListenerItems <- serializedSendItem
}

/*
**
Mapping B Model -> DB Model
**
*/

func mapBItemToDBItem(input model.BItem) model.DBItem {
	return model.DBItem{ProductID: input.ProductID, ItemID: input.ItemID, ReceivingDate: input.ReceivingDate}
}

func mapBProductToDBProduct(input model.BProduct) model.DBProduct {
	return model.DBProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}
