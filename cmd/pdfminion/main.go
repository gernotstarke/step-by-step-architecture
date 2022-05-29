package main

import (
	version "pdfminion"
	"pdfminion/internal/domain"
	"pdfminion/internal/gui"
)

func main() {

	domain.SetupConfiguration()

	// create and run the user interface
	gui.CreateMainUI(version.VersionStr + version.Status())

}
