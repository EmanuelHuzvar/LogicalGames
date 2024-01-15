package customButtons

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type CustomRadioButton struct {
	widget.BaseWidget
	Selected     string
	SelectIcon   *canvas.Image
	DeselectIcon *canvas.Image
}

//go:embed picture/close.png
var deselectIconImg []byte

//go:embed picture/square.png
var selectIconImg []byte

func NewCustomRadioButton() *CustomRadioButton {
	imgResource := fyne.NewStaticResource("close.png", deselectIconImg)
	deselectionIcon := canvas.NewImageFromResource(imgResource)
	imgResourceSel := fyne.NewStaticResource("close.png", selectIconImg)
	selectionIcon := canvas.NewImageFromResource(imgResourceSel)

	r := &CustomRadioButton{
		SelectIcon:   selectionIcon,
		DeselectIcon: deselectionIcon,
	}

	r.ExtendBaseWidget(r)
	return r
}

func (r *CustomRadioButton) Tapped(event *fyne.PointEvent) {
	if r.Selected == "x" {

		r.Selected = "square"
	} else {
		r.Selected = "x"
	}

	fmt.Println(r)
	r.Refresh()
}

func (r *CustomRadioButton) CreateRenderer() fyne.WidgetRenderer {
	icon := r.DeselectIcon
	if r.Selected == "square" {
		icon = r.SelectIcon
	}

	return &CustomRadioButtonRenderer{
		radioButton: r,
		icon:        icon,
	}
}

type CustomRadioButtonRenderer struct {
	radioButton *CustomRadioButton
	icon        *canvas.Image
}

func (r *CustomRadioButtonRenderer) Layout(size fyne.Size) {
	r.icon.Resize(size)
}

func (r *CustomRadioButtonRenderer) MinSize() fyne.Size {
	return fyne.NewSize(30, 30) // Set the minimum size for the custom radio button
}

func (r *CustomRadioButtonRenderer) Refresh() {
	canvas.Refresh(r.icon)
	if r.radioButton.Selected == "x" {
		r.icon = r.radioButton.SelectIcon

	} else {
		r.icon = r.radioButton.DeselectIcon

	}

}

func (r *CustomRadioButtonRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *CustomRadioButtonRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.icon}
}

func (r *CustomRadioButtonRenderer) Destroy() {
	//
}
