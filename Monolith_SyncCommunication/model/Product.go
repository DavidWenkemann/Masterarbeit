//contains product models for all three layers

package model

type DBProduct struct {
	ProductID int     `sql:"id"`
	EAN       string  `sql:"ean"`
	Name      string  `sql:"name"`
	Price     float64 `sql:"price"`
}

type BProduct struct {
	ProductID int
	EAN       string
	Name      string
	Price     float64
}

type APIProduct struct {
	ProductID int     `json:"id"`
	EAN       string  `json:"ean"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}
