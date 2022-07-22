package main

import (
	"github.com/DavidWenkemann/Masterarbeit/database"
	userinterface "github.com/DavidWenkemann/Masterarbeit/ui"
)

func main() {

	database.SpinupDB()
	userinterface.StartUI()
}
