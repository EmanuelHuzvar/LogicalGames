package helpers

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ImgToCanvasObject(imgData []byte, fallbackPath string) *canvas.Image {
	var imgResource fyne.Resource

	if len(imgData) > 0 {
		imgResource = fyne.NewStaticResource("embeddedImage", imgData)
		fmt.Println("Loaded embedded image")
	} else {
		imgResource = fyne.NewStaticResource(fallbackPath, imgData)
		fmt.Println("Fallback to", fallbackPath)
	}

	img := canvas.NewImageFromResource(imgResource)
	img.FillMode = canvas.ImageFillContain
	return img
}
