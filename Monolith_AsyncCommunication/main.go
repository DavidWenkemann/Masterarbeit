package main

import (
	"fmt"

	userinterface "github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/ui"
)

func main() {

	//Clears Terminal
	fmt.Printf("\x1bc")

	//Loads Data
	SpinupDB()

	//Starts UI
	userinterface.StartUI()
}

func SpinupDB() {

}

func AddProduct() {

}

func AddItem() {

}
