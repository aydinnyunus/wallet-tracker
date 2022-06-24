package tracker

import (
	"github.com/aydinnyunus/wallet-tracker/cli/command/repository"
	models "github.com/aydinnyunus/wallet-tracker/domain/repository"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// global constants for file
const ()

// global variables (not cool) for this file
var ()

func TrackWebsocketCommand() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "websocket",
		Short: "Command to track Wallet",
		Long:  ``,
		RunE:  startTrackWebsocket,
	}

	// declaring local flags used by get trip commands.
	getCmd.Flags().String(
		"wallet", "w", "Specify specific wallet",
	)

	getCmd.Flags().String(
		"network", "n", "Specify specific network",
	)

	getCmd.Flags().BoolP(
		"verbose", "v", true, "To hide field values like headers, say --verbose=false",
	)

	getCmd.Flags().BoolP(
		"all", "a", true, "Get all transactions on blockchain.com",
	)

	return getCmd
}

// startTrack implements GetCommand logic.
func startTrackWebsocket(cmd *cobra.Command, _ []string) error {
	var (
		cliConfig models.Database
		err       error
	)

	// get cli config for authentication
	err = viper.Unmarshal(&cliConfig)
	if err != nil {
		return err
	}

	queryArgs := models.ScammerQueryArgs{Limit: 1}

	// parse and set list of suspects to be queried
	wallet, err := cmd.Flags().GetString("wallet")
	if err != nil {
		return err
	} else if len(wallet) > 0 {
		queryArgs.Wallet = append(queryArgs.Wallet, wallet)
	}

	all, err := cmd.Flags().GetBool("all")
	if err != nil {
		color.Red(err.Error())
		return err
	}

	// parse verboseness flag.
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	} else if verbose {
		queryArgs.Verbose = verbose
	}

	if all{
		repository.ConnectWebsocketAllTransactions(cliConfig, queryArgs)
	} else{
		/*TODO STRING ARRAYI OLACAK*/
		repository.ConnectWebsocketSpecificAddress(queryArgs.Wallet[0])
	}

	// print out query settings
	color.Blue(queryArgs.String())



	return nil
}