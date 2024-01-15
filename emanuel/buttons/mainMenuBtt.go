package buttons

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

/*TODO add background color*/

type ColoredTextButton struct {
	widget.BaseWidget
	Text        string
	TextColor   color.Color
	ButtonColor color.Color
	OnTapped    func()
	hovered     bool
}

func NewColoredTextButton(text string, textColor, btnColor color.Color, tapped func()) *ColoredTextButton {
	btn := &ColoredTextButton{
		Text:        text,
		TextColor:   textColor,
		ButtonColor: btnColor,
		OnTapped:    tapped,
	}
	btn.ExtendBaseWidget(btn)
	return btn
}

func (b *ColoredTextButton) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped()
	}
}

func (b *ColoredTextButton) MouseIn(*fyne.PointEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *ColoredTextButton) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (b *ColoredTextButton) CreateRenderer() fyne.WidgetRenderer {

	label := canvas.NewText(b.Text, b.TextColor)
	label.Alignment = fyne.TextAlignCenter
	bg := canvas.NewRectangle(b.ButtonColor)

	underlyingButton := widget.NewButton("", b.OnTapped)
	return &coloredTextButtonRenderer{label: label, bg: bg, button: underlyingButton, btn: b}
}

type coloredTextButtonRenderer struct {
	label  *canvas.Text
	bg     *canvas.Rectangle
	button *widget.Button
	btn    *ColoredTextButton
}

func (r *coloredTextButtonRenderer) MinSize() fyne.Size {
	labelMinSize := r.label.MinSize()
	// Add padding or any other adjustments to account for the custom content.
	padding := fyne.NewSize(theme.Padding()*24, theme.Padding()*12)
	return labelMinSize.Add(padding)
}

func (r *coloredTextButtonRenderer) Refresh() {
	r.label.Text = r.btn.Text
	r.bg.FillColor = r.btn.ButtonColor
	if r.btn.hovered {
		r.bg.FillColor = DarkenColor(r.btn.ButtonColor)
		r.label.Color = DarkenColor(r.btn.TextColor)
	}
	r.bg.Refresh()
	r.label.Refresh()
	r.button.Refresh()

}

func (r *coloredTextButtonRenderer) Layout(size fyne.Size) {
	r.bg.Resize(size)
	r.button.Resize(size)
	r.label.Resize(size)
}

func (r *coloredTextButtonRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *coloredTextButtonRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.button, r.label}
}

func (r *coloredTextButtonRenderer) Destroy() {}

// Utility function to darken a color
func DarkenColor(c color.Color) color.Color {
	const factor = 0.8
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(float32(r>>8) * factor),
		G: uint8(float32(g>>8) * factor),
		B: uint8(float32(b>>8) * factor),
		A: uint8(a >> 8),
	}
}
