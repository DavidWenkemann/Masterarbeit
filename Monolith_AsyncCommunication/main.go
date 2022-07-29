package main

import (
	"fmt"

	reportingcommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/reporting/communication"
	storecommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/communication"
	warehousecommunication "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/warehouse/communication"

	userinterface "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/ui"
)

func main() {

	//Clears Terminal
	fmt.Printf("\x1bc")

	//Loads Data
	SpinupDB()

	reportingcommunication.ReportingListener()
	storecommunication.StoreListener()
	warehousecommunication.WarehouseListener()

	//Starts UI
	userinterface.StartUI()
}

func SpinupDB() {

}

func AddProduct() {

}

func AddItem() {

}
