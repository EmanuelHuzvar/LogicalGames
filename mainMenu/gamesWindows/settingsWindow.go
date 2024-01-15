package gamesWindows

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//go:embed pictures/sun.png
var sunImg []byte

//go:embed pictures/volume.png
var volumeImg []byte

//go:embed pictures/user.png
var personImg []byte

//go:embed pictures/email.png
var emailImg []byte

//go:embed pictures/social.png
var shareImg []byte

//go:embed pictures/support.png
var helpImg []byte

func MakeOptionWindow(mainApp fyne.App) {
	// Image Resources
	imgResource := fyne.NewStaticResource("sun.png", sunImg)
	sunImgWidget := canvas.NewImageFromResource(imgResource)
	volumeResource := fyne.NewStaticResource("volume.png", volumeImg)
	volumeImgWidget := canvas.NewImageFromResource(volumeResource)
	personResource := fyne.NewStaticResource("user.png", personImg)
	personImgWidget := canvas.NewImageFromResource(personResource)
	emailResource := fyne.NewStaticResource("email.png", emailImg)
	emailImgWidget := canvas.NewImageFromResource(emailResource)
	socialResource := fyne.NewStaticResource("social.png", shareImg)
	socialImgWidget := canvas.NewImageFromResource(socialResource)
	supportResource := fyne.NewStaticResource("support.png", helpImg)
	supportImgWidget := canvas.NewImageFromResource(supportResource)

	sunImgWidget.SetMinSize(fyne.NewSize(50, 50))
	volumeImgWidget.SetMinSize(fyne.NewSize(50, 50))
	personImgWidget.SetMinSize(fyne.NewSize(50, 50))
	emailImgWidget.SetMinSize(fyne.NewSize(50, 50))
	socialImgWidget.SetMinSize(fyne.NewSize(50, 50))
	supportImgWidget.SetMinSize(fyne.NewSize(50, 50))

	brightnessSlider := widget.NewSlider(0, 100)
	volumeSlider := widget.NewSlider(0, 100)

	personButton := widget.NewButton("YOUR NAME", func() {
		handleButtonClick() // The function that handles the button click
	})

	brightnessSlider.Resize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	sunContainer := container.NewCenter(
		sunImgWidget,
	)
	volumeContainer := container.NewCenter(
		volumeImgWidget,
	)
	personContainer := container.NewCenter(
		personImgWidget,
	)

	emailButton := widget.NewButtonWithIcon("", emailResource, func() {
		// Code to execute when email image is clicked
	})
	emailButton.Importance = widget.LowImportance
	emailButton.MinSize()

	downContainer := container.NewHBox(
		emailImgWidget,
		layout.NewSpacer(),
		socialImgWidget,
		layout.NewSpacer(),
		supportImgWidget,
	)

	// Main Content
	mainContent := container.NewVBox(
		sunContainer,
		brightnessSlider,
		volumeContainer,
		volumeSlider,
		personContainer,
		personButton,
		layout.NewSpacer(),
		downContainer,
	)

	// Window Settings
	win := mainApp.NewWindow("Info")
	win.SetContent(mainContent)
	win.Resize(fyne.NewSize(400, 500))
	win.CenterOnScreen()
	win.Show()
}

func handleButtonClick() {
	/*TODO MAKE WINDOW THAT OPENS SING IN*/
}
func sunImgClicked() {
	fmt.Println("marek")
}
