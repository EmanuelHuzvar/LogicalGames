package gamesWindows

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

//go:embed pictures/check.png
var checkImg []byte

//go:embed pictures/WinLevelImg.png
var winLevelImg []byte

//go:embed pictures/looseIcon.png
var LooseLevelImg []byte

func MakeWinWindow(mainApp fyne.App, wind fyne.Window, mainContent fyne.CanvasObject, game string) {

	checkImgResource := fyne.NewStaticResource("check.png", checkImg)
	WinWindowImgResource := fyne.NewStaticResource("WinLevelImg.png", winLevelImg)
	checkImgWidget := canvas.NewImageFromResource(checkImgResource)
	windWindowWidget := canvas.NewImageFromResource(WinWindowImgResource)

	checkImgWidget.SetMinSize(fyne.NewSize(128, 128))
	btn := widget.Button{Text: "next level"}
	btn2 := widget.Button{Text: "games"}
	btn3 := widget.Button{Text: "levels"}
	winLabel := canvas.Text{Text: "You successfully solved the level"}
	winLabel.TextSize = 24
	winLabel.Color = color.RGBA{
		R: 128,
		G: 128,
		B: 128,
		A: 255,
	}
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
			level := increamentLevel(LevelInProggress)

			if level == "15" {
				level = "1"
			}

			content := makeLevel(level)
			btnSetContent(wind, content)
		}
		if game == "paint" {
			level := increamentLevel(LevelInProggressPaint)
			fmt.Println(level)
			if level == "15" {
				level = "1"
			}
			fmt.Println(level)
			fmt.Println(level)
			fmt.Println(level)
			fmt.Println(level)
			fmt.Println(level)
			content := MakeGamePaintFloor(level)
			containeris := content.Content()
			btnSetContent(wind, containeris)
			SetUpPaintFloorWindow(wind, ContentPaint, PsWindow)
		}
		if game == "nonogram" {
			level := increamentLevel(LevelInProggressNonogram)
			if level == "15" {
				level = "1"
			}
			content := makeLevelNonogram(level)
			btnSetContent(wind, content)
		}

		win.Close()

	}
	windWindowWidget.Translucency = 0.15
	combinedContainer := container.NewStack(
		windWindowWidget,
		mainContentContainer,
	)
	win.SetContent(combinedContainer)
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
		fmt.Println(err)
		return "1"
	}
	num++
	return strconv.Itoa(num)
}
