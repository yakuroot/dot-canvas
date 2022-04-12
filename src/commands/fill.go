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

type fillCommand struct{}

func (fillCommand) Run(ctx *framework.Interaction, args []string) error {
	if !fillFilter(ctx) {
		return nil
	}

	if len(args) < 4 {
		return nil
	}

	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[2])
	color := strings.ToLower(args[3])

	hex, exist := base.ColorList[color]
	if !exist {
		return nil
	}

	go draw(x, y, hex, ctx.Author.ID)

	go ctx.Reply(framework.MessageOptions{
		Embeds: []discord.Embed{{
			Color: base.ColorGreen,
			Description: base.CheckSign + " " + locales.Text("fill.complete", ctx.Language, map[string]interface{}{
				"x": x, "y": y, "color": locales.Text("command_options.fill.color.options."+color, ctx.Language)})}},
		Ephemeral: true})

	return nil
}

var Fill = &base.Command{
	Name:                       "fill",
	NameLocalizationKey:        "command_name.fill",
	Description:                "Color a coordinate.",
	DescriptionLocalizationKey: "command_description.fill",
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
		base.ChoiceStructure{
			Name:                       "color",
			Description:                "Select a color",
			NameLocalizationKey:        "command_options.fill.color.name",
			DescriptionLocalizationKey: "command_options.fill.color.description",
			Choices: []base.ChoiceOptions{
				{Name: "Red", NameLocalizationKey: "command_options.fill.color.options.red", Value: "red"},
				{Name: "Orange", NameLocalizationKey: "command_options.fill.color.options.orange", Value: "orange"},
				{Name: "Yellow", NameLocalizationKey: "command_options.fill.color.options.yellow", Value: "yellow"},
				{Name: "Lime", NameLocalizationKey: "command_options.fill.color.options.lime", Value: "lime"},
				{Name: "Green", NameLocalizationKey: "command_options.fill.color.options.green", Value: "green"},
				{Name: "Olive", NameLocalizationKey: "command_options.fill.color.options.olive", Value: "olive"},
				{Name: "Aqua", NameLocalizationKey: "command_options.fill.color.options.aqua", Value: "aqua"},
				{Name: "Teal", NameLocalizationKey: "command_options.fill.color.options.teal", Value: "teal"},
				{Name: "Blue", NameLocalizationKey: "command_options.fill.color.options.blue", Value: "blue"},
				{Name: "Navy", NameLocalizationKey: "command_options.fill.color.options.navy", Value: "navy"},
				{Name: "Purple", NameLocalizationKey: "command_options.fill.color.options.purple", Value: "purple"},
				{Name: "Fuchsia", NameLocalizationKey: "command_options.fill.color.options.fuchsia", Value: "fuchsia"},
				{Name: "Marron", NameLocalizationKey: "command_options.fill.color.options.marron", Value: "marron"},
				{Name: "Brown", NameLocalizationKey: "command_options.fill.color.options.brown", Value: "brown"},
				{Name: "White", NameLocalizationKey: "command_options.fill.color.options.white", Value: "white"},
				{Name: "Sliver", NameLocalizationKey: "command_options.fill.color.options.sliver", Value: "sliver"},
				{Name: "Gray", NameLocalizationKey: "command_options.fill.color.options.gray", Value: "gray"},
				{Name: "Black", NameLocalizationKey: "command_options.fill.color.options.black", Value: "black"},
			},
			Required: true,
		},
	},
	CommandRunable: fillCommand{},
}
