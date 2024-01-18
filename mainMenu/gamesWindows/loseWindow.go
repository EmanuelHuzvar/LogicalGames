package gamesWindows

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MakeLoseWindow(mainApp fyne.App, wind fyne.Window, mainContent fyne.CanvasObject, game string) {

	btn := widget.Button{Text: "Restart"}
	btn2 := widget.Button{Text: "main menu"}
	winLabel := canvas.Text{Text: "Game Over"}
	winLabel.TextSize = 24
	winLabel.TextStyle.Bold = true

	btnConrainer := container.NewHBox(
		layout.NewSpacer(),
		&btn,
		&btn2,
		layout.NewSpacer(),
	)
	labelContainer := container.NewCenter(
		&winLabel,
	)

	// Main ContentPaint
	mainContentContainer := container.NewVBox(
		labelContainer,
		btnConrainer,
	)

	// Window Settings
	win := mainApp.NewWindow("Win")

	btn2.OnTapped = func() {
		btnSetContent(wind, mainContent)
		win.Close()
	}

	btn.OnTapped = func() {
		if game == "2048" {
			// Reset the game state and re-render the grid
			newState := NewGameState(4)
			addRandomTile(newState)
			gameStateInProgress = newState // Update the global game state
			renderGrid(newState, Game2048WindowInProggress)
			Game2048WindowInProggress.window.SetContent(container.NewVBox(Game2048WindowInProggress.scoreLabel, Game2048WindowInProggress.gridLayout))

			// Update the keyboard listener to use the new game state
			setUpKeyboardListener(Game2048WindowInProggress.window, Game2048WindowInProggress, newState)
		}
		win.Close()
	}
	win.SetContent(mainContentContainer)
	win.CenterOnScreen()

	win.Resize(fyne.NewSize(400, 220))
	win.Show()
}
