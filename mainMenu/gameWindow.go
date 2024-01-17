package mainMenu

import (
	"ProjectMarekEmanuel/mainMenu/gamesWindows"
	"ProjectMarekEmanuel/mainMenu/gamesWindows/customButtons"
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

//go:embed pictures/back-arrow.png
var BackBtnImg []byte

//go:embed pictures/bubbleSortGame.png
var BubbleSortGameImg []byte

//go:embed pictures/nonogramGame.png
var NonogramGameImg []byte

//go:embed pictures/paintFloorGame.png
var PaintFloorGameImg []byte

//go:embed pictures/game2048Game.png
var game2048Img []byte

func (ps *PlayScreen) Render() {
	backBtnImgResource := fyne.NewStaticResource("back-arrow.png", BackBtnImg)
	ImgResourceBubble := fyne.NewStaticResource("bubbleSortGame.png", BubbleSortGameImg)
	ImgResourceNonogram := fyne.NewStaticResource("nonogramGame.png", NonogramGameImg)
	ImgResourcePaintFloor := fyne.NewStaticResource("paintFloorGame.png", PaintFloorGameImg)
	ImgResource2048 := fyne.NewStaticResource("game2048Game.png", game2048Img)
	backButton := widget.NewButtonWithIcon("", backBtnImgResource, func() {
		ps.window.SetContent(MainMenuContent)
	})

	//optionsButton := widget.NewButtonWithIcon("", theme.HelpIcon(), func() {
	//	gamesWindows.MakeOptionWindow(ps.app)
	//})
	gameButtons := []fyne.CanvasObject{

		customButtons.NewImageButton(ImgResourceBubble, "Bubble Sort", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "bubble", MainMenuContent).Render()
			//gamesWindows.NewBubbleScreen(ps.window, ps.app, GameWindowContainer).Render()
		}),
		customButtons.NewImageButton(ImgResourceNonogram, "Nonogram", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "nonogram", MainMenuContent).Render()
		}),
		customButtons.NewImageButton(ImgResourcePaintFloor, "Paint floor", func() {
			gamesWindows.NewLevelScreen(ps.window, ps.app, GameWindowContainer, "paint", MainMenuContent).Render()
		}),
		customButtons.NewImageButton(ImgResource2048, "2048", func() {
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
