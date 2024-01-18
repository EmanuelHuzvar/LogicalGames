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

	//lvl3 := gamesWindows.NewLevel{
	//	ID:         "lvl3", // Complete document name
	//	Dimensions: []int{7, 7},
	//	Map: []int{
	//		0, 0, 1, 1, 0, 0, 0,
	//		0, 0, 1, 1, 0, 1, 0,
	//		2, 0, 0, 0, 0, 1, 0,
	//		0, 1, 1, 1, 0, 1, 0,
	//		0, 1, 1, 1, 0, 1, 0,
	//		0, 1, 1, 1, 0, 1, 0,
	//		0, 0, 0, 0, 0, 0, 0,
	//	},
	//}
	//
	//lvl4 := gamesWindows.NewLevel{
	//	ID:         "lvl4", // Complete document name
	//	Dimensions: []int{10, 8},
	//	Map: []int{1, 1, 1, 0, 0, 0, 0, 0,
	//		1, 1, 1, 0, 1, 1, 1, 0,
	//		1, 0, 2, 0, 0, 0, 0, 0,
	//		1, 1, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 1, 1, 1, 1, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 0, 1, 0, 0, 0, 0,
	//		0, 0, 0, 1, 0, 1, 1, 1,
	//		0, 1, 1, 1, 0, 1, 1, 1,
	//	},
	//}
	//
	//lvl5 := gamesWindows.NewLevel{
	//	ID:         "lvl5", // Complete document name
	//	Dimensions: []int{8, 8},
	//	Map: []int{
	//		0, 0, 0, 0, 1, 0, 0, 2,
	//		0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 1, 1, 0,
	//		1, 0, 0, 0, 0, 0, 0, 0,
	//		1, 0, 1, 0, 1, 0, 1, 0,
	//		1, 0, 1, 0, 1, 0, 1, 0,
	//		1, 0, 1, 0, 1, 0, 0, 0,
	//		1, 0, 0, 0, 0, 0, 1, 1,
	//	},
	//}
	//
	//lvl6 := gamesWindows.NewLevel{
	//	ID:         "lvl6", // Complete document name
	//	Dimensions: []int{10, 9},
	//	Map: []int{
	//		1, 1, 0, 1, 1, 0, 0, 0, 0,
	//		1, 1, 0, 1, 0, 0, 0, 1, 1,
	//		1, 1, 0, 1, 0, 0, 0, 0, 1,
	//		1, 1, 2, 1, 0, 0, 1, 0, 1,
	//		0, 0, 0, 0, 0, 0, 0, 0, 1,
	//		0, 1, 0, 1, 1, 0, 1, 1, 0,
	//		0, 1, 0, 1, 0, 0, 0, 0, 0,
	//		0, 1, 0, 1, 1, 0, 1, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 1,
	//	},
	//}
	//
	//lvl7 := gamesWindows.NewLevel{
	//	ID:         "lvl7", // Complete document name
	//	Dimensions: []int{10, 8},
	//	Map: []int{
	//		0, 0, 0, 0, 0, 1, 0, 0,
	//		1, 1, 0, 0, 0, 0, 0, 0,
	//		1, 1, 1, 1, 0, 0, 1, 0,
	//		1, 1, 1, 1, 0, 0, 1, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 1, 0, 1, 0, 1, 0,
	//		0, 1, 1, 0, 0, 0, 0, 0,
	//		0, 0, 1, 0, 0, 2, 1, 0,
	//		1, 0, 0, 0, 0, 0, 1, 0,
	//		0, 0, 0, 1, 0, 0, 0, 0,
	//	},
	//}
	//
	//lvl8 := gamesWindows.NewLevel{
	//	ID:         "lvl8", // Complete document name
	//	Dimensions: []int{10, 8},
	//	Map: []int{
	//		0, 0, 0, 0, 0, 1, 1, 0,
	//		0, 0, 0, 0, 0, 0, 1, 0,
	//		1, 0, 1, 1, 0, 0, 1, 0,
	//		1, 0, 1, 1, 0, 0, 1, 0,
	//		1, 0, 1, 1, 0, 0, 0, 0,
	//		0, 0, 1, 0, 0, 1, 0, 0,
	//		0, 0, 0, 2, 1, 1, 0, 0,
	//		0, 1, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 1,
	//		1, 0, 0, 0, 0, 0, 0, 0,
	//	},
	//}
	//
	//lvl9 := gamesWindows.NewLevel{
	//	ID:         "lvl9", // Complete document name
	//	Dimensions: []int{10, 8},
	//	Map: []int{
	//		0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 1,
	//		1, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 2, 1, 1, 1, 1, 0, 0,
	//		0, 0, 1, 1, 1, 1, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 1, 1, 1, 1, 0, 0,
	//		0, 1, 1, 1, 1, 1, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0,
	//	},
	//}
	//
	//lvl10 := gamesWindows.NewLevel{
	//	ID:         "lvl10", // Complete document name
	//	Dimensions: []int{11, 8},
	//	Map: []int{
	//		1, 0, 0, 1, 0, 0, 0, 0,
	//		1, 0, 0, 1, 0, 1, 0, 0,
	//		1, 0, 0, 0, 0, 1, 0, 0,
	//		0, 1, 0, 0, 0, 1, 0, 0,
	//		0, 1, 1, 1, 1, 1, 0, 0,
	//		0, 1, 1, 1, 1, 1, 0, 0,
	//		0, 1, 1, 1, 1, 1, 0, 0,
	//		0, 1, 1, 0, 0, 0, 0, 0,
	//		0, 0, 2, 0, 0, 0, 0, 1,
	//		0, 1, 0, 0, 0, 0, 0, 0,
	//		1, 0, 0, 0, 0, 0, 0, 0,
	//	},
	//}
	//
	//lvl11 := gamesWindows.NewLevel{
	//	ID:         "lvl11", // Complete document name
	//	Dimensions: []int{11, 9},
	//	Map: []int{
	//		1, 0, 1, 1, 1, 1, 0, 0, 1,
	//		1, 0, 1, 1, 1, 0, 0, 0, 1,
	//		1, 0, 0, 0, 0, 0, 0, 0, 1,
	//		1, 0, 1, 0, 0, 0, 0, 0, 0,
	//		0, 0, 1, 0, 1, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 1, 0, 0,
	//		1, 0, 0, 0, 0, 0, 0, 1, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 0, 0, 0, 0, 2, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 1,
	//	},
	//}
	//
	//lvl12 := gamesWindows.NewLevel{
	//	ID:         "lvl12", // Complete document name
	//	Dimensions: []int{9, 9},
	//	Map: []int{
	//		0, 0, 1, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 1, 0, 0, 0, 0,
	//		0, 1, 1, 1, 0, 0, 0, 0, 1,
	//		0, 1, 1, 1, 1, 1, 1, 0, 1,
	//		0, 1, 1, 1, 1, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 1, 0, 0,
	//		1, 1, 1, 1, 0, 0, 0, 0, 0,
	//		1, 1, 1, 0, 0, 2, 1, 0, 1,
	//		1, 1, 1, 0, 0, 0, 1, 0, 0,
	//	},
	//}
	//
	//lvl13 := gamesWindows.NewLevel{
	//	ID:         "lvl13", // Complete document name
	//	Dimensions: []int{9, 9},
	//	Map: []int{
	//		1, 0, 1, 0, 0, 0, 0, 0, 0,
	//		1, 0, 1, 0, 0, 0, 0, 0, 2,
	//		1, 0, 1, 0, 1, 1, 1, 1, 1,
	//		1, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 1, 1, 1, 1, 1, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 0, 0, 0, 0, 0, 0, 0,
	//	},
	//}
	//
	//lvl14 := gamesWindows.NewLevel{
	//	ID:         "lvl14", // Complete document name
	//	Dimensions: []int{11, 9},
	//	Map: []int{
	//		0, 0, 0, 1, 0, 0, 0, 0, 1,
	//		0, 1, 0, 1, 0, 1, 1, 0, 1,
	//		0, 1, 0, 1, 0, 1, 0, 0, 1,
	//		0, 1, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		1, 0, 0, 0, 0, 0, 0, 0, 0,
	//		1, 0, 0, 0, 0, 0, 0, 1, 0,
	//		1, 1, 0, 1, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 0, 1, 0, 1, 1, 0, 0,
	//		0, 0, 0, 0, 0, 2, 1, 0, 0,
	//	},
	//}
	//
	//lvl15 := gamesWindows.NewLevel{
	//	ID:         "lvl15", // Complete document name
	//	Dimensions: []int{10, 9},
	//	Map: []int{
	//		0, 0, 0, 0, 0, 0, 0, 0, 1,
	//		1, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 1, 1, 0, 0,
	//		0, 1, 1, 1, 1, 0, 0, 0, 0,
	//		0, 1, 1, 1, 0, 0, 0, 0, 0,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 1, 1, 0, 0, 1, 0, 1,
	//		0, 0, 0, 0, 0, 0, 0, 0, 0,
	//		0, 1, 1, 1, 0, 0, 1, 0, 0,
	//		0, 0, 0, 0, 2, 0, 1, 0, 0,
	//	},
	//}
	//
	//err := gamesWindows.AddPaintFloorLevel(lvl3)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl4)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl5)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl6)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl7)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl8)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl9)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl10)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl11)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl12)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl13)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl14)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = gamesWindows.AddPaintFloorLevel(lvl15)
	//if err != nil {
	//	fmt.Println(err)
	//}
	mainMenu.MakeMenu().ShowAndRun()
	//app := app.New()
	//gamesWindows.MakeTryWindow(app)

}
