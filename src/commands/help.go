package commands

import (
	"github.com/Neoration/dot-canvas/src/base"
	"github.com/Neoration/dot-canvas/src/config"
	"github.com/Neoration/dot-canvas/src/framework"
	"github.com/Neoration/dot-canvas/src/locales"
	"github.com/diamondburned/arikawa/v3/discord"
)

type helpCommand struct{}

func (helpCommand) Run(ctx *framework.Interaction, _ []string) error {
	ctx.Reply(framework.MessageOptions{
		Embeds: []discord.Embed{{
			Color:       base.ColorAqua,
			Title:       base.AlertSign + " " + locales.Text("help.title", ctx.Language),
			Description: locales.Text("help.description", ctx.Language),
			Fields: []discord.EmbedField{
				{
					Name:   locales.Text("help.fill.name", ctx.Language),
					Value:  locales.Text("help.fill.value", ctx.Language),
					Inline: true,
				},
				{
					Name:   locales.Text("help.fillhex.name", ctx.Language),
					Value:  locales.Text("help.fillhex.value", ctx.Language),
					Inline: true,
				},
				{
					Name:   locales.Text("help.canvas.name", ctx.Language),
					Value:  locales.Text("help.canvas.value", ctx.Language),
					Inline: true,
				},
			},
			Image: &discord.EmbedImage{URL: locales.Text("help.imageUrl", ctx.Language)},
		}},
		Components: discord.ContainerComponents{
			&discord.ActionRowComponent{
				&discord.ButtonComponent{
					Style: discord.LinkButtonStyle(config.GetInviteURL()),
					Label: locales.Text("help.button.invite", ctx.Language),
				},
				&discord.ButtonComponent{
					Style: discord.LinkButtonStyle("https://github.com/Neoration/dot-canvas"),
					Label: locales.Text("help.button.github", ctx.Language),
				},
				&discord.ButtonComponent{
					Style: discord.LinkButtonStyle("https://discord.com/users/726534821572116512"),
					Label: locales.Text("help.button.profile", ctx.Language),
				},
			},
		},
	})
	return nil
}

var Help = &base.Command{
	Name:                       "help",
	NameLocalizationKey:        "command_name.help",
	Description:                "Explore the BOT's help.",
	DescriptionLocalizationKey: "command_description.help",
	CommandRunable:             helpCommand{},
}
