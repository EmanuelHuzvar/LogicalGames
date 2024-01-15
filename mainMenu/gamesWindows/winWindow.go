package gamesWindows

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

//go:embed pictures/check.png
var checkImg []byte

func MakeWinWindow(mainApp fyne.App, wind fyne.Window, mainContent fyne.CanvasObject, game string) {

	checkImgResource := fyne.NewStaticResource("check.png", checkImg)
	checkImgWidget := canvas.NewImageFromResource(checkImgResource)

	checkImgWidget.SetMinSize(fyne.NewSize(128, 128))
	btn := widget.Button{Text: "next level"}
	btn2 := widget.Button{Text: "main menu"}
	btn3 := widget.Button{Text: "levels"}
	winLabel := canvas.Text{Text: "You successfully solved the level"}
	winLabel.TextSize = 24
	winLabel.TextStyle.Bold = true

	imgContainer := container.NewCenter(
		layout.NewSpacer(),
		checkImgWidget,
	)
	btnConrainer := container.NewHBox(
		layout.NewSpacer(),
		&btn,
		&btn3,
		&btn2,
		layout.NewSpacer(),
	)
	labelContainer := container.NewCenter(
		&winLabel,
	)

	// Main ContentPaint
	mainContentContainer := container.NewVBox(
		labelContainer,
		imgContainer,
		btnConrainer,
	)

	// Window Settings
	win := mainApp.NewWindow("Win")

	btn2.OnTapped = func() {
		btnSetContent(wind, mainContent)
		win.Close()
	}
	btn3.OnTapped = func() {
		btnSetContent(wind, LevelContent)
		win.Close()
	}
	btn.OnTapped = func() {
		if game == "bubble" {
			content := makeLevel(increamentLevel(LevelInProggress))
			btnSetContent(wind, content)
		}
		if game == "paint" {
			content := MakeGamePaintFloor(increamentLevel(LevelInProggressPaint))
			containeris := content.Content()
			btnSetContent(wind, containeris)
			SetUpPaintFloorWindow(wind, ContentPaint, PsWindow)
		}
		if game == "nonogram" {
			content := makeLevelNonogram(increamentLevel(LevelInProggressNonogram))
			btnSetContent(wind, content)
		}

		win.Close()

	}
	win.SetContent(mainContentContainer)
	win.CenterOnScreen()

	win.Resize(fyne.NewSize(400, 220))
	win.Show()
}

func btnSetContent(win fyne.Window, mainContent fyne.CanvasObject) {
	win.SetContent(mainContent)

}
func increamentLevel(level string) string {
	num, err := strconv.Atoi(level)
	if err != nil {
		return "1"
	}
	num++

	return strconv.Itoa(num)
}
