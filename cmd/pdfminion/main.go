package main

import (
	"schritt-4/internal/domain"
	"schritt-4/internal/gui"
)

func main() {

	domain.SetupConfiguration()

	// create and run the user interface
	gui.CreateMainUI()

}
