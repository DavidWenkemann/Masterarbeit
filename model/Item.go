package model

import "time"

type APIItem struct {
	Product       BProduct
	ItemID        string
	ReceivingDate time.Time
	SellingDate   time.Time
}

type BItem struct {
	Product       BProduct
	ItemID        string
	ReceivingDate time.Time
	SellingDate   time.Time
}

type DBItem struct {
	ProductID     int
	ItemID        string
	ReceivingDate time.Time
	SellingDate   time.Time
}
