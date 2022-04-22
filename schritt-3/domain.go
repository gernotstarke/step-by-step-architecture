package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// PageCount expects fileName to exist
func PageCount(fileName string) int {
	// TODO: handle some errors, e.g. file non pdf, zero length file, nonexistent file
	nrOfPages, err := api.PageCountFile(fileName)
	if err != nil {
		fmt.Println("Cannot determine pageCount:", err)
	}
	return nrOfPages
}
