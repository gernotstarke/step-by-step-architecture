package main

import (
	"internal/domain"
	"internal/gui"
)

func main() {

	domain.SetupConfiguration()

	// create and run the user interface
	gui.CreateMainUI()

}
