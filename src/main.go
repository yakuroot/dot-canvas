package main

import (
	"context"
	"log"

	"github.com/Neoration/dot-canvas/src/base"
	"github.com/Neoration/dot-canvas/src/commands"
	"github.com/Neoration/dot-canvas/src/config"
	"github.com/Neoration/dot-canvas/src/event"
	"github.com/Neoration/dot-canvas/src/task"
	"github.com/Neoration/dot-canvas/src/web"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/robfig/cron/v3"
)

var (
	ShardManager *shard.Manager
)

func init() {
	newShard := state.NewShardFunc(func(m *shard.Manager, s *state.State) {
		s.AddIntents(gateway.IntentGuilds)
		s.AddHandler(func(i *gateway.InteractionCreateEvent) { event.InteractionCreate(s, i) })
		s.AddHandler(event.GuildCreateHandler)
		s.AddHandler(event.GuildDeleteHandler)
	})

	var err error

	ShardManager, err = shard.NewManager("Bot "+config.Token, newShard)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if err := ShardManager.Open(context.Background()); err != nil {
		log.Fatalf("Error: %v", err)
	}

	var shardNum int

	ShardManager.ForEach(func(s shard.Shard) {
		state, ok := s.(*state.State)
		if !ok {
			return
		}

		g, err := state.Guilds()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Shard %d/%d (%d Guilds) Started.", shardNum, ShardManager.NumShards()-1, len(g))
	})
}

func init() {
	base.Commands.Register(
		commands.Fill,
		commands.FillHex,
		commands.Canvas,
		commands.Help,
	)

	cmds := base.Commands.SlashCommandBuilder()

	ShardManager.ForEach(func(shard shard.Shard) {
		state, ok := shard.(*state.State)
		if !ok {
			return
		}

		for _, cmd := range cmds {
			if _, err := state.CreateCommand(state.Ready().Application.ID, cmd); err != nil {
				log.Fatalf("Error: %v", err)
			}
		}

		state.PresenceSet(discord.NullGuildID, &discord.Presence{
			Status:     discord.OnlineStatus,
			Activities: []discord.Activity{{Name: "DotCanvas!", Type: discord.GameActivity}}}, true)
	})
}

func init() {
	c := cron.New()
	c.AddFunc("0 0 1 * *", task.MonthlyInit)
	c.AddFunc("1/2 * * * *", func() { task.BotInfoStatus(ShardManager) })
	c.AddFunc("*/2 * * * *", func() { task.BotCommandInfoStatus(ShardManager) })
	c.Start()
}

func main() {
	defer ShardManager.Close()
	web.Start()
}
