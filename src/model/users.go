package model

import "time"

type UserRecordProps struct {
	X         int       `bson:"x"`
	Y         int       `bson:"y"`
	Color     string    `bson:"color"`
	CreatedAt time.Time `bson:"createdAt"`
}

type UserProps struct {
	ID      string            `bson:"_id"`
	Records []UserRecordProps `bson:"records"`
}
