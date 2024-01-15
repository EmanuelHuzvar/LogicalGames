package gamesWindows

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
}

func NewGame2048Screen(window fyne.Window, app fyne.App, mainMenuContent fyne.CanvasObject) *Game2048Screen {
	mainApp = app
	return &Game2048Screen{window: window, app: app, mainMenuContent: mainMenuContent}
}

func (g48 *Game2048Screen) Render() {
	wind = g48.window
	mainContent = g48.mainMenuContent

	app := app.New()

	menuWindow := app.NewWindow("Logitec App")

	menuWindow.SetOnClosed(func() {
		app.Quit()
	})

	content := MakeGameGame2048().Content()
	g48.window.SetContent(content)
	g48.window.CenterOnScreen()

}

func MakeGameGame2048() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("2048")

	gameState := NewGameState(4) // Assuming a 4x4 grid for 2048
	addRandomTile(gameState)

	renderGrid(myWindow, gameState)

	return myWindow
}

func NewGameState(size int) *GameState {
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
		state.Grid[position[0]][position[1]].Value = 64
	}
}

func renderGrid(window fyne.Window, state *GameState) {
	gridLayout := container.NewGridWithColumns(len(state.Grid))
	for i := range state.Grid {
		for j := range state.Grid[i] {
			tile := state.Grid[i][j]
			rect := canvas.NewRectangle(tileColor(tile.Value))
			rect.SetMinSize(fyne.NewSize(75, 75)) // Set the size of the tile

			label := canvas.NewText(formatTileValue(tile.Value), color.Black)
			label.TextStyle.Bold = true
			label.Alignment = fyne.TextAlignCenter

			overlay := container.NewStack(rect, label)
			gridLayout.Add(overlay)
		}
	}
	window.SetContent(gridLayout)
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
	// Ensure that the adjusted value is between 0 and 255
	adjustedValue := (value * 7) % 255
	// Keep red at full intensity and vary blue and green
	return color.RGBA{
		R: 255,                        // Red at full intensity
		G: 255 - uint8(adjustedValue), // Varying green
		B: 255 - uint8(adjustedValue), // Varying blue
		A: 255,
	}
}

func shadeOfGreen(value int) color.Color {
	// Ensure that the adjusted value is between 0 and 255
	adjustedValue := (value * 7) % 255

	// Keeping red at full intensity and varying green slightly
	// to give different shades of yellow.
	// The blue component is kept low to maintain the yellow color.
	return color.RGBA{
		R: uint8(255 - adjustedValue),
		G: 255,                        // Reducing green creates different shades
		B: uint8(255 - adjustedValue), // Keeping blue low
		A: 255,
	}
}

func shadeOfBlue(value int) color.Color {
	// Ensure that the adjusted value is between 0 and 255
	adjustedValue := (value * 7) % 255

	// Keep blue at full intensity and vary red and green
	return color.RGBA{
		R: uint8(255 - adjustedValue), // Varying red
		G: uint8(255 - adjustedValue), // Varying green
		B: 255,                        // Blue at full intensity
		A: 255,
	}
}

func setUpKeyboardListener(window fyne.Window, gameState *GameState) {
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyUp:
			//moveTilesUp(gameState)
			addRandomTile(gameState)
			renderGrid(window, gameState)

		case fyne.KeyRight:
			//moveTilesRight(gameState)
			addRandomTile(gameState)
			renderGrid(window, gameState)

		case fyne.KeyDown:
			//moveTilesDown(gameState)
			addRandomTile(gameState)
			renderGrid(window, gameState)

		case fyne.KeyLeft:
			//moveTilesLeft(gameState)
			addRandomTile(gameState)
			renderGrid(window, gameState)
		}

	})
}
