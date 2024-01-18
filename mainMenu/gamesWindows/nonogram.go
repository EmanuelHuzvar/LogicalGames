package gamesWindows

import (
	"ProjectMarekEmanuel/mainMenu/gamesWindows/customButtons"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image/color"
	"strconv"
)

type CellState int

const strokeWidthNanogram = 5
const strokeWidthLine = strokeWidthNanogram * 2
const (
	StateEmpty CellState = iota
	StateMarked
	StateFilled
)

var level LevelNonogram
var gridSize int
var colsNumbers []int
var rowsNumbers []int
var colsWinNumbers []int
var rowWinNumbers []int
var cellsOfNonogram []*NonogramCell
var cellsOfNonogramWin []*NonogramCell
var RadioBtn *customButtons.CustomRadioButton

type NonogramCell struct {
	widget.BaseWidget
	state       CellState
	Pos         [][]int
	FillColor   color.Color // Add fill color
	StrokeColor color.Color // Add stroke color
	StrokeWidth float32
}

func NewNonogramCell(fillColor, strokeColor color.Color, strokeWidth float32) *NonogramCell {
	cell := &NonogramCell{
		state:       StateEmpty,
		FillColor:   fillColor,
		StrokeColor: strokeColor,
		StrokeWidth: strokeWidth,
	}
	cell.ExtendBaseWidget(cell)
	cell.Refresh() // Important for lifecycle management
	return cell
}

func (c *NonogramCell) Tapped(event *fyne.PointEvent) {
	fmt.Println(c)
	//rowOfCell := c.Pos[0][0]
	//colOfCell := c.Pos[0][1]

	if RadioBtn.Selected == "x" && c.state == 1 {
		c.state = 0
		c.Refresh()
		return
	}
	if RadioBtn.Selected == "x" && c.state != 1 {
		c.state = 1
		c.Refresh()
		if checkWinConditionNonogram(cellsOfNonogram, cellsOfNonogramWin) {
			MakeWinWindow(mainApp, wind, mainContent, "nonogram")
		}
		return
	}
	if RadioBtn.Selected == "square" && c.state == 2 {
		c.state = 0
		c.Refresh()
		return
	}
	if RadioBtn.Selected == "square" && c.state != 2 {
		c.state = 2
		c.Refresh()
		if checkWinConditionNonogram(cellsOfNonogram, cellsOfNonogramWin) {
			MakeWinWindow(mainApp, wind, mainContent, "nonogram")
		}
		return
	}

	c.Refresh()

}

func (c *NonogramCell) CreateRenderer() fyne.WidgetRenderer {
	// Initial setup for the rectangle, this will be the background
	rect := canvas.NewRectangle(color.White)
	rect.StrokeColor = colornames.Gray // Set the color for the border
	rect.StrokeWidth = 1               // Set the width of the border

	// The cross mark can be represented by two lines
	line1 := canvas.NewLine(color.Black)
	line2 := canvas.NewLine(color.Black)
	line1.Hidden = true // Initially hidden
	line2.Hidden = true // Initially hidden
	//TODO
	line1.StrokeWidth = strokeWidthLine
	line2.StrokeWidth = strokeWidthLine

	return &NonogramCellRenderer{
		cell:  c,
		rect:  rect,
		line1: line1,
		line2: line2,
	}
}

type NonogramCellRenderer struct {
	cell  *NonogramCell
	rect  *canvas.Rectangle
	line1 *canvas.Line
	line2 *canvas.Line
}

func (r *NonogramCellRenderer) Layout(size fyne.Size) {
	// Set the size for the rectangle and lines
	r.rect.Resize(size)

	padding := size.Width / 8
	r.line1.Position1 = fyne.NewPos(padding, padding)
	r.line1.Position2 = fyne.NewPos(size.Width-padding, size.Height-padding)
	r.line2.Position1 = fyne.NewPos(size.Width-padding, padding)
	r.line2.Position2 = fyne.NewPos(padding, size.Height-padding)

}

func (r *NonogramCellRenderer) MinSize() fyne.Size {
	// Minimum size of the cell
	return fyne.NewSize(30, 30)
}

