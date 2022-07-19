//models for all three layers for product

package model

type DBProduct struct {
	ProductID int     `sql:"id"`
	EAN       string  `sql:"ean"`
	Name      string  `sql:"name"`
	Price     float64 `sql:"price"`
	//Quantity int     `sql:"quantity"`
}

type BProduct struct {
	EAN   string
	Name  string
	Price float64
	//Quantity int
}

type APIProduct struct {
	EAN   string  `json:"ean"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	//Quantity int     `json:"quantity"`
}
