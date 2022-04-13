package commands

import (
	"context"
	"time"

	"github.com/Neoration/dot-canvas/src/base"
	"github.com/Neoration/dot-canvas/src/cache"
	"github.com/Neoration/dot-canvas/src/canvas"
	"github.com/Neoration/dot-canvas/src/database"
	"github.com/Neoration/dot-canvas/src/framework"
	"github.com/Neoration/dot-canvas/src/locales"
	"github.com/Neoration/dot-canvas/src/model"
	"github.com/Neoration/dot-canvas/src/queue"
	"github.com/diamondburned/arikawa/v3/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const coolDown = 15 // seconds

var (
	imageNameQueue  = queue.New[string](100)
	canvasImageName = ""
)

func fillFilter(ctx *framework.Interaction) (pass bool) {
	t, ok := cache.UserContainer.Get(ctx.Author.ID)

	if ok {
		duration := t.Sub(time.Now())
		ctx.Reply(framework.MessageOptions{
			Embeds: []discord.Embed{{
				Color: base.ColorRed,
				Description: base.XSign + " " + locales.Text("fill.timelimit", ctx.Language, map[string]interface{}{
					"seconds": int(duration.Minutes())*60 + int(duration.Seconds())})}},
			Ephemeral: true})
		return false
	}
	return true
}

func draw(x, y int, color string, executor discord.UserID) {
	cache.UserContainer.Set(executor, time.Now().Add(coolDown*time.Second))
	time.AfterFunc(coolDown*time.Second, func() { cache.UserContainer.Remove(executor) })

	if imageNameQueue.Size() > 0 {
		oldName := imageNameQueue.Pop()
		canvas.RemoveImage(oldName)
	}

	canvas.SetDot(x, y, color)

	fileName := base.GetRandCode()
	canvas.SaveNamedImage(fileName)
	imageNameQueue.Append(fileName)
	canvasImageName = fileName

	go database.Canvas.UpdateOne(
		context.Background(),
		bson.M{"x": x, "y": y},
		bson.M{"$set": bson.M{
			"color": color, "recentlyUser": executor.String()}},
		&options.UpdateOptions{Upsert: base.Pointer(true)})

	go database.Users.UpdateOne(
		context.Background(),
		bson.M{"_id": executor.String()},
		bson.M{"$push": bson.M{
			"records": model.UserRecordProps{
				X: x, Y: y, Color: color, CreatedAt: time.Now()}}},
		&options.UpdateOptions{Upsert: base.Pointer(true)})
}
