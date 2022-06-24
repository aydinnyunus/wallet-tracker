package commands

import (
	"github.com/aydinnyunus/wallet-tracker/cli/command/neodash"
	"github.com/aydinnyunus/wallet-tracker/cli/command/redis"
	"github.com/aydinnyunus/wallet-tracker/cli/command/tracker"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// NewWalletTrackerCommand is the highest command in the hierarchy and all commands root from it.
//nolint:funlen
func NewWalletTrackerCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	// Define our command
	rootCmd := &cobra.Command{
		Use:   "wallet-tracker",
		Short: "The Wallet-Tracker Command Line Interface",
		Long:  `Detect real scammers with Wallet-Tracker CLI from anywhere.`,
		RunE:  ShowHelp(os.Stdout),
	}

	// Add subcommands
	rootCmd.AddCommand(
		redis.NewRedisCommand(),
		neodash.NewNeodashCommand(),
		tracker.NewTrackCommand(),
	)

	return rootCmd
}


// ShowHelp shows the command help.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(err)
		d := color.New(color.FgBlue, color.Bold)
		_, _ = d.Print(walletTrackerAscii)
		d = color.New(color.FgGreen, color.Bold)
		_, _ = d.Print(normalOwl)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}

const (
	normalOwl = `
  ___
 (o,o)
 {'"'}
 -"-"-
`
	walletTrackerAscii = `
  __      __        .__  .__          __          ___________                     __                 
/  \    /  \_____  |  | |  |   _____/  |_        \__    ___/___________    ____ |  | __ ___________ 
\   \/\/   /\__  \ |  | |  | _/ __ \   __\  ______ |    |  \_  __ \__  \ _/ ___\|  |/ // __ \_  __ \
 \        /  / __ \|  |_|  |_\  ___/|  |   /_____/ |    |   |  | \// __ \\  \___|    <\  ___/|  | \/
  \__/\  /  (____  /____/____/\___  >__|           |____|   |__|  (____  /\___  >__|_ \\___  >__|   
       \/        \/               \/                                   \/     \/     \/    \/       
`
)
