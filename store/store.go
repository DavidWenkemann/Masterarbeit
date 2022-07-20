package store

import "github.com/DavidWenkemann/Masterarbeit/database"

//Sells Item and Sets SellingDate on time now
func SellItem(itemID string) {
	database.SetItemSelledDate(itemID)
}
