package gamesWindows

import (
	"ProjectMarekEmanuel/mainMenu/gamesWindows/customButtons"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type LevelScreen struct {
	window          fyne.Window
	app             fyne.App
	game            string
	mainMenuContent fyne.CanvasObject
	mainMenuWindow  fyne.CanvasObject
}

var RedColorWin = color.RGBA{
	R: 255,
	G: 179,
	B: 186,
	A: 255,
}
var GreenColorWin = color.RGBA{
	R: 186,
	G: 255,
	B: 201,
	A: 255,
}
var YellowColorWin = color.RGBA{
	R: 255,
	G: 255,
	B: 186,
	A: 255,
}

func NewLevelScreen(window fyne.Window, app fyne.App, mainMenuContent fyne.CanvasObject, game string, mainMenuWindow fyne.CanvasObject) *LevelScreen {
	return &LevelScreen{window: window, app: app, mainMenuContent: mainMenuContent, game: game, mainMenuWindow: mainMenuWindow}
}

var CurrentLevel string

var LevelContent fyne.CanvasObject

func (ls *LevelScreen) Render() {

	BackBtnImgResource := fyne.NewStaticResource("back-arrow.png", backBtnImg)

	app := app.New()
	myWindow := app.NewWindow("Logitec App")
	myWindow.SetOnClosed(func() {
		app.Quit()
	})

	grid := container.NewGridWithColumns(5) // Create a grid with 5 columns

	// Function to create a new button that prints its label when clicked
	makeButton := func(label string, colorOfBtn color.Color) *customButtons.ColorButton {
		btn := customButtons.NewColorButton(colorOfBtn, label, func() {
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

		return btn
	}

	// Add buttons with numbers to the grid
	for i := 1; i <= 15; i++ {
		if i > 0 && i < 6 {
			grid.Add(makeButton(fmt.Sprintf("%d", i), GreenColorWin))
		}
		if i > 5 && i < 11 {
			grid.Add(makeButton(fmt.Sprintf("%d", i), YellowColorWin))
		}
		if i > 10 {
			grid.Add(makeButton(fmt.Sprintf("%d", i), RedColorWin))
		}

	}
	backButton := widget.NewButtonWithIcon("", BackBtnImgResource, func() {
		ls.window.SetContent(ls.mainMenuContent)
	})

	topLeftContainer := container.NewVBox(

		backButton,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	finalContainer := container.NewHBox(

		topLeftContainer,
		layout.NewSpacer(),
	)

	gameGridContainer := container.NewBorder(nil, nil, nil, nil, grid)
	gameGridContainer.Add(finalContainer)
	LevelContent = gameGridContainer

	ls.window.SetContent(gameGridContainer)
	ls.window.CenterOnScreen()
}
