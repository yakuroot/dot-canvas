package base

import (
	"bytes"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/Neoration/dot-canvas/src/config"
	"github.com/Neoration/dot-canvas/src/framework"
	"github.com/Neoration/dot-canvas/src/locales"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
)

type (
	Command struct {
		Name,
		Description,
		NameLocalizationKey,
		DescriptionLocalizationKey string
		Structure []CommandStructure
		CommandRunable
	}

	CommandRunable interface {
		Run(ctx *framework.Interaction, args []string) error
	}

	CommandHandler struct {
		Cmd []*Command
	}
)

var Commands = &CommandHandler{make([]*Command, 0)}

func (h *CommandHandler) Register(commands ...*Command) *CommandHandler {
	h.Cmd = append(h.Cmd, commands...)
	return h
}

func (h *CommandHandler) Get(name string) (*Command, bool) {
	for _, cmd := range h.Cmd {
		if cmd.Name == name {
			return cmd, true
		}
	}
	return nil, false
}

func (h *CommandHandler) SlashCommandBuilder() (cmd []api.CreateCommandData) {
	for _, c := range h.Cmd {
		cmd = append(cmd, c.SlashCommandBuilder())
	}
	return
}

func (c *Command) SlashCommandBuilder() api.CreateCommandData {
	nameLocalization := make(discord.StringLocales)
	descriptionLocalization := make(discord.StringLocales)
	for _, lng := range locales.Languages {
		nameLocalization[lng] = locales.Text(c.NameLocalizationKey, string(lng))
		descriptionLocalization[lng] = locales.Text(c.DescriptionLocalizationKey, string(lng))
	}

	cmd := api.CreateCommandData{
		Name:                     c.Name,
		NameLocalizations:        nameLocalization,
		Description:              c.Description,
		DescriptionLocalizations: descriptionLocalization,
		Type:                     discord.ChatInputCommand,
	}

	if len(c.Structure) < 1 {
		return cmd
	}

	cmd.Options = make(discord.CommandOptions, 0)
	for _, s := range c.Structure {
		cmd.Options = append(cmd.Options, s.SlashCommandOptionBuilder())
	}

	return cmd
}

func (c *Command) RunCommand(ctx *framework.Interaction, args []string) {
	if ctx.State == nil {
		return
	}

	err := make(chan error)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				switch data := r.(type) {
				case error:
					err <- data
				case string:
					err <- errors.New(data)
				default:
					err <- fmt.Errorf("error: %v", err)
				}
			}
		}()

		err <- c.Run(ctx, args)
	}()

	if e := <-err; e != nil {
		errorHandler(ctx, args, e)
	}
}

func errorHandler(ctx *framework.Interaction, args []string, err error) {
	errString := "Error: \n" + err.Error() + "\n\nStack: \n" + string(debug.Stack())
	reader := bytes.NewReader([]byte(errString))

	ctx.State.SendMessageComplex(
		config.ErrorLogChannel,
		api.SendMessageData{
			Content: fmt.Sprintf(
				"Error Occured.\nAuthor: %s (%s)\nCommand: %s",
				ctx.Author.Tag(),
				ctx.Author.ID.String(),
				strings.Join(args, " ")),
			Files: []sendpart.File{{
				Name:   "error.txt",
				Reader: reader}}})
}
