package neodash

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// global constants for file
const ()

// global variables (not cool) for this file
var ()

// NewNeodashCommand serves the base command for querying neodash.
//
// Adds flags and subcommands for neodash commands here,
func NewNeodashCommand() *cobra.Command {
	var neodashCommands = &cobra.Command{
		Use:   "neodash",
		Short: "Starting Neodash Service",
		Long:  `Starting Neodash Service to visualize Neo4J Database`,
		RunE:  ShowHelp(os.Stdout),
	}

	neodashCommands.AddCommand(
		StartCommand(),
	)
	return neodashCommands
}

// ShowHelp shows the command help.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}
