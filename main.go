package main

import "ProjectMarekEmanuel/mainMenu"

func main() {
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
	//levelID := "2"
	//cols := []int{2, 3, 2, 1}
	//rows := []int{1, 2, 4, 1}
	//
	//// Call the AddLevelForNonogram method with the example values
	//err := gamesWindows.AddLevelForNonogram(levelID, cols, rows)
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
