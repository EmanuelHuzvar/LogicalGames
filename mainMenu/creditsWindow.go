package mainMenu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type CreditsScreen struct {
	window fyne.Window
	app    fyne.App
}

func NewCreditsScreen(window fyne.Window, app fyne.App) *CreditsScreen {
	return &CreditsScreen{window: window, app: app}
}

func (cs *CreditsScreen) Render() {
	BackBtnImgResource := fyne.NewStaticResource("back-arrow.png", BackBtnImg)
	imgResource := fyne.NewStaticResource("credits-background.png", BackgroundImgCredits)
	backgroundImage := canvas.NewImageFromResource(imgResource)
	backgroundImage.FillMode = canvas.ImageFillStretch
	backgroundImage.Translucency = 0.30
	backButton := widget.NewButtonWithIcon("", BackBtnImgResource, func() {
		cs.window.SetContent(MainMenuContent)
	})
	topLeftContainer := container.NewVBox(
		backButton,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	finalContainerB := container.NewHBox(

		topLeftContainer,
		layout.NewSpacer(),
	)

	// The text for the credits
	credits := []string{
		"",
		"Credits",
		"",
		"",
		"Creators: ",
		"Marek Leščinský",
		"Emanuel Hužvár",
		"",
		"Visuals Provided By:",
		"DALL.E",
		"",
		"Design Approval:",
		"Jan Nosek",
		"",
		"Back-end Approval:",
		"Samuel Mražík",
		"",
		"Main Sponsor:",
		"IT Valley",
		"",
		"Acknowledgments:",
		"Thank YOU for downloading our application",
		"It has been a tremendous pleasure",
		"to develop this game",
		"We hope you enjoy using this application",
		"as much as we've enjoyed creating it",
		"",
		"",
		"Feedback:",
		"For questions, suggestions, or feedback",
		"please reach out to us at",
		"marek.lescinsky@kosickaakademia.com",
		"",
		"",
		"Special Thanks:",
		"Our heartfelt gratitude goes to KASV ",
		"for their unwavering support ",
		"and for making this project possible",
		"",
		"",
		"",
		"P.S.",
		"A big shout-out to all our classmates",
		"at the KASV",
		"Your support, feedback, and comradeship ",
		"have been invaluable",
		"Together, we've turned challenges ",
		"into achievements",
		"Thank you!",
	}

	// Create text objects for each line of credits
	textObjects := make([]*canvas.Text, len(credits))
	for i, line := range credits {
		text := canvas.NewText(line, color.Black)
		text.TextStyle.Bold = true
		text.Alignment = fyne.TextAlignCenter
		text.TextSize = 30 // set your desired text size
		textObjects[i] = text
	}

	// Create a container for text objects
	textContainer := container.NewVBox()
	for _, text := range textObjects {
		textContainer.Add(text)
	}

	// Create a scroll container for the credits
	scroll := container.NewVScroll(textContainer)
	scroll.SetMinSize(fyne.NewSize(400, 300)) // Set a minimum size for the scrolling area

	// Scroll the credits automatically using a goroutine
	go func() {
		ticker := time.NewTicker(50 * time.Millisecond) // Adjust the speed as necessary
		for range ticker.C {
			scroll.Offset.Y += 1 // Move the credits up
			scroll.Refresh()

			// Stop scrolling if we've reached the top
			if scroll.Offset.Y <= -textContainer.MinSize().Height {
				ticker.Stop()
			}
		}
	}()

	// Use a border layout to position the back button and scroll container
	content := container.NewBorder(finalContainerB, nil, nil, nil, scroll)

	// Layer the content over the background image
	finalContainer := container.NewMax(backgroundImage, content)

	cs.window.SetContent(finalContainer)
}
