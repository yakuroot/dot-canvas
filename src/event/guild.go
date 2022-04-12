package event

import (
	"github.com/Neoration/dot-canvas/src/cache"
	"github.com/diamondburned/arikawa/v3/gateway"
)

func GuildCreateHandler(g *gateway.GuildCreateEvent) {
	guildID := g.ID
	go cache.GuildContainer.Set(guildID, struct{}{})
}

func GuildDeleteHandler(g *gateway.GuildDeleteEvent) {
	guildID := g.ID
	go cache.GuildContainer.Remove(guildID)
}
