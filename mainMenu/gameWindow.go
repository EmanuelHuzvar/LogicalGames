package mainMenu

import (
	"ProjectMarekEmanuel/mainMenu/gamesWindows"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type PlayScreen struct {
	window fyne.Window
	app    fyne.App
}

func NewPlayScreen(window fyne.Window, app fyne.App) *PlayScreen {
	return &PlayScreen{window: window, app: app}
}

var GameWindowContainer fyne.CanvasObject

func (ps *PlayScreen) Render() {
	backButton := widget.NewButton("Back", func() {
		ps.window.SetContent(MainMenuContent)
	})
	optionsButton := widget.NewButtonWithIcon("", theme.HelpIcon(), func() {
		gamesWindows.MakeOptionWindow(ps.app)
	})
	gameButtons := []fyne.CanvasObject{
		widget.NewButton("Bubble Sort", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "bubble").Render()
			//gamesWindows.NewBubbleScreen(ps.window, ps.app, GameWindowContainer).Render()
		}),
		widget.NewButton("Nonogram", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "nonogram").Render()
		}),
		widget.NewButton("Roll Ball", func() { /* Your click handler code */ }),
		widget.NewButton("Unblock", func() { /* Your click handler code */ }),
	}
	gameGrid := container.NewAdaptiveGrid(2, gameButtons...)
	topLeftContainer := container.NewVBox(
		optionsButton,
		layout.NewSpacer(),
	)
	finalContainer := container.NewHBox(
		topLeftContainer,
		layout.NewSpacer(),
	)
	gameGridContainer := container.NewBorder(nil, backButton, nil, nil, gameGrid)

	// Add other elements of the Play screen here

	ps.window.SetContent(container.NewVBox(
		finalContainer,
		backButton,
		// other widgets
	))
	GameWindowContainer = gameGridContainer
	ps.window.SetContent(gameGridContainer)
}
