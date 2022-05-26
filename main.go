// prepare for automated testing
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type protoui struct {
	app    fyne.App
	win    fyne.Window
	status *widget.Label
	input  *widget.Entry
	accept *widget.Button
	cancel *widget.Button
}

func makeUI(p *protoui) fyne.CanvasObject {

	// application and visible window
	p.app = app.New()
	p.win = p.app.NewWindow("Take-2")

	// widgets
	p.status = widget.NewLabel("")
	p.input = widget.NewEntry()
	p.input.PlaceHolder = "enter something"

	// upon button press, set status widget
	p.accept = widget.NewButton("Action", func() {
		p.status.SetText(p.input.Text)
	})

	// cancel button quits application
	p.cancel = widget.NewButton("Cancel", func() {
		p.app.Quit()
	})

	return container.NewVBox(p.status, p.input,
		container.New(layout.NewHBoxLayout(), p.cancel, p.accept))
}

func main() {

	// initialize prototype UI
	p := &protoui{}
	content := makeUI(p)

	p.win.SetContent(content)
	p.win.ShowAndRun()
}
