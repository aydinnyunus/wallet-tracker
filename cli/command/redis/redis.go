package redis

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// global constants for file
const ()

// global variables (not cool) for this file
var ()

// NewRedisCommand serves the base command for querying redis.
//
// Adds flags and subcommands for redis commands here,
func NewRedisCommand() *cobra.Command {
	var redisCommands = &cobra.Command{
		Use:   "redis",
		Short: "Query and Play with Redis",
		Long:  `Query and play with trips in the Redis of your choice.`,
		RunE:  ShowHelp(os.Stdout),
	}

	redisCommands.AddCommand(
		GetCommand(),
	)
	return redisCommands
}

// ShowHelp shows the command help.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}