func (r *NonogramCellRenderer) Refresh() {
	switch r.cell.state {
	case StateEmpty:
		r.rect.FillColor = color.White
		r.line1.Hidden = true
		r.line2.Hidden = true
	case StateMarked:
		r.rect.FillColor = color.Black
		r.line1.Hidden = false
		r.line2.Hidden = false
	case StateFilled:
		r.rect.FillColor = color.White
		r.line1.Hidden = false
		r.line2.Hidden = false
	}
	r.rect.StrokeColor = colornames.Black
	r.rect.StrokeWidth = strokeWidthNanogram
	canvas.Refresh(r.rect)
	canvas.Refresh(r.line1)
	canvas.Refresh(r.line2)
}

func (r *NonogramCellRenderer) BackgroundColor() color.Color {
	return color.Transparent
}
func (r *NonogramCellRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect, r.line1, r.line2}
}

func (r *NonogramCellRenderer) Destroy() {
	// Destroy any resources created by the renderer if necessary
}

type NonogramScreen struct {
	window          fyne.Window
	app             fyne.App
	mainMenuContent fyne.CanvasObject
	level           string
}

var LevelInProggressNonogram string

func NewNonogramScreen(window fyne.Window, app fyne.App, mainMenuContent fyne.CanvasObject, level string) *NonogramScreen {
	mainApp = app
	return &NonogramScreen{window: window, app: app, mainMenuContent: mainMenuContent, level: level}
}

func (ns *NonogramScreen) Render() {
	wind = ns.window
	mainContent = ns.mainMenuContent

	app := app.New()

	menuWindow := app.NewWindow("Logitec App")

	menuWindow.SetOnClosed(func() {
		app.Quit()
	})

	cont := makeLevelNonogram(ns.level)

	ns.window.SetContent(cont)

	ns.window.CenterOnScreen()

}

func makeLevelNonogram(level string) *fyne.Container {
	colsNumbers = nil
	rowsNumbers = nil
	colsWinNumbers = nil
	cellsOfNonogramWin = nil
	gird := createNonogramGridWithClues(level)
	LevelInProggressNonogram = level
	return gird

}
func createNonogramGridWithClues(nonogramLevel string) *fyne.Container {
	imgResource := fyne.NewStaticResource("backgroundimgnonogram.png", backgroundImgNonogram)
	backgroundImage := canvas.NewImageFromResource(imgResource)
	backgroundImage.FillMode = canvas.ImageFillStretch
	BackBtnImgResource := fyne.NewStaticResource("back-arrow.png", backBtnImg)
	colsNumbers = nil
	rowsNumbers = nil
	colsWinNumbers = nil
	cellsOfNonogramWin = nil
	cellsOfNonogram = nil
	LevelInProggressNonogram = ""
	level = LevelNonogram{}
	level, _ = GetLevelNonogramByID(nonogramLevel)
	LevelInProggressNonogram = nonogramLevel
	cols := len(level.Cols)
	rows := len(level.Rows)
	gridSize = len(level.Cols)

	colsNumbers = level.Cols
	rowsNumbers = level.Rows
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

	colsWinNumbers, _ = binaryStringsToInts(level.ColsWin)
	cellsOfNonogramWin = createWinConditionGrid(gridSize, colsWinNumbers)

	fmt.Println(level.ColsWin)
	// Create the grid container with extra space for clues
	grid := container.New(layout.NewGridLayoutWithColumns(cols + 1))

	// Add column clues at the top (for simplicity, here as labels)
	for i := 0; i <= cols; i++ {
		if i == 0 {
			// First cell is empty because it's the corner of the clues
			grid.Add(finalContainer)
		} else {
			if colsNumbers[i-1] > 100 {
				firstNum := colsNumbers[i-1] / 100
				secondNum := (colsNumbers[i-1] / 10) % 10
				thirdNum := colsNumbers[i-1] % 10

				label := widget.NewLabel(fmt.Sprintf("%d\n%d\n%d", firstNum, secondNum, thirdNum))

				grid.Add(label)
			} else if colsNumbers[i-1] > 10 {
				firstNum := colsNumbers[i-1] / 10
				secondNum := colsNumbers[i-1] % 10

				label := widget.NewLabel(fmt.Sprintf("%d\n%d", firstNum, secondNum))

				grid.Add(label)
			} else {

				grid.Add(widget.NewLabel(fmt.Sprintf("%d", colsNumbers[i-1])))
			}

		}
	}

	// Add the Nonogram cells and row clues
	for y := 0; y < rows; y++ {
		if colsNumbers[y] > 100 {
			firstNum := rowsNumbers[y] / 100
			secondNum := (rowsNumbers[y] / 10) % 10
			thirdNum := rowsNumbers[y] % 10

			label := widget.NewLabel(fmt.Sprintf("%d\n%d\n%d", firstNum, secondNum, thirdNum))

			grid.Add(label)
		} else if rowsNumbers[y] > 10 {
			firstNum := rowsNumbers[y] / 10
			secondNum := rowsNumbers[y] % 10
			grid.Add(widget.NewLabel(fmt.Sprintf("%d     %d", firstNum, secondNum)))
		} else {
			grid.Add(widget.NewLabel(fmt.Sprintf("%d", rowsNumbers[y])))
		}
		// Add row clue before each row (for simplicity, here as a label)

		// Add cells for the current row
		for x := 0; x < cols; x++ {
			cell := NewNonogramCell(colornames.White, colornames.Black, strokeWidthNanogram)

			poss := [][]int{
				{y, x},
			}

			cell.Pos = poss
			cellsOfNonogram = append(cellsOfNonogram, cell)
			grid.Add(cell)

		}
	}

	radio := customButtons.NewCustomRadioButton()

	grid.Add(radio)

	radio.Selected = "x"
	radio.Refresh()
	RadioBtn = radio
	backgroundImage.Translucency = 0.90
	combinedContainer := container.NewStack(
		backgroundImage,
		grid,
	)
	return combinedContainer
}
func intToString(num int) string {
	return strconv.Itoa(num)
}

