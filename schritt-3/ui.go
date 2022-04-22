package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type protoui struct {
	app           fyne.App
	win           fyne.Window
	file          *widget.Label
	fileInput     *widget.Entry
	result        *widget.Label
	srcFileButton *widget.Button
	count         *widget.Button
	cancel        *widget.Button
}

func buttonGroup(p *protoui) fyne.CanvasObject {

	p.count = widget.NewButton("Count Pages", func() {
		// call domain logic here!

		// set label appropriately
	})

	p.cancel = widget.NewButton("Cancel", func() {
		p.app.Quit()
	})

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		p.cancel,
		p.count)

	okCancelPanel := widget.NewCard("", "Processing", buttons)

	return okCancelPanel
}

func srcFileGroup(p *protoui) *fyne.Container {
	srcDirField := widget.NewEntry()
	srcDirField.SetText("Step-3")

	p.srcFileButton = widget.NewButton("File", func() {
		dialog.ShowFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil { // there was an error - tell user
					dialog.ShowError(err, p.win)
					return
				}
				if reader == nil { // user cancelled
					return
				}
				// reader contains the file
				srcDirField.SetText(reader.URI().Name())

			}, p.win)

	})

	srcDirStatusLabel := canvas.NewText("nothing selected", color.Gray{})
	srcDirStatusLabel.TextSize = 9

	inputStatus := container.New(layout.NewVBoxLayout(), srcDirField, srcDirStatusLabel)

	return container.New(layout.NewFormLayout(), p.srcFileButton, inputStatus)
}

func createUI() {

	p := &protoui{}
	// application and visible window
	p.app = app.New()
	p.win = p.app.NewWindow("Step-3")

	// upon button press, set status widget
	p.count = widget.NewButton("Count pages", func() {
		p.result.SetText(string(PageCount(p.fileInput.Text)))
	})

	p.win.SetContent(container.NewVBox(srcFileGroup(p),
		container.New(layout.NewHBoxLayout(), buttonGroup(p))))

	p.win.ShowAndRun()
}
