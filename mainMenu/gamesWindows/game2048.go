package gamesWindows

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type Tile struct {
	Value  int
	Merged bool
}

type GameState struct {
	Grid     [][]*Tile
	Score    int
	GameOver bool
}
type Game2048Screen struct {
	window          fyne.Window
	app             fyne.App
	mainMenuContent fyne.CanvasObject
	scoreLabel      *canvas.Text
	gridLayout      *fyne.Container
	mainMenuWindow  fyne.CanvasObject
	backButton      fyne.Widget
}

var Game2048WindowInProggress *Game2048Screen

var gameStateInProgress *GameState

func NewGame2048Screen(window fyne.Window, app fyne.App, mainMenuContent fyne.CanvasObject, mainMenuWindow fyne.CanvasObject) *Game2048Screen {
	mainApp = app
	//window.SetFixedSize(true)
	//window.Resize(fyne.NewSize(800, 800))
	return &Game2048Screen{window: window, app: app, mainMenuContent: mainMenuContent, mainMenuWindow: mainMenuWindow}
}

func (g48 *Game2048Screen) Render() {
	wind = g48.window
	mainContent = g48.mainMenuContent

	app := app.New()

	menuWindow := app.NewWindow("Logitec App")

	menuWindow.SetOnClosed(func() {
		app.Quit()
	})
	MakeGameGame2048(g48)
	content := MakeGameGame2048(g48).Content()
	g48.window.SetContent(content)
	setUpKeyboardListener(g48.window, g48, gameStateInProgress)
	g48.window.CenterOnScreen()

}

func MakeGameGame2048(g48 *Game2048Screen) fyne.Window {
	BackBtnImgResource := fyne.NewStaticResource("back-arrow.png", backBtnImg)
	myApp := app.New()
	myWindow := myApp.NewWindow("2048")
	//Game2048WindowInProggress = nil

	g48.backButton = widget.NewButtonWithIcon("", BackBtnImgResource, func() {
		g48.window.SetContent(g48.mainMenuContent)
	})

	gameState := NewGameState(4) // Assuming a 4x4 grid for 2048
	addRandomTile(gameState)
	//test score label
	g48.scoreLabel = canvas.NewText(fmt.Sprintf("Score: %d", gameState.Score), color.Black)
	g48.scoreLabel.TextStyle.Bold = true

	g48.gridLayout = createGridContainer(gameState)

	g48.gridLayout.Resize(fyne.NewSize(1000, 1000))
	renderGrid(gameState, g48)

	//verticalLayout := container.NewVBox(g48.backButton, g48.scoreLabel, g48.gridLayout)

	//myWindow.SetContent(verticalLayout)

	Game2048WindowInProggress = g48
	vBoxContainer := setUpLayout(g48)
	myWindow.SetContent(vBoxContainer)
	return myWindow
}

func NewGameState(size int) *GameState {
	gameStateInProgress = nil
	state := &GameState{
		Grid:     make([][]*Tile, size),
		Score:    0,
		GameOver: false,
	}

	for i := range state.Grid {
		state.Grid[i] = make([]*Tile, size)
		for j := range state.Grid[i] {
			state.Grid[i][j] = &Tile{Value: 0}
		}
	}
	gameStateInProgress = state
	// Optionally, add two random tiles
	return state
}

func formatTileValue(value int) string {
	if value == 0 {
		return ""
	}
	return fmt.Sprintf("%d", value)
}

func addRandomTile(state *GameState) {
	rand.NewSource(time.Now().UnixNano())
	var emptyTiles [][2]int
	for i, row := range state.Grid {
		for j, tile := range row {
			if tile.Value == 0 {
				emptyTiles = append(emptyTiles, [2]int{i, j})
			}
		}
	}

	if len(emptyTiles) > 0 {
		randomIndex := rand.Intn(len(emptyTiles))
		position := emptyTiles[randomIndex]
		state.Grid[position[0]][position[1]].Value = 2
	}
}

func renderGrid(state *GameState, g48 *Game2048Screen) {

	gridLayout := container.NewGridWithColumns(len(state.Grid))
	tileSize := 150
	for i := range state.Grid {
		for j := range state.Grid[i] {
			tile := state.Grid[i][j]
			rect := canvas.NewRectangle(tileColor(tile.Value))
			rect.SetMinSize(fyne.NewSize(float32(tileSize), float32(tileSize))) // Set the size of the tile

			label := canvas.NewText(formatTileValue(tile.Value), color.Black)
			label.TextStyle.Bold = true
			label.TextSize = 50
			label.Alignment = fyne.TextAlignCenter

			overlay := container.NewStack(rect, label)

			gridLayout.Add(overlay)

		}
	}

	g48.gridLayout = gridLayout
	g48.gridLayout.Refresh()
	g48.scoreLabel.Text = fmt.Sprintf("Score: %d", state.Score)

	g48.gridLayout.Refresh()
	g48.scoreLabel.Refresh()
}

