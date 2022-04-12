package event

import (
	"strings"

	"github.com/Neoration/dot-canvas/src/base"
	"github.com/Neoration/dot-canvas/src/framework"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

func InteractionCreate(s *state.State, i *gateway.InteractionCreateEvent) {
	sender := i.Sender()

	if sender == nil || !i.SenderID().IsValid() {
		return
	}

	if sender.Bot {
		return
	}

	data, ok := i.InteractionEvent.Data.(*discord.CommandInteraction)
	if !ok {
		return
	}

	commandName := data.Name
	args := []string{data.Name}
	if data.Options != nil && len(data.Options) != 0 {
		for _, opt := range data.Options {
			args = append(args, getArgs(opt)...)
		}
	}

	if cmd, exist := base.Commands.Get(commandName); exist {
		cmd.RunCommand(framework.NewInteractionFramework(s, &i.InteractionEvent), args)
	}

}

func getArgs(data discord.CommandInteractionOption) []string {
	args := make([]string, 0)

	if v := data.Value.String(); v != "" {
		args = append(args, strings.ReplaceAll(v, `"`, ""))
	} else {
		args = append(args, data.Name)
	}

	if data.Options != nil && len(data.Options) != 0 {
		for _, otp := range data.Options {
			args = append(args, getArgs(otp)...)
		}
	}

	return args
}
