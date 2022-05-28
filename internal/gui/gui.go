package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	domain "schritt-4/internal/domain"
	"schritt-4/resources"
)

// Appl exposes the fyne application - mainly to enable the quit-function to stop the app.
var Appl fyne.App
var appVersion string

var UIState uiState

// Window is the main application window
var Window fyne.Window

// CreateMainUI creates and shows the main graphical user interface, second version.
// It creates by delegating to "Panel" functions which will create their respective panel.
func CreateMainUI() {

	Appl = app.New()

	Window = Appl.NewWindow(domain.AppName)

	content := container.New(layout.NewVBoxLayout(),
		logoHeaderPanel(),
		directoriesPanel(),
		widget.NewSeparator(),
		okCancelPanel())

	Window.SetContent(content)
	Window.Resize(fyne.NewSize(600, 400))
	Window.SetFixedSize(true)
	Window.CenterOnScreen()
	Window.ShowAndRun()
}

func logoHeaderPanel() *fyne.Container {
	arc42Logo := canvas.NewImageFromResource(resources.Arc42LogoPNG)
	arc42Logo.FillMode = canvas.ImageFillContain
	arc42Logo.SetMinSize(fyne.NewSize(80, 40))
	arc42Logo.Resize(fyne.NewSize(80, 40))

	/*
		arc42URL, _ := url.Parse("https://arc42.org")
		arc42Link := widget.NewHyperlinkWithStyle("arc42.org",
			arc42URL, fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	*/
	appLogo := canvas.NewImageFromResource(resources.PDFminionLogoPNG)
	appLogo.FillMode = canvas.ImageFillContain
	appLogo.SetMinSize(fyne.NewSize(200, 120))

	appLogo.Resize(fyne.NewSize(200, 120))

	return container.New(layout.NewHBoxLayout(),
		container.New(layout.NewVBoxLayout(),
			arc42Logo,
			layout.NewSpacer(),
			//arc42Link),
		),
		layout.NewSpacer(),
		appLogo,
	)
}

func directoriesPanel() fyne.CanvasObject {

	dirContainer := container.New(layout.NewVBoxLayout(),
		srcDirSelectorGroup())
	dirPanel := widget.NewCard("", "Directory", dirContainer)

	return dirPanel
}

func srcDirSelectorGroup() *fyne.Container {
	srcDirField := widget.NewEntry()
	srcDirField.SetText(domain.SourceDirName())

	srcDirButton := widget.NewButton("Source", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			if list == nil {
				return
			}

			_, err = list.List()
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			srcDirField.SetText(list.Path())
			domain.SetSourceDirName(list.Path())

			if domain.CheckSrcDirectoryStatus(list.Path()) {
				fmt.Printf("Folder %s :\n%s", list.Path(), list.String())
				// dialog.ShowInformation("Folder Open", out, Window)
			}
		}, Window)
	})
	srcDirButton.SetIcon(theme.FolderOpenIcon())

	domain.CheckSrcDirectoryStatus(domain.SourceDirName())
	srcDirStatusLabel := canvas.NewText("nothing selected", color.Gray{128})
	srcDirStatusLabel.TextSize = 9

	inputStatus := container.New(layout.NewVBoxLayout(), srcDirField, srcDirStatusLabel)

	return container.New(layout.NewFormLayout(), srcDirButton, inputStatus)

}

func okCancelPanel() fyne.CanvasObject {

	OKButton := widget.NewButton("Process PDFs", func() {})
	OKButton.Disable()

	CancelButton := widget.NewButton("Cancel", quitApp)

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		CancelButton,
		OKButton)

	okCancelPanel := widget.NewCard("", "Processing", buttons)

	return okCancelPanel
}

func statusLine(msg string) *fyne.Container {

	statusMsg := canvas.NewText(msg, DarkRedColor)
	statusMsg.TextSize = 9
	statusMsg.Alignment = fyne.TextAlignTrailing

	versionLabel := canvas.NewText("v."+domain.VersionStr, NavyColor)
	versionLabel.TextSize = 9
	versionLabel.Alignment = fyne.TextAlignCenter

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		versionLabel,
		widget.NewSeparator(),
		layout.NewSpacer(),
		statusMsg,
	)
}

func quitApp() {
	Appl.Quit()
}

func setUIState() uiState {
	uis := UIState

	uis.sourceDirName = ""
	return uis
}