func binaryStringsToInts(binStrs []string) ([]int, error) {
	var ints []int
	for _, binStr := range binStrs {
		num, err := strconv.ParseInt(binStr, 2, 64)
		if err != nil {
			// Handle the error in case the binary number is not valid.
			return nil, fmt.Errorf("failed to parse binary string %s to int: %v", binStr, err)
		}
		ints = append(ints, int(num))
	}
	return ints, nil
}
func createWinConditionGrid(gridSize int, colsWinNumbers []int) []*NonogramCell {
	cellsOfNonogramWin := make([]*NonogramCell, gridSize*gridSize)
	for colIndex, colWinNum := range colsWinNumbers {
		for rowIndex := 0; rowIndex < gridSize; rowIndex++ {
			// Calculate the index of the cell in the cellsOfNonogramWin slice
			cellIndex := rowIndex*gridSize + colIndex

			// Create a new NonogramCell for the win condition grid
			cell := NewNonogramCell(colornames.White, colornames.Black, strokeWidthNanogram)
			cell.Pos = [][]int{{

				rowIndex, colIndex}}

			// Check the bit to determine if the cell should be filled
			if (colWinNum>>(gridSize-1-rowIndex))&1 != 0 {
				cell.state = StateMarked
			} else {
				// If the bit is not set, so the cell is either empty or marked
				cell.state = StateEmpty // or StateMarked based on your win condition logic
			}

			// Place the new cell in the cellsOfNonogramWin array
			cellsOfNonogramWin[cellIndex] = cell
		}
	}
	return cellsOfNonogramWin
}
func checkWinConditionNonogram(cellsOfNonogram []*NonogramCell, cellsOfNonogramWin []*NonogramCell) bool {
	fmt.Println(cellsOfNonogram)
	for i, cell := range cellsOfNonogram {

		if cell.state == StateEmpty || cell.state == StateFilled {
			cell.state = StateEmpty
		}

		if cell.state != cellsOfNonogramWin[i].state {

			return false
		}
	}
	return true
}
