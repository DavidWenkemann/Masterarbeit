package basedata

type product struct {
	EAN      string `json: "ean"`
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}

var products = []product{
	{EAN: "4011803092174", Name: "Spezi", Quantity: 2},
	{EAN: "4066600641919", Name: "Paulaner Hefe Weißbier alkoholfrei", Quantity: 5},
	{EAN: "4029764001807", Name: "Clubmate", Quantity: 6},
}

func GetProductsByEAN(key string) product {

	var p product

	for i := range products {

		if products[i].EAN == key {
			return products[i]
		}
	}

	return p

}
