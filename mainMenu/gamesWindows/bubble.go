package gamesWindows

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image/color"
	"time"
)

const maxBubblesPerTube = 4

type BubbleScreen struct {
	window          fyne.Window
	app             fyne.App
	mainMenuContent fyne.CanvasObject
	level           string
}
type Vector2 struct {
	X float32
	Y float32
}
type Bunk struct {
	Buttons []*CircleButton
}

var LevelInProggress string

//var Bunks []Bunk
//var whiteColor = colornames.White
//var mainApp fyne.App

var (
	Bunks        []Bunk
	mainApp      fyne.App
	levelForUse  Level
	colorNameMap map[color.Color]string
)
var (
	whiteColor   = colornames.White
	redColor     = colornames.Red
	greenColor   = colornames.Green
	blueColor    = colornames.Blue
	cyanColor    = colornames.Cyan
	purpleColor  = colornames.Purple
	magentaColor = colornames.Magenta
	orangeColor  = colornames.Orange
	yellowColor  = colornames.Yellow
)

//go:embed pictures/backroundBubble.jpg
var backgroundImg []byte

//go:embed pictures/backgroundimgnonogram.png
var backgroundImgNonogram []byte

//go:embed pictures/back-arrow.png
var backBtnImg []byte

//go:embed pictures/red.png
var red []byte

//go:embed pictures/green.jpg
var green []byte

//go:embed pictures/orange.jpg
var orange []byte

func NewBubbleScreen(window fyne.Window, app fyne.App, mainMenuContent fyne.CanvasObject, level string) *BubbleScreen {
	mainApp = app
	return &BubbleScreen{window: window, app: app, mainMenuContent: mainMenuContent, level: level}
}

func init() {
	colorNameMap = make(map[color.Color]string)
	for name, c := range colornames.Map {
		colorNameMap[c] = name
	}
}

var strokeColor = colornames.Black
var lastClickedButton *CircleButton = nil
var wind fyne.Window
var mainContent fyne.CanvasObject
var contentOfWindow fyne.Window

const (
	strokeWidth = float32(5)
)

type CircleButton struct {
	widget.BaseWidget
	OnTapped       func(*CircleButton)
	FillColor      color.Color
	StrokeColor    color.Color
	StrokeWidth    float32
	PositionOff    bool
	Disabled       bool
	PositionInBunk [2]int
}

func NewCircleButton(fillColor, strokeColor color.Color, strokeWidth float32, onTapped func(*CircleButton)) *CircleButton {
	button := &CircleButton{
		OnTapped:    onTapped,
		FillColor:   fillColor,
		StrokeColor: strokeColor,
		StrokeWidth: strokeWidth,
	}
	button.ExtendBaseWidget(button)
	return button
}

func (b *CircleButton) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
	if b.Disabled {
		return
	}

	if lastClickedButton != nil && lastClickedButton != b {
		swapColors(lastClickedButton, b)
		MovePositionDown(b)
		lastClickedButton = nil
	} else {
		lastClickedButton = b
	}
	if checkWinCondition(Bunks) {
		MakeWinWindow(mainApp, wind, mainContent, "bubble")
	}
}

func (b *CircleButton) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	circle := canvas.NewCircle(b.FillColor)
	circle.StrokeColor = b.StrokeColor
	circle.StrokeWidth = b.StrokeWidth

	objects := []fyne.CanvasObject{circle}
	return &circleButtonRenderer{
		button:  b,
		circle:  circle,
		objects: objects,
	}
}

type circleButtonRenderer struct {
	button  *CircleButton
	circle  *canvas.Circle
	objects []fyne.CanvasObject
}

func (r *circleButtonRenderer) Destroy() {}

func (r *circleButtonRenderer) Layout(size fyne.Size) {
	r.circle.Resize(size)
}

func (r *circleButtonRenderer) MinSize() fyne.Size {
	return r.circle.MinSize().Add(fyne.NewSize(r.circle.StrokeWidth*10, r.circle.StrokeWidth*10))
}

func (r *circleButtonRenderer) Refresh() {
	r.circle.FillColor = r.button.FillColor
	r.circle.StrokeColor = r.button.StrokeColor
	r.circle.StrokeWidth = r.button.StrokeWidth
	canvas.Refresh(r.circle)
}

