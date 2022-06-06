package basedata

type product struct {
	EAN      string `json: "ean"`
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}

var products = []product{
	{EAN: "4011803092174", Name: "Spezi", Quantity: 2},
	{EAN: "4066600641919", Name: "Paulaner Weissbier alk.frei", Quantity: 5},
	{EAN: "4029764001807", Name: "Clubmate", Quantity: 6},
	{EAN: "123", Name: "Bier", Quantity: 10},
}

//searches and returnes specific product by ean
func GetProductByEAN(key string) product {

	var p product

	//searches product with same ean in slice and returns it
	for i := range products {

		if products[i].EAN == key {
			return products[i]
		}
	}

	//returns an empty product if ean wasnt found
	return p

}

//returns all products in the slice
func GetAllProducts() []product {
	return products
}

//Adds a new product to the slice, quanity is 0
func AddProduct(ean string, name string) product {

	//TODO: wenn ean schon vorhanden -> Fehler

	//Initialize new Product
	p := product{
		EAN:  ean,
		Name: name,
	}

	//Add new Product to the slice
	products = append(products, p)

	//return the new product
	return p
}

//removes specific product, changes sorting of slice
func RemoveProductByEAN(key string) {

	//searchs product with same ean in slice
	for i := range products {
		if products[i].EAN == key {

			//replacing found product with last product in slice
			products[i] = products[len(products)-1]
			//shorten up slice by 1
			products = products[:len(products)-1]
			return
		}
	}
	return
}
