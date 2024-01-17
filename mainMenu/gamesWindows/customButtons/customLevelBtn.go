package customButtons

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image/color"
)

type ImageButton struct {
	widget.BaseWidget
	image  *canvas.Image
	label  *canvas.Text
	tapped func()
}

func NewImageButton(imageResource fyne.Resource, text string, tapped func()) *ImageButton {
	img := canvas.NewImageFromResource(imageResource)
	lbl := canvas.NewText(text, color.White)
	lbl.TextStyle.Bold = true

	btn := &ImageButton{
		image:  img,
		label:  lbl,
		tapped: tapped,
	}
	btn.ExtendBaseWidget(btn)
	return btn
}

func (b *ImageButton) CreateRenderer() fyne.WidgetRenderer {
	b.label.Color = colornames.Black
	b.label.TextSize = 25
	b.image.Translucency = 0.45
	return &imageButtonRenderer{
		button:  b,
		image:   b.image,
		label:   b.label,
		objects: []fyne.CanvasObject{b.image, b.label},
	}
}

type imageButtonRenderer struct {
	button  *ImageButton
	image   *canvas.Image
	label   *canvas.Text
	objects []fyne.CanvasObject
}

func (r *imageButtonRenderer) MinSize() fyne.Size {
	return r.image.MinSize().Max(r.label.MinSize())
}

func (r *imageButtonRenderer) Layout(size fyne.Size) {
	r.image.Resize(size)

	// Calculate the size of the text to center it
	textSize := r.label.MinSize()
	r.label.Resize(textSize)
	r.label.Move(fyne.NewPos(
		(size.Width-textSize.Width)/2,   // Center horizontally
		(size.Height-textSize.Height)/2, // Center vertically
	))
}

func (r *imageButtonRenderer) Refresh() {
	r.label.Color = colornames.Black // Ensure the text color is black
	r.label.TextSize = 25            // Adjust the text size as needed
	r.label.Alignment = fyne.TextAlignCenter
	r.image.Refresh()
	r.label.Refresh()
}

func (r *imageButtonRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *imageButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *imageButtonRenderer) Destroy() {}

func (b *ImageButton) Tapped(_ *fyne.PointEvent) {
	if b.tapped != nil {
		b.tapped()
	}
}

func (b *ImageButton) TappedSecondary(_ *fyne.PointEvent) {}
