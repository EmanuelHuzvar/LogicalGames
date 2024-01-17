package customButtons

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type ColorButton struct {
	widget.BaseWidget
	background *canvas.Rectangle
	label      *canvas.Text
	tapped     func()
}

func NewColorButton(bgColor color.Color, text string, tapped func()) *ColorButton {
	background := canvas.NewRectangle(bgColor)
	lbl := canvas.NewText(text, color.White)
	lbl.TextStyle.Bold = true

	btn := &ColorButton{
		background: background,
		label:      lbl,
		tapped:     tapped,
	}
	btn.ExtendBaseWidget(btn)
	return btn
}

func (b *ColorButton) CreateRenderer() fyne.WidgetRenderer {
	b.label.Color = color.Black
	b.label.TextSize = 25
	return &colorButtonRenderer{
		button:     b,
		background: b.background,
		label:      b.label,
		objects:    []fyne.CanvasObject{b.background, b.label},
	}
}

type colorButtonRenderer struct {
	button     *ColorButton
	background *canvas.Rectangle
	label      *canvas.Text
	objects    []fyne.CanvasObject
}

func (r *colorButtonRenderer) MinSize() fyne.Size {
	return r.background.MinSize().Max(r.label.MinSize())
}

func (r *colorButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	// Calculate the size of the text to center it
	textSize := r.label.MinSize()
	r.label.Resize(textSize)
	r.label.Move(fyne.NewPos(
		(size.Width-textSize.Width)/2,   // Center horizontally
		(size.Height-textSize.Height)/2, // Center vertically
	))
}

func (r *colorButtonRenderer) Refresh() {
	r.label.Color = color.Black // Ensure the text color is black
	r.label.TextSize = 25       // Adjust the text size as needed
	r.label.Alignment = fyne.TextAlignCenter
	r.background.Refresh()
	r.label.Refresh()
}

func (r *colorButtonRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *colorButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *colorButtonRenderer) Destroy() {}

func (b *ColorButton) Tapped(_ *fyne.PointEvent) {
	if b.tapped != nil {
		b.tapped()
	}
}

func (b *ColorButton) TappedSecondary(_ *fyne.PointEvent) {}
