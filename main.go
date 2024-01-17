package main

import (
	"ProjectMarekEmanuel/mainMenu"
)

func main() {
	//marek games
	//PaintFloor.MakeGame().ShowAndRun()
	//Game2048.MakeGame().ShowAndRun()

	//fields := [][]string{
	//	{"red", "green", "blue", "yellow"},
	//	{"cyan", "purple", "magenta", "orange"},
	//	{"green", "red", "yellow", "blue"},
	//	{"purple", "cyan", "orange", "magenta"},
	//	{"blue", "yellow", "red", "green"},
	//	{"magenta", "orange", "cyan", "purple"},
	//	{"yellow", "blue", "green", "red"},
	//	{"orange", "magenta", "purple", "cyan"},
	//}
	//levelID := "5"
	//cols := []int{3, 11, 3}
	//rows := []int{3, 0, 3}
	//winCil := []string{"101", "101", "101"}
	//err := gamesWindows.AddLevelForNonogram(levelID, cols, rows, winCil)
	//if err != nil {
	//	fmt.Printf("Error adding level for Nonogram: %s\n", err)
	//} else {
	//	fmt.Println("Level added successfully for Nonogram.")
	//}
	//gamesWindows.AddLevelForBubble("4", fields)

	mainMenu.MakeMenu().ShowAndRun()
	//app := app.New()
	//gamesWindows.MakeTryWindow(app)

}
