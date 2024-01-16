package gamesWindows

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/option"
)

var (
	ctx       = context.Background()
	projectID = "taskmanager-24b1f"
	keyPath   = "mainMenu/gamesWindows/authentification.json"
)

type Level struct {
	ID     string     `firestore:"Level"`
	Fields [][]string `firestore:"-"`
}
type LevelNonogram struct {
	ID      string   `firestore:"id"`
	Cols    []int    `firestore:"cols"`
	Rows    []int    `firestore:"rows"`
	ColsWin []string `firestore:"winCol"`
}

func GetLevelByID(levelID string, gameName string) (Level, error) {
	var level Level

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(keyPath))
	if err != nil {
		return level, err
	}
	defer client.Close()

	doc, err := client.Collection(gameName).Doc(levelID).Get(ctx)
	if err != nil {
		return level, err
	}

	level.ID = levelID // Assign the document ID to the level ID

	data := doc.Data()
	// Initialize the Fields slice based on the expected number of rows
	for i := 0; i <= 7; i++ {
		key := fmt.Sprintf("%d", i)
		if row, ok := data[key].([]interface{}); ok {
			var rowStrings []string
			for _, col := range row {
				if colStr, ok := col.(string); ok {
					rowStrings = append(rowStrings, colStr)
				} else {
					return level, fmt.Errorf("expected a string in row, got %T", col)
				}
			}
			level.Fields = append(level.Fields, rowStrings)
		} else {
			return level, fmt.Errorf("expected '%s' to be []interface{}, got %T", key, data[key])
		}
	}

	return level, nil
}
func AddLevelForBubble(levelID string, fields [][]string) error {
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(keyPath))
	if err != nil {
		return err
	}
	defer client.Close()

	// Prepare the data to be written to Firestore
	data := make(map[string]interface{})
	for i, row := range fields {
		data[fmt.Sprintf("%d", i)] = row
	}

	// Use the Set method to create a new document with the provided level ID
	_, err = client.Collection("Bubble").Doc(levelID).Set(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func IterateBunks(level Level) {
	for bunkID, bunkMap := range level.Fields {
		fmt.Println("Bunk ID:", bunkID)
		for key, value := range bunkMap {
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		}
	}
}
func GetLevelNonogramByID(levelID string) (LevelNonogram, error) {
	var levelNonogram LevelNonogram

	// Initialize the Firestore client.
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(keyPath))
	if err != nil {
		return levelNonogram, err
	}
	defer client.Close()

	// Retrieve the document with the given ID from the Nonogram collection.
	doc, err := client.Collection("Nonogram").Doc(levelID).Get(ctx)
	if err != nil {
		return levelNonogram, err
	}

	// Extract the document data.
	data := doc.Data()

	// Extract and assign the "cols" field.
	if cols, ok := data["cols"].([]interface{}); ok {
		for _, col := range cols {
			colInt, ok := col.(int64) // Firestore stores integers as int64
			if !ok {
				return levelNonogram, fmt.Errorf("type assertion to int64 failed for cols")
			}
			levelNonogram.Cols = append(levelNonogram.Cols, int(colInt))
		}
	} else {
		return levelNonogram, fmt.Errorf("expected 'cols' to be a one-dimensional array of integers")
	}

	// Extract and assign the "rows" field.
	if rows, ok := data["rows"].([]interface{}); ok {
		for _, row := range rows {
			rowInt, ok := row.(int64) // Firestore stores integers as int64
			if !ok {
				return levelNonogram, fmt.Errorf("type assertion to int64 failed for rows")
			}
			levelNonogram.Rows = append(levelNonogram.Rows, int(rowInt))
		}
	} else {
		return levelNonogram, fmt.Errorf("expected 'rows' to be a one-dimensional array of integers")
	}
	if colsWin, ok := data["winCol"].([]interface{}); ok {
		for _, row := range colsWin {
			rowInt, ok := row.(string) // Firestore stores integers as int64
			if !ok {
				return levelNonogram, fmt.Errorf("type assertion to int64 failed for rows")
			}
			levelNonogram.ColsWin = append(levelNonogram.ColsWin, rowInt)
		}
	} else {
		return levelNonogram, fmt.Errorf("expected 'rows' to be a one-dimensional array of integers")
	}
	// Assign the document ID to the level ID.
	levelNonogram.ID = levelID

	return levelNonogram, nil
}
func AddLevelForNonogram(levelID string, cols []int, rows []int, colsWin []string) error {
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(keyPath))
	if err != nil {
		return err
	}
	defer client.Close()

	// Prepare the data to be written to Firestore
	data := map[string]interface{}{
		"cols":   cols,
		"rows":   rows,
		"winCol": colsWin,
	}

	// Use the Set method to create a new document with the provided level ID
	_, err = client.Collection("Nonogram").Doc(levelID).Set(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
func LoadLevelData(levelID string) ([]int, []int, error) {
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(keyPath))
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	// Retrieve the level document from Firestore.
	doc, err := client.Collection("PaintFloor").Doc(levelID).Get(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Map the document to the Level struct.
	var levelData LevelPaintFloor
	if err := doc.DataTo(&levelData); err != nil {
		return nil, nil, err
	}

	return levelData.Dimensions, levelData.Map, nil
}
