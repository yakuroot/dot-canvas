package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Neoration/dot-canvas/src/base"
	"github.com/Neoration/dot-canvas/src/config"
	"github.com/Neoration/dot-canvas/src/database"
	"github.com/Neoration/dot-canvas/src/framework"
	"github.com/Neoration/dot-canvas/src/locales"
	"github.com/diamondburned/arikawa/v3/discord"
	"go.mongodb.org/mongo-driver/bson"
)

type canvasCommand struct{}

func (canvasCommand) Run(ctx *framework.Interaction, _ []string) error {
	dotCount, _ := database.Canvas.CountDocuments(context.Background(), bson.M{})
	userCount, _ := database.Users.CountDocuments(context.Background(), bson.M{})

	ctx.Reply(framework.MessageOptions{
		Embeds: []discord.Embed{{
			Color: base.ColorCanvas,
			Title: base.CanvasSign + " " + locales.Text("canvas.title", ctx.Language, map[string]interface{}{
				"year":  time.Now().Year(),
				"month": fmt.Sprintf("%02d", int(time.Now().Month()))}),
			Fields: []discord.EmbedField{
				{Name: locales.Text("canvas.dot", ctx.Language), Value: fmt.Sprintf("%d", dotCount), Inline: true},
				{Name: locales.Text("canvas.users", ctx.Language), Value: fmt.Sprintf("%d", userCount), Inline: true}},
			Image: &discord.EmbedImage{URL: config.ImageURL}}},
	})
	return nil
}

var Canvas = &base.Command{
	Name:                       "canvas",
	NameLocalizationKey:        "command_name.canvas",
	Description:                "Check out the canvases that have been made so far.",
	DescriptionLocalizationKey: "command_description.canvas",
	CommandRunable:             canvasCommand{},
}
