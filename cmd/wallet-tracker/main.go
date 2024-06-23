package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/aydinnyunus/wallet-tracker/cli/command/commands"
)

var rootCmd *cobra.Command

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rootCmd = commands.NewWalletTrackerCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
