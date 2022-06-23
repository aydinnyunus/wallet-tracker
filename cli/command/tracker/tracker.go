package tracker

import (
	"github.com/spf13/cobra"
	"io"
	"os"
)

// global constants for file
const ()

// global variables (not cool) for this file
var ()

// NewTrackCommand serves the base command for querying redis.
//
// Adds flags and subcommands for redis commands here,
func NewTrackCommand() *cobra.Command {
	var trackCommands = &cobra.Command{
		Use:   "tracker",
		Short: "Track Scammers with Track Command",
		Long:  `Track Scammers with Track Command.`,
		RunE:  ShowHelp(os.Stdout),
	}

	trackCommands.AddCommand(
		TrackCommand(),
		TrackWebsocketCommand(),
	)
	return trackCommands
}

// ShowHelp shows the command help.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}
