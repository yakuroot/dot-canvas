package task

import (
	"context"
	"fmt"

	"github.com/Neoration/dot-canvas/src/cache"
	"github.com/Neoration/dot-canvas/src/database"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/dustin/go-humanize"
	"go.mongodb.org/mongo-driver/bson"
)

func BotInfoStatus(m *shard.Manager) {
	dotCount, _ := database.Canvas.CountDocuments(context.Background(), bson.M{})
	userCount, _ := database.Users.CountDocuments(context.Background(), bson.M{})
	guildSize := cache.GuildContainer.Size()

	m.ForEach(func(shard shard.Shard) {
		state, ok := shard.(*state.State)
		if !ok {
			return
		}

		state.Gateway().Send(
			context.Background(),
			&gateway.UpdatePresenceCommand{
				Activities: []discord.Activity{{
					Name: humanize.Comma(dotCount) + " Dots | " +
						humanize.Comma(userCount) + " Users | " +
						humanize.Comma(int64(guildSize)) + " Guilds",
					Type: discord.GameActivity,
				}},
				Status: discord.OnlineStatus})
	})
}

func BotCommandInfoStatus(m *shard.Manager) {
	m.ForEach(func(shard shard.Shard) {
		state, ok := shard.(*state.State)
		if !ok {
			return
		}

		if err := state.Gateway().Send(
			context.Background(),
			&gateway.UpdatePresenceCommand{
				Activities: []discord.Activity{{
					Name: "/fill <x> <y> <color>",
					Type: discord.GameActivity,
				}},
				Status: discord.OnlineStatus}); err != nil {
			fmt.Println(err)
		}
	})
}
