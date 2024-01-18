package gamesWindows

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func MakeLoseWindow(mainApp fyne.App, wind fyne.Window, mainContent fyne.CanvasObject, game string) {
	looseImgRes := fyne.NewStaticResource("looseIcon.png", LooseLevelImg)
	looseIcon := canvas.NewImageFromResource(looseImgRes)
	WinWindowImgResource := fyne.NewStaticResource("WinLevelImg.png", winLevelImg)
	windWindowWidget := canvas.NewImageFromResource(WinWindowImgResource)
	windWindowWidget.Translucency = 0.25
	btn := widget.Button{Text: "Restart"}
	btn2 := widget.Button{Text: "main menu"}
	winLabel := canvas.Text{Text: "Game Over"}
	winLabel.Alignment = fyne.TextAlignCenter
	winLabel.TextSize = 24
	winLabel.TextStyle.Bold = true
	winLabel.Color = color.RGBA{
		R: 128,
		G: 128,
		B: 128,
		A: 255,
	}

	imgContainer := container.NewCenter(
		layout.NewSpacer(),
		looseIcon,
	)

	btnConrainer := container.NewHBox(
		layout.NewSpacer(),
		&btn,
		&btn2,
		layout.NewSpacer(),
	)
	labelContainer := container.NewVBox(
		&winLabel,
	)

	// Main ContentPaint
	mainContentContainer := container.NewVBox(
		labelContainer,
		layout.NewSpacer(),
		imgContainer,
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
			Game2048WindowInProggress.window.SetContent(container.NewVBox(setUpLayout(Game2048WindowInProggress)))

			// Update the keyboard listener to use the new game state
			setUpKeyboardListener(Game2048WindowInProggress.window, Game2048WindowInProggress, newState)
		}
		win.Close()
	}
	combinedContainer := container.NewStack(
		windWindowWidget,
		mainContentContainer,
	)
	win.SetContent(combinedContainer)
	win.CenterOnScreen()

	win.Resize(fyne.NewSize(400, 220))
	win.Show()
}
