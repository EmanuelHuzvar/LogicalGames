package PaintFloor

import (
	"bufio"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

func MakeGame(levelFilename string) fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Game")

	// Load level data from the file
	if err := loadLevelFromFile(levelFilename); err != nil {
		// Handle the error, maybe load a default level
	}

	// Create a container with the grid layout
	gridLayout := container.NewGridWithColumns(gridWidth)
	for y := 0; y < gridWidth; y++ {
		for x := 0; x < gridHeight; x++ {
			gridLayout.Add(grid[x][y])
		}
	}

	// Handle key inputs for movement
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case fyne.KeyUp, fyne.KeyDown, fyne.KeyLeft, fyne.KeyRight:
			movePlayer(string(key.Name))
			if checkLevelComplete() {
				// Level is complete, perform necessary action
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Grid Game",
					Content: "Level complete!",
				})
			}
		}
	})

	myWindow.SetContent(gridLayout)
	myWindow.Resize(fyne.NewSize(float32(gridWidth*cellSize), float32(gridHeight*cellSize)))
	return myWindow
}

func movePlayer(direction string) {
	dx, dy := 0, 0
	switch direction {
	case string(fyne.KeyUp):
		dy = -1
	case string(fyne.KeyDown):
		dy = 1
	case string(fyne.KeyLeft):
		dx = -1
	case string(fyne.KeyRight):
		dx = 1
	}

	// Continue moving in the direction with a delay for smoother animation
	for {
		newX, newY := playerX+dx, playerY+dy
		if newX >= 0 && newX < gridWidth && newY >= 0 && newY < gridHeight && !isObstacle(newX, newY) {
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
		for y := 0; y < gridHeight; y++ {
			if grid[x][y].FillColor != color.Black && grid[x][y].FillColor != obstacleColor {
				return false
			}
		}
	}
	return true
}

func loadLevelFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line to get the grid size
	if !scanner.Scan() {
		return errors.New("failed to read grid size")
	}
	sizeStr := scanner.Text()
	sizeParts := strings.Split(sizeStr, "x")
	if len(sizeParts) != 2 {
		return errors.New("invalid grid size format")
	}
	gridHeight, err = strconv.Atoi(sizeParts[0])
	if err != nil {
		return err
	}
	gridWidth, err = strconv.Atoi(sizeParts[1])
	if err != nil {
		return err
	}

	// Initialize the grid with the read size
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
				// Add more cases as needed
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

	return nil
}