func (r *circleButtonRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *circleButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (bs *BubbleScreen) Render() {
	wind = bs.window
	mainContent = bs.mainMenuContent

	app := app.New()
	fmt.Println(bs.level)
	menuWindow := app.NewWindow("Logitec App")

	menuWindow.SetOnClosed(func() {
		app.Quit()
	})

	cont := makeLevel(bs.level)

	bs.window.SetContent(cont)
	bs.window.CenterOnScreen()

}

func createCircleButton(fillColor, strokeColor color.Color, strokeWidth float32) *CircleButton {
	circleBtn := NewCircleButton(fillColor, strokeColor, strokeWidth, func(btn *CircleButton) {
		if btn.Disabled {
			return
		}
		fmt.Println("Button clicked:", btn)
		if lastClickedButton != nil {
			if lastClickedButton.PositionOff {
				MovePositionDown(lastClickedButton)
			}
		}
		if btn.PositionOff {
			MovePositionDown(btn)
		} else {
			MovePositionUp(btn)
		}

	})

	for i, bunk := range Bunks {

		for j := 0; j < 4; j++ {
			if bunk.Buttons[j].PositionInBunk == [2]int{} {
				bunk.Buttons[j].PositionInBunk = [2]int{i, j}
			}

		}

	}

	circleBtn.PositionOff = false
	return circleBtn
}
func createFourByOneBunk(colorOne color.Color, colorTwo color.Color, colorThree color.Color, colorFour color.Color) *fyne.Container {
	circleOne := createCircleButton(colorOne, strokeColor, strokeWidth)
	circleTwo := createCircleButton(colorTwo, strokeColor, strokeWidth)
	circleThree := createCircleButton(colorThree, strokeColor, strokeWidth)
	circleFour := createCircleButton(colorFour, strokeColor, strokeWidth)
	btns := []*CircleButton{circleOne, circleTwo, circleThree, circleFour}
	smallBunk := Bunk{btns}
	if len(Bunks) == 10 {
		circleOne.PositionInBunk = [2]int{9, 0}
		circleTwo.PositionInBunk = [2]int{9, 1}
		circleThree.PositionInBunk = [2]int{9, 2}
		circleFour.PositionInBunk = [2]int{9, 3}
	}

	addToBunks(smallBunk)

	circleTwo.Disabled = true
	circleThree.Disabled = true
	circleFour.Disabled = true

	bunkContainer := container.NewVBox(
		circleOne,
		circleTwo,
		circleThree,
		circleFour,
	)

	return bunkContainer
}
func createFourByOneBunkWhite() *fyne.Container {
	circleOne := createCircleButton(colornames.White, strokeColor, strokeWidth)
	circleTwo := createCircleButton(colornames.White, strokeColor, strokeWidth)
	circleThree := createCircleButton(colornames.White, strokeColor, strokeWidth)
	circleFour := createCircleButton(colornames.White, strokeColor, strokeWidth)
	btns := []*CircleButton{circleOne, circleTwo, circleThree, circleFour}
	smallBunk := Bunk{btns}
	addToBunks(smallBunk)
	if len(Bunks) == 10 {
		circleOne.PositionInBunk = [2]int{9, 0}
		circleTwo.PositionInBunk = [2]int{9, 1}
		circleThree.PositionInBunk = [2]int{9, 2}
		circleFour.PositionInBunk = [2]int{9, 3}
	}
	if len(Bunks) == 6 {
		circleOne.PositionInBunk = [2]int{5, 0}
		circleTwo.PositionInBunk = [2]int{5, 1}
		circleThree.PositionInBunk = [2]int{5, 2}
		circleFour.PositionInBunk = [2]int{5, 3}
	}
	if len(Bunks) == 8 {
		circleOne.PositionInBunk = [2]int{7, 0}
		circleTwo.PositionInBunk = [2]int{7, 1}
		circleThree.PositionInBunk = [2]int{7, 2}
		circleFour.PositionInBunk = [2]int{7, 3}
	}
	bunkContainer := container.NewVBox(
		circleOne,
		circleTwo,
		circleThree,
		circleFour,
	)

	return bunkContainer
}
func makeLevel(level string) *fyne.Container {
	BackBtnImgResource := fyne.NewStaticResource("back-arrow.png", backBtnImg)
	imgResource := fyne.NewStaticResource("backroundBubble.jpg", backgroundImg)
	backgroundImage := canvas.NewImageFromResource(imgResource)
	backgroundImage.FillMode = canvas.ImageFillStretch
	emptyBunk := []Bunk{}
	Bunks = emptyBunk
	lev, _ := GetLevelByID(level, "Bubble")
	fmt.Println(lev.Fields)
	bunkForDb := []*fyne.Container{}
	for i := 0; i < len(lev.Fields); i++ {

		bun := createFourByOneBunk(MatchColorFromDB(lev.Fields[i][0]), MatchColorFromDB(lev.Fields[i][1]), MatchColorFromDB(lev.Fields[i][2]), MatchColorFromDB(lev.Fields[i][3]))
		bunkForDb = append(bunkForDb, bun)

	}

	gridLength := (len(bunkForDb) + 2) / 2
	grid := container.New(layout.NewGridLayout(gridLength))

	grid2 := container.New(layout.NewGridLayout(gridLength))

	for i := 0; i < len(bunkForDb)/2; i++ {
		grid.Add(bunkForDb[i])
	}
	for i := len(bunkForDb) / 2; i < len(bunkForDb); i++ {
		grid2.Add(bunkForDb[i])
	}
	grid.Add(createFourByOneBunkWhite())
	grid2.Add(createFourByOneBunkWhite())
	backButton := widget.NewButtonWithIcon("", BackBtnImgResource, func() {
		wind.SetContent(mainContent)
	})
	topLeftContainer := container.NewVBox(
		backButton,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	finalContainer := container.NewHBox(

		topLeftContainer,
		layout.NewSpacer(),
	)
	topContainer := container.NewVBox(
		finalContainer,
		layout.NewSpacer(),
		grid,
		layout.NewSpacer(),
		grid2,
		layout.NewSpacer(),
	)
	backgroundImage.Translucency = 0.55
	combinedContainer := container.NewStack(
		backgroundImage,
		topContainer,
	)

	startBunksToColorRoutine(Bunks)
	return combinedContainer

}

func MovePositionUp(btn *CircleButton) {

	currentPosition := btn.Position()
	offset := fyne.NewSize(0, -40)
	newPosition := currentPosition.Add(offset)

	btn.PositionOff = true
	btn.Move(newPosition)

}
func MovePositionDown(btn *CircleButton) {
	currentPosition := btn.Position()
	offset := fyne.NewSize(0, 40)
	newPosition := currentPosition.Add(offset)
	btn.PositionOff = false
	btn.Move(newPosition)

}
func swapColors(b1, b2 *CircleButton) {
	if b2.FillColor != whiteColor {
		return
	}

	colorOfLastCircle := getLastNonWhiteColorInBunk(Bunks, b2.PositionInBunk[0])
	fmt.Println("last color ", colorOfLastCircle)
	fmt.Println(b1.FillColor, colorOfLastCircle)

	if colorOfLastCircle == nil {
		colorOfLastCircle = whiteColor
	}

	if colorOfLastCircle != b1.FillColor && colorOfLastCircle != whiteColor {
		return
	}

	fmt.Println(b1.FillColor, colorOfLastCircle)

	lastWhiteButton := getLastWhiteButtonInBunk(Bunks, b2.PositionInBunk[0])
	if lastWhiteButton == nil {
		return // No white button to swap with
	}

	b1.FillColor, lastWhiteButton.FillColor = lastWhiteButton.FillColor, b1.FillColor

	b1.Refresh()
	lastWhiteButton.Refresh()
}

func addToBunks(bunk Bunk) {
	Bunks = append(Bunks, bunk)
}
func bunksToColor(bunks []Bunk) {
	for _, bunk := range bunks {
		for j := 0; j < 3; j++ {
			if bunk.Buttons[j].FillColor == whiteColor {
				bunk.Buttons[j+1].Disabled = false
			} else {
				bunk.Buttons[j+1].Disabled = true
			}
		}

	}

}
func startBunksToColorRoutine(bunks []Bunk) {
	ticker := time.NewTicker(300 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				bunksToColor(bunks)
			}
		}
	}()
}
func ColorName(c color.Color) string {
	if name, ok := colorNameMap[c]; ok {
		return name
	}
	return "unknown" // or any default value you prefer
}

