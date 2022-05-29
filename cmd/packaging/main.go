package main

import (
	"fmt"
	version "pdfminion"
)

// demonstrate the use of the cmd-internal packing approach, see ADR-0003

func main() {
	fmt.Println("Checking packaging.")
	fmt.Println("Version: ", version.Version())
	fmt.Println("Status: ", version.Status())
}
