package mainMenu

import (
	"ProjectMarekEmanuel/mainMenu/gamesWindows"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CreditsScreen struct {
	window fyne.Window
	app    fyne.App
}

func NewCreditsScreen(window fyne.Window, app fyne.App) *CreditsScreen {
	return &CreditsScreen{window: window, app: app}
}

func (cs *CreditsScreen) Render() {
	backButton := widget.NewButton("Back", func() {
		cs.window.SetContent(MainMenuContent)
	})
	optionsButton := widget.NewButtonWithIcon("", theme.HelpIcon(), func() {
		gamesWindows.MakeOptionWindow(cs.app)
	})

	// Add other elements of the Play screen here

	cs.window.SetContent(container.NewVBox(
		optionsButton,
		backButton,
		// other widgets
	))
}
