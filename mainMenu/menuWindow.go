package mainMenu

import (
	"ProjectMarekEmanuel/emanuel/buttons"
	"ProjectMarekEmanuel/mainMenu/gamesWindows"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
)

//go:embed pictures/background.jpg
var backgroundImg []byte

var MainMenuContent fyne.CanvasObject

func MakeMenu() fyne.Window {
	imgResource := fyne.NewStaticResource("background.jpg", backgroundImg)
	backgroundImage := canvas.NewImageFromResource(imgResource)
	backgroundImage.FillMode = canvas.ImageFillStretch
	app := app.New()
	menuWindow := app.NewWindow("Logitec App")
	menuWindow.SetOnClosed(func() {
		app.Quit()
	})

	exitButton := buttons.NewColoredTextButton("Quit", colornames.Blue, colornames.Azure, func() {
		menuWindow.Close()
	})

	playButton := buttons.NewColoredTextButton("Play", colornames.Blue, colornames.Azure, func() {
		playScreen := NewPlayScreen(menuWindow, app)
		playScreen.Render()
	})
	creditsButton := buttons.NewColoredTextButton("Credits", colornames.Blue, colornames.Azure, func() {
		creditsScreen := NewCreditsScreen(menuWindow, app)
		creditsScreen.Render()
	})
	questionMarkButton := widget.NewButtonWithIcon("", theme.HelpIcon(), func() {
		gamesWindows.MakeOptionWindow(app)
	})
	buttonContainer := container.NewVBox(
		playButton,
		creditsButton,
		exitButton,
	)
	topLeftContainer := container.NewVBox(
		questionMarkButton,
		layout.NewSpacer(),
	)
	finalContainer := container.NewHBox(
		topLeftContainer,
		layout.NewSpacer(),
	)
	centerContainer := container.New(
		layout.NewCenterLayout(),
		buttonContainer,
	)
	contentContainer := container.NewStack(
		backgroundImage,
		centerContainer,
		finalContainer,
	)

	MainMenuContent = contentContainer

	menuWindow.SetContent(contentContainer)
	menuWindow.Resize(fyne.NewSize(700, 700))
	menuWindow.CenterOnScreen()
	return menuWindow

}
