package commands

import (
	"strconv"
	"strings"

	"github.com/Neoration/dot-canvas/src/base"
	"github.com/Neoration/dot-canvas/src/config"
	"github.com/Neoration/dot-canvas/src/framework"
	"github.com/Neoration/dot-canvas/src/locales"
	"github.com/diamondburned/arikawa/v3/discord"
)

type fillHexCommand struct{}

func (fillHexCommand) Run(ctx *framework.Interaction, args []string) error {
	if !fillFilter(ctx) {
		return nil
	}

	if len(args) < 4 {
		return nil
	}

	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[2])
	hex := strings.ToLower(args[3])
	if !strings.HasPrefix(hex, "#") {
		hex = "#" + hex
	}

	if _, err := strconv.ParseInt(strings.ReplaceAll(hex, "#", ""), 16, 32); err != nil {
		ctx.Reply(framework.MessageOptions{
			Embeds: []discord.Embed{{
				Color:       base.ColorRed,
				Description: base.XSign + " " + locales.Text("fill.unknown_color", ctx.Language)}},
			Ephemeral: true})
		return nil
	}

	go draw(x, y, hex, ctx.Author.ID)

	go ctx.Reply(framework.MessageOptions{
		Embeds: []discord.Embed{{
			Color: base.ColorGreen,
			Description: base.CheckSign + " " + locales.Text("fill.complete", ctx.Language, map[string]interface{}{
				"x": x, "y": y, "color": hex})}},
		Ephemeral: true})

	return nil
}

var FillHex = &base.Command{
	Name:                       "fillhex",
	NameLocalizationKey:        "command_name.fillhex",
	Description:                "Fill one coordinate with the desired color.",
	DescriptionLocalizationKey: "command_description.fillhex",
	Structure: []base.CommandStructure{
		base.MaxMinIntegerStructure{
			Name:                       "x",
			Description:                "Input X Coordinates",
			NameLocalizationKey:        "command_options.fill.x.name",
			DescriptionLocalizationKey: "command_options.fill.x.description",
			Min:                        0,
			Max:                        config.CanvasWidth,
			Required:                   true,
		},
		base.MaxMinIntegerStructure{
			Name:                       "y",
			Description:                "Input Y Coordinates",
			NameLocalizationKey:        "command_options.fill.y.name",
			DescriptionLocalizationKey: "command_options.fill.y.description",
			Min:                        0,
			Max:                        config.CanvasHeigh,
			Required:                   true,
		},
		base.StringStructure{
			Name:                       "color",
			Description:                "Please input a hexadecimal color. It must be of the same format as #000000",
			NameLocalizationKey:        "command_options.fill.color_hex.name",
			DescriptionLocalizationKey: "command_options.fill.color_hex.description",
			Required:                   true,
		},
	},
	CommandRunable: fillHexCommand{},
}
