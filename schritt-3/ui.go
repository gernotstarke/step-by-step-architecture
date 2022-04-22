package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type protoui struct {
	app           fyne.App
	win           fyne.Window
	file          *widget.Label
	fileInput     *widget.Entry
	result        *canvas.Text
	srcFileButton *widget.Button
	count         *widget.Button
	cancel        *widget.Button
}

func buttonGroup(p *protoui) fyne.CanvasObject {

	p.count = widget.NewButton("Count Pages", func() {

		p.result.Refresh()

		pc := PageCount(p.fileInput.Text)
		p.result.Text = p.fileInput.Text + fmt.Sprintf(" has %d pages", pc)
		//p.result.SetText(fmt.Sprint(pc))
	})

	p.cancel = widget.NewButton("Cancel", func() {
		p.app.Quit()
	})

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		p.cancel,
		layout.NewSpacer(),
		p.count)

	return buttons

}

func srcFileGroup(p *protoui) *fyne.Container {
	p.fileInput = widget.NewEntry()
	p.fileInput.SetText("./pdfs/1pg.pdf")

	/*
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

		})*/

	p.result = canvas.NewText("nothing selected", color.Black)

	// TODO: smaller size for result-text
	//p.result.TextStyle = 9

	return container.New(layout.NewVBoxLayout(), p.fileInput, p.result)

}

func createUI() {

	p := &protoui{}
	// application and visible window
	p.app = app.New()
	p.win = p.app.NewWindow("Step-3")

	createAndDisplaySplash()

	p.win.SetContent(container.NewVBox(srcFileGroup(p),
		container.New(layout.NewHBoxLayout(), buttonGroup(p))))

	p.win.ShowAndRun()
}

func createAndDisplaySplash() {
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		splashWindow := drv.CreateSplashWindow()

		splashWindow.SetContent(
			fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(500, 500)),
				canvas.NewImageFromFile("../assets/logo.png")))

		splashWindow.Show()

		go func() {
			time.Sleep(time.Second * 5)
			splashWindow.Close()
		}()
	}
}
