package main

import (
	"fmt"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/database"
	userinterface "github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/ui"
)

func main() {

	//Clears Terminal
	fmt.Printf("\x1bc")

	//Loads Data
	database.SpinupDB()

	//Starts UI
	userinterface.StartUI()
}
