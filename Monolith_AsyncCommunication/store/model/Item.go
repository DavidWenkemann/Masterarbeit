//contains item models for all three layers
package model

import "time"

type APIItem struct {
	Product APIProduct
	ItemID  string
}

type BItem struct {
	ProductID     int
	ItemID        string
	ReceivingDate time.Time
	SellingDate   time.Time
}

type DBItem struct {
	ProductID     int
	ItemID        string
	ReceivingDate time.Time
}
