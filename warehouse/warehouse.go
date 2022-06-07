package warehouse

import (
	"example/Masterarbeit/basedata"
)

func StockProduct(ean string) bool {

	//Proof if ean is in basedata
	if basedata.GetProductByEAN(ean).EAN != ean {
		//Failure, because product isnt available
		return false
	} else {
		//Send information to reporting ToDo
		return true
	}

}
