package task

import (
	"context"
	"fmt"
	"time"

	"github.com/Neoration/dot-canvas/src/canvas"
	"github.com/Neoration/dot-canvas/src/database"
	"github.com/Neoration/dot-canvas/src/model"
	"go.mongodb.org/mongo-driver/bson"
)

func MonthlyInit() {
	now := time.Now()
	canvas.SaveNamedImage(fmt.Sprintf("%d%d", now.Year(), now.Month()))

	tryCounts := 0

	if res, err := database.Users.Find(context.Background(), bson.M{}); err == nil {
		var users []model.UserProps
		if err := res.All(context.Background(), &users); err == nil {
			for _, u := range users {
				tryCounts += len(u.Records)
			}
		}
	}

	dotCounts, _ := database.Canvas.CountDocuments(context.Background(), bson.M{})
	userCounts, _ := database.Users.CountDocuments(context.Background(), bson.M{})

	database.Records.InsertOne(
		context.Background(),
		model.RecordProps{
			ID:         fmt.Sprintf("%d-%s", now.Year(), now.Month().String()),
			DotCounts:  int(dotCounts),
			UserCounts: int(userCounts),
			TryCounts:  tryCounts,
		})

	database.Canvas.DeleteMany(context.Background(), bson.M{})
	database.Users.DeleteMany(context.Background(), bson.M{})
}
