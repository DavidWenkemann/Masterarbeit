package reporting

type Product struct {
	EAN      string `json:"ean"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var productsInStock = []Product{
	{EAN: "4011803092174", Name: "Spezi", Quantity: 2},
	{EAN: "4066600641919", Name: "Paulaner Weissbier alk.frei", Quantity: 2},
	{EAN: "4029764001807", Name: "Clubmate", Quantity: 2},
}

func StockProductByOne(ean string) bool {

	for i := range productsInStock {

		if productsInStock[i].EAN == ean {
			productsInStock[i].Quantity += 1
			return true
		}
	}
	return false
}

func RemoveProductByOne(ean string) bool {

	for i := range productsInStock {

		if productsInStock[i].EAN == ean {
			productsInStock[i].Quantity -= 1
			return true
		}
	}
	return false
}
