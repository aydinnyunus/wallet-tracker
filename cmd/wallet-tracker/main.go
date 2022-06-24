package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/aydinnyunus/wallet-tracker/cli/command/commands"
)

var rootCmd *cobra.Command

func main() {
	rootCmd = commands.NewWalletTrackerCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
