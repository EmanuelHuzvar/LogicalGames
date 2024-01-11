package db

import (
	"ProjectMarekEmanuel/marek-games/PaintFloor/models"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
)

var (
	ctx       = context.Background()
	projectID = "taskmanager-24b1f"
	keyPath   = "marek-games/PaintFloor/db/authentification.json"
)

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
	var levelData models.Level
	if err := doc.DataTo(&levelData); err != nil {
		return nil, nil, err
	}

	return levelData.Dimensions, levelData.Map, nil
}