func tileColor(value int) color.Color {
	if value == 0 {
		return color.RGBA{R: 242, G: 243, B: 244, A: 255} // Color for empty tile
	}

	// Calculate the number of merges
	merges := math.Log2(float64(value)) / math.Log2(2) // log base 2 of value

	// Determine the color cycle based on the number of merges
	switch int(merges) / 5 { // Divide by 3 to change colors every three merges
	case 0:
		// Shades of red for the first cycle
		return shadeOfRed(value)
	case 1:
		// Shades of yellow for the second cycle
		return shadeOfGreen(value)
	case 2:
		// Shades of blue for the third cycle
		return shadeOfBlue(value)
	// Add more cases as needed
	default:
		return color.RGBA{R: 200, G: 200, B: 200, A: 255}
	}
}

func shadeOfRed(value int) color.Color {
	var greenBlueComponent uint8
	maxComponent := 200
	minComponent := 50
	componentRange := maxComponent - minComponent
	increments := float64(componentRange) / 4

	// Calculate the position of the value in its group of 4
	merges := int(math.Log2(float64(value))) // log base 2 of value
	positionInGroup := (merges - 1) % 4      // Position in the group of 4

	greenBlueComponent = uint8(maxComponent - int(increments)*positionInGroup)

	return color.RGBA{
		R: 255, // Red at full intensity
		G: greenBlueComponent,
		B: greenBlueComponent,
		A: 255,
	}
}

func shadeOfGreen(value int) color.Color {
	maxComponent := 200
	minComponent := 50
	componentRange := maxComponent - minComponent
	increments := float64(componentRange) / 5

	merges := int(math.Log2(float64(value)))
	positionInGroup := merges % 5
	greenBlueComponent := uint8(maxComponent - int(increments)*positionInGroup)

	return color.RGBA{
		R: greenBlueComponent,
		G: 255,
		B: greenBlueComponent,
		A: 255,
	}
}

func shadeOfBlue(value int) color.Color {
	maxComponent := 200
	minComponent := 50
	componentRange := maxComponent - minComponent
	increments := float64(componentRange) / 5

	merges := int(math.Log2(float64(value)))
	positionInGroup := merges % 5
	redGreenComponent := uint8(maxComponent - int(increments)*positionInGroup)

	return color.RGBA{
		R: redGreenComponent,
		G: redGreenComponent,
		B: 255,
		A: 255,
	}
}

func setUpKeyboardListener(window fyne.Window, g48 *Game2048Screen, gameState *GameState) {
	g48.window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyUp:
			if !canMove(gameState) {
				//MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}
			if canMoveUp(gameState) {
				fmt.Println()
				moveTilesUp(gameState)
				addRandomTile(gameState)
				renderGrid(gameState, g48)
				g48.window.SetContent(container.NewVBox(setUpLayout(g48)))
				fmt.Printf("Current Score: %d\n", gameState.Score)
			}
			if !canMove(gameState) {
				MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}

		case fyne.KeyRight:
			if !canMove(gameState) {
				MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}
			if canMoveRight(gameState) {
				moveTilesRight(gameState)
				addRandomTile(gameState)
				renderGrid(gameState, g48)
				g48.window.SetContent(container.NewVBox(setUpLayout(g48)))
				fmt.Printf("Current Score: %d\n", gameState.Score)
			}
			if !canMove(gameState) {
				MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}
		case fyne.KeyDown:
			if !canMove(gameState) {
				MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}
			if canMoveDown(gameState) {
				moveTilesDown(gameState)
				addRandomTile(gameState)
				renderGrid(gameState, g48)
				g48.window.SetContent(container.NewVBox(setUpLayout(g48)))
				fmt.Printf("Current Score: %d\n", gameState.Score)
			}
			if !canMove(gameState) {
				MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}
		case fyne.KeyLeft:
			if !canMove(gameState) {
				MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}
			if canMoveLeft(gameState) {
				moveTilesLeft(gameState)
				addRandomTile(gameState)
				renderGrid(gameState, g48)
				g48.window.SetContent(container.NewVBox(setUpLayout(g48)))
				fmt.Printf("Current Score: %d\n", gameState.Score)
			}
			if !canMove(gameState) {
				MakeLoseWindow(g48.app, g48.window, g48.mainMenuContent, "2048")
			}

		}
	})
}

