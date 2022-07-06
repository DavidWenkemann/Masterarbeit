package database

//pruducts emulates a Database table with 3 columns EAN (unique) ,NAME and Price
var products = []Product{
	{EAN: "4011803092174", Name: "Spezi", Price: 0.75},
	{EAN: "4066600641919", Name: "Paulaner Weissbier alk.frei", Price: 1.39},
	{EAN: "4029764001807", Name: "Clubmate", Price: 0.95},
}

type Product struct {
	EAN   string  `json:"ean"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
