package mainMenu

import (
	"ProjectMarekEmanuel/mainMenu/gamesWindows"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	//optionsButton := widget.NewButtonWithIcon("", theme.HelpIcon(), func() {
	//	gamesWindows.MakeOptionWindow(ps.app)
	//})
	gameButtons := []fyne.CanvasObject{
		widget.NewButton("Bubble Sort", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "bubble", MainMenuContent).Render()
			//gamesWindows.NewBubbleScreen(ps.window, ps.app, GameWindowContainer).Render()
		}),
		widget.NewButton("Nonogram", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "nonogram", MainMenuContent).Render()
		}),
		widget.NewButton("Paint floor", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "paint", MainMenuContent).Render()
		}),
		widget.NewButton("2048", func() {
			gamesWindows.NewGame2048Screen(ps.window, ps.app, GameWindowContainer, MainMenuContent).Render()
		}),
	}
	gameGrid := container.NewAdaptiveGrid(2, gameButtons...)
	topLeftContainer := container.NewVBox(
		backButton,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	finalContainer := container.NewHBox(

		topLeftContainer,
		layout.NewSpacer(),
	)

	gameGridContainer := container.NewBorder(nil, nil, nil, nil, gameGrid)
	gameGridContainer.Add(finalContainer)

	// Add other elements of the Play screen here

	ps.window.SetContent(container.NewVBox(
		finalContainer,
		backButton,
		// other widgets
	))
	GameWindowContainer = gameGridContainer
	ps.window.SetContent(gameGridContainer)
}
