// COMMUNICATION
//
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
