package redis

import (
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/aydinnyunus/wallet-tracker/cli/command/repository"
	models "github.com/aydinnyunus/wallet-tracker/domain/repository"
)


// global constants for file
const ()

// global variables (not cool) for this file
var ()


func GetCommand() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Command to query exchanges",
		Long:  `Command to query exchanges on your Redis Database`,
		RunE:  getRedisCmd,
	}

	// declaring local flags used by get trip commands.
	getCmd.Flags().StringSliceP(
		"exchanges", "e", nil, "Comma separated values of Exchange names. "+
			"--exchanges exchange-name,exchange-name2",
	)

	getCmd.Flags().IntP("limit", "l", 5, "Max number of trips you want to be displayed --limit 5")

	getCmd.Flags().BoolP(
		"verbose", "v", true, "To hide field values like headers, say --verbose=false",
	)

	return getCmd
}

// getRedisCmd implements GetCommand logic.
func getRedisCmd(cmd *cobra.Command, _ []string) error {
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
	exchanges, err := cmd.Flags().GetStringSlice("exchanges")
	if err != nil {
		return err
	} else if len(exchanges) > 0 {
		queryArgs.Exchanges = exchanges
	}

	// parse max limit of rows displayed.
	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		return err

	} else if limit > 0 {
		queryArgs.Limit = limit
	}

	// parse verboseness flag.
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	} else if limit > 0 {
		queryArgs.Verbose = verbose
	}

	// get exchanges with parsed query arguments for this subcommand
	exchange, err := get(cliConfig, queryArgs)
	if err != nil {
		return err
	}

	// print out query settings
	color.Blue(queryArgs.String())

	if exchange == nil || len(exchange) == 0 {
		color.Blue("0 result.")
		return nil
	}

	// print out result
	_, err = pp.Print(exchange)
	if err != nil {
		return err
	}

	return nil
}


// ScammerQueryArgs for more information about existing filters.
func get(dbConfig models.Database, args models.ScammerQueryArgs) ([]string, error) {
	var (
		err    error
		_     *redis.Client
		result []string
	)
	// connect to database
	rdb, ctx, err := repository.ConnectToRedis(dbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "can not establish connection to agent database")
	}

	// nobody wants to retrieve all hundreds of thousands of results.
	if args.Limit == 0 {
		args.Limit = 10
	}

	// filter by suspect ids
	if args.Exchanges != nil {
		for i,_ := range args.Exchanges{
			result = append(result, repository.ReadRedis(rdb,ctx,args.Exchanges[i], args.Limit)...)
		}
	}

	return result, err
}
