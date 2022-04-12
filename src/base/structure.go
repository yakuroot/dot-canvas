package base

import (
	"github.com/Neoration/dot-canvas/src/locales"
	"github.com/diamondburned/arikawa/v3/discord"
)

type StringStructure struct {
	Name,
	Description,
	NameLocalizationKey,
	DescriptionLocalizationKey string
	Required bool
}

type MaxMinIntegerStructure struct {
	Name,
	Description,
	NameLocalizationKey,
	DescriptionLocalizationKey string
	Max,
	Min int
	Required bool
}

type (
	ChoiceOptions struct {
		Name                string
		NameLocalizationKey string
		Value               string
	}

	ChoiceStructure struct {
		Name,
		Description,
		NameLocalizationKey,
		DescriptionLocalizationKey string
		Choices  []ChoiceOptions
		Required bool
	}
)

type CommandStructure interface {
	SlashCommandOptionBuilder() discord.CommandOption
	//Filter(ctx *framework.Interaction, args []string) (pass bool)
}

func (s StringStructure) SlashCommandOptionBuilder() discord.CommandOption {
	nameLocalization := make(discord.StringLocales)
	descriptionLocalization := make(discord.StringLocales)
	for _, lng := range locales.Languages {
		nameLocalization[lng] = locales.Text(s.NameLocalizationKey, string(lng))
		descriptionLocalization[lng] = locales.Text(s.DescriptionLocalizationKey, string(lng))
	}

	return &discord.StringOption{
		OptionName:               s.Name,
		OptionNameLocalizations:  nameLocalization,
		Description:              s.Description,
		DescriptionLocalizations: descriptionLocalization,
		Required:                 s.Required,
	}
}

func (s MaxMinIntegerStructure) SlashCommandOptionBuilder() discord.CommandOption {
	nameLocalization := make(discord.StringLocales)
	descriptionLocalization := make(discord.StringLocales)
	for _, lng := range locales.Languages {
		nameLocalization[lng] = locales.Text(s.NameLocalizationKey, string(lng))
		descriptionLocalization[lng] = locales.Text(s.DescriptionLocalizationKey, string(lng))
	}

	return &discord.IntegerOption{
		OptionName:               s.Name,
		OptionNameLocalizations:  nameLocalization,
		Description:              s.Description,
		DescriptionLocalizations: descriptionLocalization,
		Required:                 s.Required,
		Max:                      &s.Max,
		Min:                      &s.Min,
	}
}

func (s ChoiceStructure) SlashCommandOptionBuilder() discord.CommandOption {
	choices := make([]discord.StringChoice, 0)
	for _, opt := range s.Choices {
		nameLocalization := make(discord.StringLocales)
		for _, lng := range locales.Languages {
			nameLocalization[lng] = locales.Text(s.NameLocalizationKey, string(lng))
		}

		choices = append(choices, discord.StringChoice{
			Name:              opt.Name,
			NameLocalizations: nameLocalization,
			Value:             opt.Value})
	}

	nameLocalization := make(discord.StringLocales)
	descriptionLocalization := make(discord.StringLocales)
	for _, lng := range locales.Languages {
		nameLocalization[lng] = locales.Text(s.NameLocalizationKey, string(lng))
		descriptionLocalization[lng] = locales.Text(s.DescriptionLocalizationKey, string(lng))
	}

	return &discord.StringOption{
		OptionName:               s.Name,
		OptionNameLocalizations:  nameLocalization,
		Description:              s.Description,
		DescriptionLocalizations: descriptionLocalization,
		Required:                 s.Required,
		Choices:                  choices,
	}
}
