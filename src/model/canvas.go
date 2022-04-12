package model

type CanvasProps struct {
	X            int    `bson:"x"`
	Y            int    `bson:"y"`
	Color        string `bson:"color"` // "#000000"
	RecentlyUser string `bson:"recentlyUser"`
}
