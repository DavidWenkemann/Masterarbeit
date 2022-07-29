//store communication sends no informations, its only receiving

package communication

import (
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/model"

	"encoding/json"
)

var ReportingListenerProducts chan []byte
var ReportingListenerItems chan []byte

func ReportingListener() {

	//Create channels
	ReportingListenerProducts = make(chan []byte, 1024)
	ReportingListenerItems = make(chan []byte, 1024)

	//listener for products
	go func() {
		for {

			select {
			case msg := <-ReportingListenerProducts:

				receivedReportingBProduct := model.BProduct{}

				err := json.Unmarshal(msg, &receivedReportingBProduct)
				if err != nil {
					panic(err)
				}

				//SAVE TO DB
				saveReceivedProduct(receivedReportingBProduct)
				//fmt.Println("(msg)json :", string(msg), "; reporting product go :", receivedReportingBProduct)

				//case <-quit.C:
				//break
			}
		}
	}()

	//listener for items
	go func() {
		for {

			select {
			case msg := <-ReportingListenerItems:

				receivedReportingBItem := model.BItem{}

				err := json.Unmarshal(msg, &receivedReportingBItem)
				if err != nil {
					panic(err)
				}
				//SAVE TO DB
				saveReceivedItem(receivedReportingBItem)

				//case <-quit.C:
				//break
			}
		}
	}()
}

func saveReceivedProduct(receivedProduct model.BProduct) {
	//if new Product
	_, err := database.GetProductByEan(receivedProduct.EAN)

	if err != nil {
		database.ReceiveNewProduct(mapBProductToDBProduct(receivedProduct))
		return
	}

	//if NOT New Product
	//If Not deleted
	if receivedProduct.EAN == "" {
		//panic("Hilfe!")

		database.ReceiveRemoveProduct(receivedProduct.ProductID)
	} else {
		database.ReceiveEditProduct(mapBProductToDBProduct(receivedProduct))

	}

}

//Only new Items possible
func saveReceivedItem(receivedItem model.BItem) {
	database.ReceiveNewItem(mapBItemToDBItem(receivedItem))
}

/*
**
Mapping B Model -> DB Model
**
*/

func mapBItemToDBItem(input model.BItem) model.DBItem {
	return model.DBItem{ProductID: input.ProductID, ItemID: input.ItemID, ReceivingDate: input.ReceivingDate, SellingDate: input.SellingDate}
}

func mapBProductToDBProduct(input model.BProduct) model.DBProduct {
	return model.DBProduct{ProductID: input.ProductID, EAN: input.EAN, Name: input.Name, Price: input.Price}
}
