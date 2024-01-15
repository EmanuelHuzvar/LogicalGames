package gamesWindows

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var lowImportance = widget.HighImportance
var mediumImportance = widget.WarningImportance
var highImportance = widget.DangerImportance

type LevelScreen struct {
	window          fyne.Window
	app             fyne.App
	game            string
	mainMenuContent fyne.CanvasObject
}

func NewLevelScreen(window fyne.Window, app fyne.App, mainMenuContent fyne.CanvasObject, game string) *LevelScreen {
	return &LevelScreen{window: window, app: app, mainMenuContent: mainMenuContent, game: game}
}

var CurrentLevel string

var LevelContent fyne.CanvasObject

func (ls *LevelScreen) Render() {
	app := app.New()
	myWindow := app.NewWindow("Logitec App")
	myWindow.SetOnClosed(func() {
		app.Quit()
	})

	grid := container.NewGridWithColumns(5) // Create a grid with 5 columns

	// Function to create a new button that prints its label when clicked
	makeButton := func(label string, importance widget.Importance) *widget.Button {
		btn := widget.NewButton(label, func() {
			fmt.Println("Button clicked:", label)
			CurrentLevel = label
			if ls.game == "bubble" {
				NewBubbleScreen(ls.window, ls.app, ls.mainMenuContent, label).Render()
			}
			if ls.game == "nonogram" {
				NewNonogramScreen(ls.window, ls.app, ls.mainMenuContent, label).Render()
			}
			if ls.game == "paint" {
				NewPaintFloorScreen(ls.window, ls.app, ls.mainMenuContent, label).Render()
			}
			if ls.game == "2048" {
				NewNonogramScreen(ls.window, ls.app, ls.mainMenuContent, label).Render()
			}

		})
		btn.Importance = importance
		return btn
	}

	// Add buttons with numbers to the grid
	for i := 1; i <= 15; i++ {
		if i > 0 && i < 6 {
			grid.Add(makeButton(fmt.Sprintf("%d", i), lowImportance))
		}
		if i > 5 && i < 11 {
			grid.Add(makeButton(fmt.Sprintf("%d", i), mediumImportance))
		}
		if i > 10 {
			grid.Add(makeButton(fmt.Sprintf("%d", i), highImportance))
		}

	}

	LevelContent = grid
	ls.window.SetContent(grid)
	ls.window.CenterOnScreen()
}
