package gamesWindows

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"time"
)

const (
	cellSize = 30
)

type LevelPaintFloor struct {
	Dimensions []int `firestore:"dimensions"`
	Map        []int `firestore:"map"`
}
type PaintFloorScreen struct {
	window          fyne.Window
	app             fyne.App
	mainMenuContent fyne.CanvasObject
	level           string
}

var PsWindow *PaintFloorScreen
var LevelInProggressPaint string
var ContentPaint *fyne.Container

func NewPaintFloorScreen(window fyne.Window, app fyne.App, mainMenuContent fyne.CanvasObject, level string) *PaintFloorScreen {
	mainApp = app
	return &PaintFloorScreen{window: window, app: app, mainMenuContent: mainMenuContent, level: level}
}

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

func (ps *PaintFloorScreen) Render() {

	emptyObject := canvas.NewRectangle(nil)
	containeris := container.NewVBox(emptyObject)
	ps.window.SetContent(containeris)
	levelComplete = false
	wind = ps.window
	mainContent = ps.mainMenuContent

	app := app.New()

	menuWindow := app.NewWindow("Logitec App")

	menuWindow.SetOnClosed(func() {
		app.Quit()
	})

	cont := MakeGamePaintFloor(ps.level)
	cont.Canvas()

	ps.window.SetContent(ContentPaint)
	SetUpPaintFloorWindow(ps.window, ContentPaint, ps)

	PsWindow = ps
	ps.window.CenterOnScreen()
}

func MakeGamePaintFloor(levelId string) fyne.Window {
	levelComplete = false
	ContentPaint = nil
	myApp := app.NewWithID("PaintFloor")
	myWindow := myApp.NewWindow("Griddy Game")

	currentLevel, _ := StringToInt(levelId)
	LevelInProggressPaint = levelId
	// Initialize current level
	mapData, err := getLevelData(&currentLevel)

	if err != nil {
		return nil
	}
	gridLayout := container.NewGridWithColumns(gridWidth)
	// Load level data from the file
	erroris, gridLayout := loadLevelFromData(mapData, gridLayout)
	if erroris != nil {
		// Handle the error, maybe load a default level
	}

	// Function to create and show the level complete menu
	levelCompleteMenu = createLevelCompleteMenu(myWindow, &currentLevel, gridLayout)
	levelCompleteMenu.Hide()

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
				print("level complete from makegame function")
				// Initialize the level completion menu but don't show it yet
				levelCompleteMenu.Show()
				myWindow.Canvas().Refresh(myWindow.Content())

			}
		}
	})
	ContentPaint = gridLayout

	myWindow.SetContent(gridLayout)
	myWindow.Resize(fyne.NewSize(float32(gridWidth*cellSize+350), float32(gridHeight*cellSize+350)))
	myWindow.Canvas().Refresh(myWindow.Content())
	return myWindow
}

func SetUpPaintFloorWindow(myWindow fyne.Window, gridLayout *fyne.Container, ps *PaintFloorScreen) fyne.Window {
	print("game created from setup function")

	levelComplete = false
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if ps.window.Content() == ps.mainMenuContent {
			return
		}
		//if levelComplete {
		//	MakeWinWindow(ps.app, ps.window, ps.mainMenuContent, "paint")
		//}
		gridLayout.Refresh()
		switch key.Name {
		case fyne.KeyUp, fyne.KeyDown, fyne.KeyLeft, fyne.KeyRight:
			movePlayer(string(key.Name))
			if checkLevelComplete() {
				levelComplete = true
				print("level complete from setup function")
				// Initialize the level completion menu but don't show it yet
				levelCompleteMenu.Show()
				MakeWinWindow(ps.app, ps.window, ps.mainMenuContent, "paint")
				myWindow.Canvas().Refresh(myWindow.Content())
			}
		}
	})

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
func getLevelData(currentLevel *int) ([]int, error) {
	levelID := "lvl" + strconv.Itoa(*currentLevel)
	dimensions, mapData, err := LoadLevelData(levelID)
	if err != nil {
		fmt.Println("Error loading level data:", err)
		return nil, err
	}

	gridHeight = dimensions[0]
	gridWidth = dimensions[1]

	return mapData, nil
}

func loadLevelFromData(mapData []int, gridLayout *fyne.Container) (error, *fyne.Container) {
	backButton := widget.NewButton("Back", func() {
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
	if len(mapData) != gridHeight*gridWidth {
		return errors.New("map data size does not match grid dimensions"), finalContainer
	}

	// Clear existing grid layout
	grid = make([][]*canvas.Rectangle, gridHeight)

	for x := 0; x < gridHeight; x++ {
		grid[x] = make([]*canvas.Rectangle, gridWidth)
		for y := 0; y < gridWidth; y++ {
			tileValue := mapData[x*gridWidth+y]
			switch tileValue {
			case 1: // Obstacle
				grid[x][y] = canvas.NewRectangle(obstacleColor)
			case 2: // Player starting position
				playerX, playerY = x, y
				grid[x][y] = canvas.NewRectangle(playerColor)
			default:
				grid[x][y] = canvas.NewRectangle(color.White)
			}

			grid[x][y].SetMinSize(fyne.NewSize(cellSize, cellSize))
			grid[x][y].Refresh()
			gridLayout.Add(grid[x][y])

		}
	}

	containerToReturn := container.NewStack(
		gridLayout,
		finalContainer,
	)

	gridLayout = containerToReturn
	gridLayout.Refresh() // Refresh the layout to update the UI
	return nil, gridLayout
}

func createLevelCompleteMenu(myWindow fyne.Window, currentLevel *int, gridLayout *fyne.Container) *widget.PopUp {
	fmt.Println("Creating level complete menu")

	levelCompleteMenu = widget.NewModalPopUp(container.NewVBox(
		widget.NewLabel("Level Complete!"),
		widget.NewButton("Next Level", func() {
			*currentLevel++
			mapData, err := getLevelData(currentLevel)
			if err != nil {
				fmt.Println("Error getting level data:", err)
				return // Handle the error appropriately
			}

			gridLayout.Objects = nil
			gridLayout = container.NewGridWithColumns(gridWidth)

			// Load the next level
			err, _ = loadLevelFromData(mapData, gridLayout)
			if err != nil {
				fmt.Println("Error loading level from data:", err)
				return // Handle the error appropriately
			}

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

func StringToInt(str string) (int, error) {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
