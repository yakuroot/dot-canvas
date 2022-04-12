package framework

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
)

type MessageOptions struct {
	Content    string
	TTS        bool
	Embeds     []discord.Embed
	Files      []sendpart.File
	Components discord.ContainerComponents
	AutoReply  bool
	Ephemeral  bool
}
