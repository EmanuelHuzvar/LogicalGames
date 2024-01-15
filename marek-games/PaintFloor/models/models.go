package models

type Level struct {
	Dimensions []int `firestore:"dimensions"`
	Map        []int `firestore:"map"`
}
