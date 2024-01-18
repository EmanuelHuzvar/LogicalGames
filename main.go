package main

import "ProjectMarekEmanuel/mainMenu"

func main() {
	//marek games
	//PaintFloor.MakeGame().ShowAndRun()
	//Game2048.MakeGame().ShowAndRun()

	//fields := [][]string{
	//	{"green", "yellow", "blue", "red"},
	//	{"purple", "orange", "magenta", "cyan"},
	//	{"red", "green", "yellow", "blue"},
	//	{"cyan", "magenta", "purple", "orange"},
	//	{"yellow", "blue", "red", "green"},
	//	{"orange", "purple", "cyan", "magenta"},
	//	{"blue", "red", "green", "yellow"},
	//	{"magenta", "cyan", "orange", "purple"},
	//}
	//
	//gamesWindows.AddLevelForBubble("15", fields)

	//cols := []int{3, 11, 3}
	//rows := []int{3, 0, 3}
	//winCil := []string{"101", "101", "101"}
	//err := gamesWindows.AddLevelForNonogram(levelID, cols, rows, winCil)
	//if err != nil {
	//	fmt.Printf("Error adding level for Nonogram: %s\n", err)
	//} else {
	//	fmt.Println("Level added successfully for Nonogram.")
	//}

	mainMenu.MakeMenu().ShowAndRun()
	//app := app.New()
	//gamesWindows.MakeTryWindow(app)

}
