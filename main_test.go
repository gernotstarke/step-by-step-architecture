package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestEntry(t *testing.T) {

	p := &protoui{}
	makeUI(p)

	if p.status.Text != "" {
		t.Errorf("Initial value should be empty, but was %s", p.status.Text)
	}

	test.Type(p.input, "42")
	test.Tap(p.accept)
	if p.status.Text != "42" {
		t.Errorf("Value after 'Action' should be '42' but was %s", p.status.Text)
	}
}
