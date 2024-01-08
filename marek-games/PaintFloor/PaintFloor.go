package PaintFloor

import (
	"bufio"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	cellSize = 30
)

var (
	grid             [][]*canvas.Rectangle
	gridHeight       int
	gridWidth        int
	playerX, playerY int
	obstacleColor    = color.RGBA{R: 128, G: 128, B: 128, A: 255} // Gray color for obstacles
	playerColor      = color.RGBA{R: 255, G: 0, B: 0, A: 255}     // Red color for the player
)
var levelCompleteMenu *widget.PopUp
var levelComplete = false

func MakeGame(levelFilename string) fyne.Window {
	myApp := app.NewWithID("PaintFloor")
	myWindow := myApp.NewWindow("Grid Game")

	err := getLevelSize(levelFilename)
	if err != nil {
		return nil
	}
	gridLayout := container.NewGridWithColumns(gridWidth)
	currentLevel := 1 // Initialize current level

	// Load level data from the file
	if err := loadLevelFromFile(levelFilename, gridLayout); err != nil {
		// Handle the error, maybe load a default level
	}

	// Create a container with the grid layout
	//gridLayout := container.NewGridWithColumns(gridWidth) // Use gridWidth here for the number of columns
	//for x := 0; x < gridHeight; x++ {                     // Iterate over rows first (height)
	//	for y := 0; y < gridWidth; y++ { // Then iterate over columns (width) within each row
	//		gridLayout.Add(grid[x][y])
	//	}
	//}

	// Function to create and show the level complete menu
	levelCompleteMenu = createLevelCompleteMenu(myWindow, &currentLevel, gridLayout)
	levelCompleteMenu.Hide()
	myWindow.Canvas().Refresh(myWindow.Content())

	// Handle key inputs for movement
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {

		if levelComplete {
			return
		}
		gridLayout.Refresh()
		switch key.Name {
		case fyne.KeyUp, fyne.KeyDown, fyne.KeyLeft, fyne.KeyRight:
			movePlayer(string(key.Name))
			if checkLevelComplete() {
				levelComplete = true
				print("level complete")
				// Initialize the level completion menu but don't show it yet
				levelCompleteMenu.Show()
				myWindow.Canvas().Refresh(myWindow.Content())
			}
		}
	})

	myWindow.SetContent(gridLayout)
	myWindow.Resize(fyne.NewSize(float32(gridWidth*cellSize+350), float32(gridHeight*cellSize+350)))
	return myWindow
}

func movePlayer(direction string) {
	if levelComplete {
		return // Stop player movement if level is complete
	}
	dx, dy := 0, 0
	switch direction {
	case string(fyne.KeyUp):
		dx = -1
	case string(fyne.KeyDown):
		dx = 1
	case string(fyne.KeyLeft):
		dy = -1
	case string(fyne.KeyRight):
		dy = 1
	}

	// Continue moving in the direction with a delay for smoother animation
	for {
		newX, newY := playerX+dx, playerY+dy
		if newX >= 0 && newX < gridHeight && newY >= 0 && newY < gridWidth && !isObstacle(newX, newY) {
			// Paint the current cell
			paintCell(playerX, playerY)

			// Update to new position
			playerX, playerY = newX, newY

			// Display the player at new position
			grid[playerX][playerY].FillColor = playerColor
			grid[playerX][playerY].Refresh()

			// Delay for smoother movement
			time.Sleep(50 * time.Millisecond)
		} else {
			break // Stop moving if the edge is reached
		}
	}
}

func paintCell(x, y int) {
	grid[x][y].FillColor = color.Black
	grid[x][y].Refresh()
}

func isObstacle(x, y int) bool {
	return grid[x][y].FillColor == obstacleColor
}

func checkLevelComplete() bool {
	for x := 0; x < gridHeight; x++ {
		for y := 0; y < gridWidth; y++ {
			if grid[x][y].FillColor != color.Black && grid[x][y].FillColor != obstacleColor && grid[x][y].FillColor != playerColor {
				return false
			}
		}
	}
	return true
}

// return gridWidth,gridth height,error
// first number in the first line is gridHeight
// secoond number is gridWidth
func getLevelSize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Read the first line to get the grid size
	if !scanner.Scan() {
		fmt.Println(err)
		return errors.New("failed to read grid size")
	}
	sizeStr := scanner.Text()
	sizeParts := strings.Split(sizeStr, "x")
	if len(sizeParts) != 2 {
		fmt.Println(err)
		return errors.New("invalid grid size format")
	}
	gridHeight, err = strconv.Atoi(sizeParts[0])
	fmt.Println("Number of lines: " + strconv.Itoa(gridHeight))
	if err != nil {
		fmt.Println(err)
		return err
	}
	gridWidth, err = strconv.Atoi(sizeParts[1])
	fmt.Println("Number of collumns: " + strconv.Itoa(gridWidth))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func loadLevelFromFile(filename string, gridLayout *fyne.Container) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return errors.New("failed to skip the grid size line")
	}

	// Initialize the grid with the read size
	gridLayout.Objects = nil // Clear existing objects in gridLayout
	grid = make([][]*canvas.Rectangle, gridHeight)
	for x := range grid {
		grid[x] = make([]*canvas.Rectangle, gridWidth)
		for y := range grid[x] {
			grid[x][y] = canvas.NewRectangle(color.White)
			grid[x][y].SetMinSize(fyne.NewSize(cellSize, cellSize))
			grid[x][y].Refresh()
		}
	}

	// Process the rest of the file to set up the grid
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != gridWidth {
			return errors.New("line length does not match gridWidth")
		}
		for y, char := range line {
			switch char {
			case '1': // Obstacle
				grid[x][y].FillColor = obstacleColor
			case 'S': // Starting position
				playerX, playerY = x, y
				grid[x][y].FillColor = playerColor
			}
		}
		x++
		if x > gridHeight {
			return errors.New("too many lines for declared gridHeight")
		}
	}

	if x != gridHeight {
		return errors.New("not enough grid data for declared size")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Clear existing grid layout

	for x := 0; x < gridHeight; x++ { // Iterate over rows first (height)
		for y := 0; y < gridWidth; y++ { // Then iterate over columns (width) within each row
			gridLayout.Add(grid[x][y])
		}
	}

	gridLayout.Refresh() // Refresh the layout to update the UI
	return nil
}

func createLevelCompleteMenu(myWindow fyne.Window, currentLevel *int, gridLayout *fyne.Container) *widget.PopUp {

	levelCompleteMenu = widget.NewModalPopUp(container.NewVBox(
		widget.NewLabel("Level Complete!"),
		widget.NewButton("Next Level", func() {
			*currentLevel++
			err := getLevelSize("marek-games/PaintFloor/levels/lvl" + strconv.Itoa(*currentLevel) + ".txt")

			gridLayout.Objects = nil
			gridLayout = container.NewGridWithColumns(gridWidth)

			err = loadLevelFromFile("marek-games/PaintFloor/levels/lvl"+strconv.Itoa(*currentLevel)+".txt", gridLayout)
			if err != nil {
				return
			} // Load the next level
			fmt.Println("leading new level:" + strconv.Itoa(*currentLevel))

			levelComplete = false
			myWindow.SetContent(gridLayout)
			levelCompleteMenu.Hide()
			myWindow.Canvas().Refresh(myWindow.Content())
		}),
		// Add other buttons if necessary
	), myWindow.Canvas())
	myWindow.SetContent(gridLayout)
	myWindow.Resize(fyne.NewSize(float32(gridWidth*cellSize+700), float32(gridHeight*cellSize+700)))
	return levelCompleteMenu
}
