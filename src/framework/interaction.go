package framework

import (
	"errors"

	"github.com/Neoration/dot-canvas/src/locales"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

type Interaction struct {
	State       *state.State
	GuildID     discord.GuildID
	ChannelID   discord.ChannelID
	Interaction *discord.InteractionEvent
	Author      discord.User
	Language    string
	defered     bool
}

func NewInteractionFramework(s *state.State, i *discord.InteractionEvent) *Interaction {
	ctx := new(Interaction)

	sender := i.Sender()

	ctx.State = s
	ctx.GuildID = i.GuildID
	ctx.ChannelID = i.ChannelID
	ctx.Interaction = i
	ctx.Author = *sender
	ctx.Language = string(i.Locale)

	if !locales.IsSupportLanguage(string(i.Locale)) {
		ctx.Language = "en"
	}

	return ctx
}

func (i *Interaction) interactionRespond(
	opt MessageOptions, t api.InteractionResponseType) error {
	var ephemeral api.InteractionResponseFlags

	if opt.Ephemeral {
		ephemeral = api.EphemeralResponse
	}

	if err := i.State.RespondInteraction(
		i.Interaction.ID,
		i.Interaction.Token,
		api.InteractionResponse{
			Type: t,
			Data: &api.InteractionResponseData{
				Content:    option.NewNullableString(opt.Content),
				Embeds:     &opt.Embeds,
				Components: &opt.Components,
				Files:      opt.Files,
				Flags:      ephemeral,
				TTS:        false}}); err != nil {
		return err
	}

	return nil
}

func (i *Interaction) Reply(opt MessageOptions) error {
	return i.interactionRespond(opt, api.MessageInteractionWithSource)
}

func (i *Interaction) Edit(opt MessageOptions) error {
	_, err := i.State.EditInteractionResponse(
		i.Interaction.AppID,
		i.Interaction.Token,
		api.EditInteractionResponseData{
			Content:    option.NewNullableString(opt.Content),
			Embeds:     &opt.Embeds,
			Components: &opt.Components,
			Files:      opt.Files})

	if err != nil {
		return err
	}

	return nil
}

func (i *Interaction) Update(opt MessageOptions) error {
	return i.interactionRespond(opt, api.UpdateMessage)
}

func (i *Interaction) Defer(ephemeral bool) error {
	i.defered = true

	var flag api.InteractionResponseFlags

	if ephemeral {
		flag = api.EphemeralResponse
	}

	if err := i.State.RespondInteraction(
		i.Interaction.ID,
		i.Interaction.Token,
		api.InteractionResponse{
			Type: api.DeferredMessageInteractionWithSource,
			Data: &api.InteractionResponseData{Flags: flag}}); err != nil {
		return err
	}

	return nil
}

func (i *Interaction) DeferReply(opt MessageOptions) error {
	if !i.defered {
		return errors.New("not deffered interaction")
	}
	return i.Edit(opt)
}

func (i *Interaction) DeferUpdate(opt MessageOptions) error {
	if !i.defered {
		return errors.New("not deffered interaction")
	}
	return i.interactionRespond(opt, api.DeferredMessageUpdate)
}
