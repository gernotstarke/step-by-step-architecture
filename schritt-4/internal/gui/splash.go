package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"schritt-4/resources"
	"time"
)

func createAndDisplaySplash() {
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		splashWindow := drv.CreateSplashWindow()

		splashWindow.SetContent(
			container.New(layout.NewGridWrapLayout(fyne.NewSize(400, 400)),
				canvas.NewImageFromResource(resources.PDFminionLogoPNG)))

		splashWindow.Show()

		go func() {
			time.Sleep(time.Second * 5)
			splashWindow.Close()
		}()
	}
}
