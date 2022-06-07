package warehouse

import (
	"example/Masterarbeit/basedata"
)

/*
type product struct {
	EAN      string `json:"ean"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var products = []product{
	{EAN: "4011803092174", Name: "Spezi", Quantity: 2},
	{EAN: "4066600641919", Name: "Paulaner Weissbier alk.frei", Quantity: 2},
	{EAN: "4029764001807", Name: "Clubmate", Quantity: 2},
}
*/

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