func moveTilesUp(gameState *GameState) {
	for col := 0; col < len(gameState.Grid[0]); col++ {
		// First, compress the column by moving non-zero tiles up
		compressColumnUp(gameState, col)

		// Merge tiles
		for row := 0; row < len(gameState.Grid)-1; row++ {
			if gameState.Grid[row][col].Value != 0 && gameState.Grid[row][col].Value == gameState.Grid[row+1][col].Value && !gameState.Grid[row][col].Merged && !gameState.Grid[row+1][col].Merged {
				mergedValue := gameState.Grid[row][col].Value * 2
				gameState.Grid[row][col].Value *= 2
				gameState.Grid[row+1][col].Value = 0
				gameState.Grid[row][col].Merged = true
				// Update score based on merged value
				gameState.Score += mergedValue * 10
			}
		}

		// Compress again after merging
		compressColumnUp(gameState, col)
	}
}

func compressColumnUp(gameState *GameState, col int) {
	idx := 0
	for row := 0; row < len(gameState.Grid); row++ {
		if gameState.Grid[row][col].Value != 0 {
			gameState.Grid[idx][col].Value = gameState.Grid[row][col].Value
			gameState.Grid[idx][col].Merged = false
			if idx != row {
				gameState.Grid[row][col].Value = 0
			}
			idx++
		}
	}
}

func moveTilesDown(gameState *GameState) {
	for col := 0; col < len(gameState.Grid[0]); col++ {
		// First, compress the column by moving non-zero tiles down
		compressColumnDown(gameState, col)

		// Merge tiles from bottom up
		for row := len(gameState.Grid) - 1; row > 0; row-- {
			if gameState.Grid[row][col].Value != 0 && gameState.Grid[row][col].Value == gameState.Grid[row-1][col].Value && !gameState.Grid[row][col].Merged && !gameState.Grid[row-1][col].Merged {
				mergedValue := gameState.Grid[row][col].Value * 2
				gameState.Grid[row][col].Value *= 2
				gameState.Grid[row-1][col].Value = 0
				gameState.Grid[row][col].Merged = true
				gameState.Score += mergedValue * 10
			}
		}

		// Compress again after merging
		compressColumnDown(gameState, col)
	}
}

func compressColumnDown(gameState *GameState, col int) {
	idx := len(gameState.Grid) - 1
	for row := len(gameState.Grid) - 1; row >= 0; row-- {
		if gameState.Grid[row][col].Value != 0 {
			gameState.Grid[idx][col].Value = gameState.Grid[row][col].Value
			gameState.Grid[idx][col].Merged = false
			if idx != row {
				gameState.Grid[row][col].Value = 0
			}
			idx--
		}
	}
}

func moveTilesLeft(gameState *GameState) {
	for row := 0; row < len(gameState.Grid); row++ {
		compressRowLeft(gameState, row)

		for col := 0; col < len(gameState.Grid[row])-1; col++ {
			if gameState.Grid[row][col].Value != 0 && gameState.Grid[row][col].Value == gameState.Grid[row][col+1].Value && !gameState.Grid[row][col].Merged && !gameState.Grid[row][col+1].Merged {
				mergedValue := gameState.Grid[row][col].Value * 2
				gameState.Grid[row][col].Value *= 2
				gameState.Grid[row][col+1].Value = 0
				gameState.Grid[row][col].Merged = true
				gameState.Score += mergedValue * 10
			}
		}

		compressRowLeft(gameState, row)
	}
}

func compressRowLeft(gameState *GameState, row int) {
	idx := 0
	for col := 0; col < len(gameState.Grid[row]); col++ {
		if gameState.Grid[row][col].Value != 0 {
			gameState.Grid[row][idx].Value = gameState.Grid[row][col].Value
			gameState.Grid[row][idx].Merged = false
			if idx != col {
				gameState.Grid[row][col].Value = 0
			}
			idx++
		}
	}
}

func moveTilesRight(gameState *GameState) {
	for row := 0; row < len(gameState.Grid); row++ {
		compressRowRight(gameState, row)

		for col := len(gameState.Grid[row]) - 1; col > 0; col-- {
			if gameState.Grid[row][col].Value != 0 && gameState.Grid[row][col].Value == gameState.Grid[row][col-1].Value && !gameState.Grid[row][col].Merged && !gameState.Grid[row][col-1].Merged {
				mergedValue := gameState.Grid[row][col].Value * 2
				gameState.Grid[row][col].Value *= 2
				gameState.Grid[row][col-1].Value = 0
				gameState.Grid[row][col].Merged = true
				gameState.Score += mergedValue * 10
			}
		}

		compressRowRight(gameState, row)
	}

}

func compressRowRight(gameState *GameState, row int) {
	idx := len(gameState.Grid[row]) - 1
	for col := len(gameState.Grid[row]) - 1; col >= 0; col-- {
		if gameState.Grid[row][col].Value != 0 {
			gameState.Grid[row][idx].Value = gameState.Grid[row][col].Value
			gameState.Grid[row][idx].Merged = false
			if idx != col {
				gameState.Grid[row][col].Value = 0
			}
			idx--
		}
	}
}

