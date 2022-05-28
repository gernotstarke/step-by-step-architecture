package main

import (
	"testing"
)

// TestPageCount is a trivial test for a single-page pdf.
// Additional testcases would be useful.
func TestPageCount(t *testing.T) {
	var FNAME = "pdfs/1pg.pdf"
	pc := PageCount(FNAME)
	if pc != 1 {
		t.Errorf("PageCount of %s = %d; but should be %d", FNAME, pc, 1)
	}
}
