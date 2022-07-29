//communication of basedata is only sending informations
//there is no receiving informations

package communication

import (
	//"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/database"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/basedata/model"
	reportingcommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/communication"
	storecommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/communication"
	warehousecommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/communication"

	"encoding/json"
)

/*
func ReportingListener() {

	basedataListenerProducts := make(chan []byte, 1024)

	//listener for products
	//go func() {
	for {

		select {
		case msg := <-basedataListenerProducts:

			receivedBasedataBProduct := model.BProduct{}

			err := json.Unmarshal(msg, &receivedBasedataBProduct)
			if err != nil {
				panic(err)
			}
			//SAVE TO DB

			//saveReceivedProduct(receivedBasedataBProduct)

			//fmt.Println("(msg)json :", string(msg), "; reporting product go :", receivedBasedataBProduct)

			//case <-quit.C:
			//	break
		}
	}
	//}

}
*/

/*
**
Save Functions
**
*/

/*
func saveReceivedProduct(receivedProduct model.BProduct){

	//if new Product
	if(database.GetProductByEan(receivedProduct.ean).ProductID == 0){
		database.ReceiveNewProduct(mapBItemToDBItem(receivedProduct))
	}

	//if edited Product
	if(database.GetProductByEan(receivedProduct.ean).ProductID != 0){
		database.ReceiveEditProduct(mapBItemToDBItem(receivedProduct))
	}

	//if deleted Product
	if(database.GetProductByEan(receivedProduct.ean).EAN != 0){
		database.ReceiveRemoveProduct(receivedProduct.ProductID)
	}
}
*/

/*
**
Send Functions
**
*/
func SendRemovedProduct(id int) {

	sendProduct := model.BProduct{ProductID: id}

	//serialize product
	serializedSendProduct, err := json.Marshal(sendProduct)
	if err != nil {
		panic(err)
	}

	//send products to all interested channels
	reportingcommunication.ReportingListenerProducts <- serializedSendProduct
	storecommunication.StoreListenerProducts <- serializedSendProduct
	warehousecommunication.WarehouseListenerProducts <- serializedSendProduct

}

func SendEditProduct(sendProduct model.BProduct) {

	//serialize product
	serializedSendProduct, err := json.Marshal(sendProduct)
	if err != nil {
		panic(err)
	}

	//send products to all interested channels
	reportingcommunication.ReportingListenerProducts <- serializedSendProduct
	storecommunication.StoreListenerProducts <- serializedSendProduct
	warehousecommunication.WarehouseListenerProducts <- serializedSendProduct

}

/*
**
Mapping B Model -> DB Model
**
*/
