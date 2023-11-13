package PaintFloor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"math/rand"
	"time"
)

const (
	gridSize = 10
	cellSize = 30
)

var (
	grid             [gridSize][gridSize]*canvas.Rectangle
	playerX, playerY int
	obstacleColor    = color.RGBA{R: 128, G: 128, B: 128, A: 255} // Gray color for obstacles
	playerColor      = color.RGBA{R: 255, G: 0, B: 0, A: 255}     // Red color for the player
)

func MakeGame() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Game")

	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Initialize the grid and randomly place obstacles
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			grid[x][y] = canvas.NewRectangle(color.White)
			grid[x][y].SetMinSize(fyne.NewSize(cellSize, cellSize))
			// Randomly make some cells as obstacles
			if rand.Float32() < 0.2 { // 20% chance of being an obstacle
				grid[x][y].FillColor = obstacleColor
				grid[x][y].Refresh()
			}
		}
	}

	// Create a container with the grid layout
	gridLayout := container.NewGridWithColumns(gridSize)
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			gridLayout.Add(grid[x][y])
		}
	}

	// Set the initial player position
	playerX, playerY = 0, 0
	paintCell(playerX, playerY) // Initial paint

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
	myWindow.Resize(fyne.NewSize(gridSize*cellSize, gridSize*cellSize))
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
		if newX >= 0 && newX < gridSize && newY >= 0 && newY < gridSize && !isObstacle(newX, newY) {
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
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if grid[x][y].FillColor != color.Black && grid[x][y].FillColor != obstacleColor {
				return false
			}
		}
	}
	return true
}