//getLastNonWhiteColorInBunk

func getLastNonWhiteColorInBunk(bunks []Bunk, bunkIndex int) color.Color {
	if bunkIndex < 0 || bunkIndex >= len(bunks) {
		return nil // Return zero value if the index is out of range
	}

	bunk := bunks[bunkIndex]

	// Iterate over the buttons from the start to find the first non-white color
	for i := 0; i < len(bunk.Buttons); i++ {
		if !sameColor(bunk.Buttons[i].FillColor, whiteColor) {
			return bunk.Buttons[i].FillColor
		}
	}

	return nil // Return nil if all buttons are white or there are no buttons
}

func sameColor(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}
func getLastWhiteButtonInBunk(bunks []Bunk, bunkIndex int) *CircleButton {
	if bunkIndex < 0 || bunkIndex >= len(bunks) {
		return nil
	}

	bunk := bunks[bunkIndex]
	var lastWhiteButton *CircleButton = nil

	// Iterate over the buttons in reverse to find the last white button
	for i := len(bunk.Buttons) - 1; i >= 0; i-- {
		if bunk.Buttons[i].FillColor == whiteColor {
			lastWhiteButton = bunk.Buttons[i]
			break
		}
	}

	return lastWhiteButton
}

func checkWinCondition(bunks []Bunk) bool {
	for _, bunk := range bunks {
		if !isBunkUniform(bunk) {
			return false
		}
	}
	return true
}

func isBunkUniform(bunk Bunk) bool {
	if len(bunk.Buttons) == 0 {
		return false
	}

	firstButtonColor := bunk.Buttons[0].FillColor

	for _, btn := range bunk.Buttons {
		if btn.FillColor != firstButtonColor {
			return false
		}
	}
	return true
}

func MatchColorFromDB(color string) color.Color {
	// red ,green , blue ,cyan ,purple ,magenta ,orange ,yellow
	if color == "red" {
		return redColor
	}
	if color == "green" {
		return greenColor
	}
	if color == "blue" {
		return blueColor
	}
	if color == "cyan" {
		return cyanColor
	}
	if color == "purple" {
		return purpleColor
	}
	if color == "magenta" {
		return magentaColor
	}
	if color == "orange" {
		return orangeColor
	}
	if color == "yellow" {
		return yellowColor
	}

	return nil
}
