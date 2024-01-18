package customButtons

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type CustomLabel struct {
	*widget.Label
	text      string
	textColor color.Color
	textSize  float32
}

func NewCustomLabel(text string, size float32, textColor color.Color) *CustomLabel {
	label := &CustomLabel{
		Label:     widget.NewLabel(text),
		text:      text,
		textColor: textColor,
		textSize:  size,
	}

	label.ExtendBaseWidget(label)
	return label
}

func (c *CustomLabel) CreateRenderer() fyne.WidgetRenderer {
	textObj := canvas.NewText(c.text, c.textColor)
	c.textSize = 24
	textObj.TextSize = c.textSize

	textObj.Alignment = fyne.TextAlignCenter

	return &customLabelRenderer{label: c, textObj: textObj}
}

type customLabelRenderer struct {
	label   *CustomLabel
	textObj *canvas.Text
}

func (r *customLabelRenderer) MinSize() fyne.Size {
	return r.textObj.MinSize()
}

func (r *customLabelRenderer) Layout(size fyne.Size) {
	r.textObj.Resize(fyne.NewSize(50, 50))
}

func (r *customLabelRenderer) Refresh() {
	r.textObj.Text = r.label.Text
	r.textObj.Color = r.label.textColor
	r.textObj.TextSize = r.label.textSize
	canvas.Refresh(r.textObj)
}

func (r *customLabelRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *customLabelRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.textObj}
}

func (r *customLabelRenderer) Destroy() {}
