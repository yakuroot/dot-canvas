package cache

import (
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
)

var (
	GuildContainer = Container[struct{}, discord.GuildID]{Items: map[discord.GuildID]struct{}{}}
	UserContainer  = Container[time.Time, discord.UserID]{Items: map[discord.UserID]time.Time{}}
)
