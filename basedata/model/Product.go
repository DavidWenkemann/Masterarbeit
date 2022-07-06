package model

type DBProduct struct {
	EAN   string  `sql:"ean"`
	Name  string  `sql:"name"`
	Price float64 `sql:"price"`
}

type BProduct struct {
	EAN   string
	Name  string
	Price float64
}

type APIProduct struct {
	EAN   string  `json:"ean"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