func canMoveLeft(gameState *GameState) bool {
	for row := 0; row < len(gameState.Grid); row++ {
		for col := 1; col < len(gameState.Grid[row]); col++ {
			if gameState.Grid[row][col].Value != 0 && (gameState.Grid[row][col-1].Value == 0 || gameState.Grid[row][col-1].Value == gameState.Grid[row][col].Value) {
				fmt.Println("Can move left: true")
				return true
			}
		}
	}
	fmt.Println("Can move left: false")
	return false
}

func canMoveRight(gameState *GameState) bool {
	for row := 0; row < len(gameState.Grid); row++ {
		for col := 0; col < len(gameState.Grid[row])-1; col++ {
			if gameState.Grid[row][col].Value != 0 && (gameState.Grid[row][col+1].Value == 0 || gameState.Grid[row][col+1].Value == gameState.Grid[row][col].Value) {
				fmt.Println("Can move right: true")
				return true
			}
		}
	}
	fmt.Println("Can move right: false")
	return false
}

func canMoveUp(gameState *GameState) bool {
	for col := 0; col < len(gameState.Grid[0]); col++ {
		for row := 1; row < len(gameState.Grid); row++ {
			if gameState.Grid[row][col].Value != 0 && (gameState.Grid[row-1][col].Value == 0 || gameState.Grid[row-1][col].Value == gameState.Grid[row][col].Value) {
				fmt.Println("Can move up: true")
				return true
			}
		}
	}
	fmt.Println("Can move up: false")
	return false
}

func canMoveDown(gameState *GameState) bool {
	for col := 0; col < len(gameState.Grid[0]); col++ {
		for row := 0; row < len(gameState.Grid)-1; row++ {
			if gameState.Grid[row][col].Value != 0 && (gameState.Grid[row+1][col].Value == 0 || gameState.Grid[row+1][col].Value == gameState.Grid[row][col].Value) {
				fmt.Println("Can move down: true")
				return true
			}
		}
	}
	fmt.Println("Can move down: false")
	return false
}

func canMove(gameState *GameState) bool {
	if !canMoveUp(gameState) && !canMoveRight(gameState) && !canMoveDown(gameState) && !canMoveLeft(gameState) {
		fmt.Println("Can move: false")
		return false
	}
	fmt.Println("Can move: true")
	return true
}

//	func createGridContainer(state *GameState) *fyne.Container {
//		gridLayout := container.NewGridWithColumns(len(state.Grid)) // Assuming a square grid
//
//		for i := range state.Grid {
//			for j := range state.Grid[i] {
//				tile := state.Grid[i][j]
//				rect := canvas.NewRectangle(tileColor(tile.Value))
//				rect.SetMinSize(fyne.NewSize(75, 75)) // Set the size of the tile
//
//				label := canvas.NewText(formatTileValue(tile.Value), color.Black)
//				label.TextStyle.Bold = true
//				label.Alignment = fyne.TextAlignCenter
//
//				overlay := container.NewStack(rect, label)
//				gridLayout.Add(overlay)
//			}
//		}
//
//		return gridLayout
//	}
func createGridContainer(state *GameState) *fyne.Container {
	gridSize := len(state.Grid)
	tileSize := 100 // You need to define calculateTileSize

	gridLayout := container.NewGridWithColumns(gridSize)

	for i := range state.Grid {
		for j := range state.Grid[i] {
			tile := state.Grid[i][j]
			rect := canvas.NewRectangle(tileColor(tile.Value))
			rect.SetMinSize(fyne.NewSize(float32(tileSize), float32(tileSize)))

			label := canvas.NewText(formatTileValue(tile.Value), color.Black)
			label.TextStyle.Bold = true
			label.TextSize = 50
			label.Alignment = fyne.TextAlignCenter

			overlay := container.NewStack(rect, label)
			gridLayout.Add(overlay)
		}
	}

	return gridLayout
}

func setUpLayout(g48 *Game2048Screen) *fyne.Container {
	// Create an HBox layout for the top line
	topLine := container.NewHBox()

	// Add the back button to the left
	topLine.Add(g48.backButton)

	// Add a spacer to push the score label to the center
	topLine.Add(layout.NewSpacer())

	// Add the score label in the middle
	topLine.Add(g48.scoreLabel)

	// Add another spacer to keep the score label in the center
	topLine.Add(layout.NewSpacer())

	// Create the main VBox layout including the top line and the grid layout
	mainLayout := container.NewVBox(topLine, g48.gridLayout)

	// Set the window's content to this new layout
	return mainLayout
}
