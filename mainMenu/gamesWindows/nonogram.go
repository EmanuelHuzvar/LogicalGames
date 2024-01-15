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
)

type CellState int

const strokeWidthNanogram = 5
const strokeWidthLine = strokeWidthNanogram * 2
const (
	StateEmpty CellState = iota
	StateMarked
	StateFilled
)

var gridSize int
var colsNumbers []int
var rowsNumbers []int
var cellsOfNonogram []*NonogramCell

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
	backButton := widget.NewButton("Back", func() {
		ns.window.SetContent(LevelContent)
	})

	//would get from db rows and cols
	cont := makeLevelNonogram(ns.level)
	cont.Add(backButton)

	ns.window.SetContent(cont)

	ns.window.CenterOnScreen()

}

func makeLevelNonogram(level string) *fyne.Container {
	gird := createNonogramGridWithClues(level)

	return gird

}
func createNonogramGridWithClues(nonogramLevel string) *fyne.Container {

	colsNumbers = nil
	rowsNumbers = nil

	level, _ := GetLevelNonogramByID(nonogramLevel)

	cols := len(level.Cols)
	rows := len(level.Rows)
	gridSize = len(level.Cols)

	colsNumbers = level.Cols
	rowsNumbers = level.Rows
	// Create the grid container with extra space for clues
	grid := container.New(layout.NewGridLayoutWithColumns(cols + 1))

	// Add column clues at the top (for simplicity, here as labels)
	for i := 0; i <= cols; i++ {
		if i == 0 {
			// First cell is empty because it's the corner of the clues
			grid.Add(widget.NewLabel(""))
		} else {

			if colsNumbers[i-1] > 10 {
				firstNum := colsNumbers[i-1] / 10
				secondNum := colsNumbers[i-1] % 10
				grid.Add(widget.NewLabel(fmt.Sprintf("%d\n%d", firstNum, secondNum)))
			} else {
				grid.Add(widget.NewLabel(fmt.Sprintf("%d", colsNumbers[i-1])))
			}

		}
	}

	// Add the Nonogram cells and row clues
	for y := 0; y < rows; y++ {
		if rowsNumbers[y] > 10 {
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
	return grid
}

func isValidMove(row, col, gridSize int, cellsOfNonogram []*NonogramCell, rowsNumbers, colsNumbers []int) bool {
	// Check if the row can still satisfy the row clue after this move
	if !canSatisfyClue(extractLine(cellsOfNonogram, row, gridSize, true), rowsNumbers[row]) {
		return false
	}

	// Check if the column can still satisfy the column clue after this move
	if !canSatisfyClue(extractLine(cellsOfNonogram, col, gridSize, false), colsNumbers[col]) {
		return false
	}

	return true
}

// canSatisfyClue checks if a line (row or column) can still potentially satisfy the clues given.
func canSatisfyClue(line []*NonogramCell, clue int) bool {
	// Convert clue to a slice of individual numbers if greater than 10.
	clues := decodeClues(clue)

	// Try to place the clues in every possible way in the line and check if any configuration is valid.
	for i := 0; i <= len(line)-len(clues); i++ {
		if isValidConfiguration(line, clues, i) {
			return true
		}
	}
	return false
}

// isValidConfiguration checks if the clues can be placed starting at a specific index in the line.
func isValidConfiguration(line []*NonogramCell, clues []int, start int) bool {
	clueIndex := 0
	for i := start; i < len(line); i++ {
		if clues[clueIndex] == 0 {
			// Move to the next clue and skip a space for separation.
			clueIndex++
			if clueIndex >= len(clues) {
				break // All clues are placed.
			}
			i++ // Skip a space for separation.
		} else if line[i].state == StateEmpty || line[i].state == StateFilled {
			// A black cell is found or can be placed here.
			clues[clueIndex]--
		} else {
			// A white cell is found, which is not allowed in a block.
			return false
		}
	}

	// Check if all clues are placed.
	for _, remaining := range clues[clueIndex:] {
		if remaining != 0 {
			return false // Not all clues are placed.
		}
	}

	return true // The configuration is valid.
}

func decodeClues(clue int) []int {
	if clue > 10 {
		return []int{clue / 10, clue % 10}
	}
	return []int{clue}
}
func extractLine(cells []*NonogramCell, index, gridSize int, isRow bool) []*NonogramCell {
	line := make([]*NonogramCell, gridSize)
	for i := 0; i < gridSize; i++ {
		if isRow {
			line[i] = cells[index*gridSize+i]
		} else {
			line[i] = cells[i*gridSize+index]
		}
	}
	return line
}

func checkWinConditionNonogram(gridSize int, cellsOfNonogram []*NonogramCell, rowsNumbers, colsNumbers []int) bool {
	for rowIndex := 0; rowIndex < gridSize; rowIndex++ {
		if !lineSatisfiesClue(extractLine(cellsOfNonogram, rowIndex, gridSize, true), rowsNumbers[rowIndex]) {
			return false
		}
	}
	for colIndex := 0; colIndex < gridSize; colIndex++ {
		if !lineSatisfiesClue(extractLine(cellsOfNonogram, colIndex, gridSize, false), colsNumbers[colIndex]) {
			return false
		}
	}
	return true
}

func lineSatisfiesClue(line []*NonogramCell, clue int) bool {
	clues := decodeClues(clue)
	currentClueIndex := 0
	count := 0
	for i, cell := range line {
		if cell.state == StateFilled {
			count++
			if currentClueIndex >= len(clues) || count > clues[currentClueIndex] {
				// Too many filled cells or does not match the clue.
				return false
			}
		} else {
			if count > 0 {
				if count != clues[currentClueIndex] { // The segment of filled cells does not match the size of the current clue.
					return false
				}
				// Move to the next clue.
				currentClueIndex++
				count = 0 // Reset count for the next segment.
			}
		}

		// If we reach the end of the line, we check the last segment.
		if i == len(line)-1 && count > 0 {
			if count != clues[currentClueIndex] {
				return false
			}
			currentClueIndex++
		}
	}

	// Check if all clues have been accounted for. There should be no remaining clues.
	return currentClueIndex == len(clues)
}
