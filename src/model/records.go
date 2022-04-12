package model

type RecordProps struct {
	ID         string `bson:"_id"`
	DotCounts  int    `bson:"dotCounts"`
	UserCounts int    `bson:"userCounts"`
	TryCounts  int    `bson:"tryCounts"`
}
