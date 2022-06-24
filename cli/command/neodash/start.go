package neodash

import (
	"github.com/aydinnyunus/wallet-tracker/domain/cli"
	"github.com/fatih/color"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strconv"

	models "github.com/aydinnyunus/wallet-tracker/domain/repository"
	"github.com/go-git/go-git/v5"
)

// global constants for file
const ()

// global variables (not cool) for this file
var (mydir, _ = os.Getwd())

func StartCommand() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "start",
		Short: "Command to start NeoDash",
		Long: ``,
		RunE: startNeoDash,
	}

	// declaring local flags used by start neodash commands.
	getCmd.Flags().String(
		"port", "7487", "Specify specific port",
	)

	getCmd.Flags().BoolP(
		"verbose", "v", true, "To hide field values like headers, say --verbose=false",
	)

	return getCmd
}

// startNeoDash implements GetCommand logic.
func startNeoDash(cmd *cobra.Command, _ []string) error {
	var (
		cliConfig cli.Cli
		err       error
	)

	// get cli config for authentication
	err = viper.Unmarshal(&cliConfig)
	if err != nil {
		return err
	}

	queryArgs := models.ScammerQueryArgs{Limit: 1}

	// parse and set list of suspects to be queried
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		return err
	} else if len(port) > 0 {
		queryArgs.Port, err = strconv.Atoi(port)
		if err != nil {
			return err
		}
	}

	// parse verboseness flag.
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	} else if verbose {
		queryArgs.Verbose = verbose
	}

	out, err := CreateNeodash(queryArgs)
	if err != nil {
		log.Fatal(err)
	}

	// print out query settings
	color.Blue(queryArgs.String())


	// print out result
	_, err = pp.Print(out)
	if err != nil {
		return err
	}

	return nil
}

func CreateNeodash(args models.ScammerQueryArgs) ([]byte, error) {
	_, err := git.PlainClone(mydir + "/bitcoin-to-neo4jdash", false, &git.CloneOptions{
		URL:      "https://github.com/aydinnyunus/bitcoin-to-neo4jdash",
		Progress: os.Stdout,
	})

	if err != nil{
		log.Fatal(err)
		return nil, err
	}

	color.Yellow("Define Schema is starting")
	out, err := DefineSchema(args)
	if err != nil {
		return nil, err
	}
	color.Yellow("Define Schema is finished")


	color.Yellow("Builds, (re)creates, starts, and attaches to containers for a service.")

	out, err = DockerComposeUp(args)
	if err != nil {
		return nil, err
	}
	color.Yellow("docker-compose up finished. http://localhost:80")

	return out, err
}

// DefineSchema is a temporary method to satisfy the authentication process.
func DefineSchema(args models.ScammerQueryArgs) ([]byte, error) {
	cmdStr := "sudo sh " + mydir + "/bitcoin-to-neo4jdash/define_schema.sh"
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	return out, nil
}

func DockerComposeUp(args models.ScammerQueryArgs)([]byte, error){
	cmdStr := "sudo docker-compose -f " + mydir + "/bitcoin-to-neo4jdash/docker-compose.yml up"
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()

	return out, nil
}
