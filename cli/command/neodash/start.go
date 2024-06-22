package neodash

import (
	"fmt"
	"github.com/aydinnyunus/wallet-tracker/domain/cli"
	models "github.com/aydinnyunus/wallet-tracker/domain/repository"
	"github.com/fatih/color"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// global constants for file
const ()

// global variables (not cool) for this file
var (
	mydir, _ = os.Getwd()
)

func StartCommand() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "start",
		Short: "Command to start NeoDash",
		Long:  ``,
		RunE:  startNeoDash,
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
	color.Yellow("Define Schema is starting")
	_, err := DefineSchema(args)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	color.Yellow("Define Schema is finished")

	err = checkBinaryExists("docker")
	// Check if Docker Compose is running
	if err == nil {
		cmd := exec.Command("docker", "compose", "ps", "-q")
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		if len(out) == 0 {
			// Docker Compose is not running, start it
			color.Yellow("Builds, (re)creates, starts, and attaches to containers for a service.")
			out, err = DockerComposeUp(args)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			color.Yellow("docker-compose up finished. http://localhost:80")
		} else {
			color.Yellow("Docker Compose is already running")
		}

		return out, err
	} else {
		color.Red("Docker is not installed.")
		return nil, err
	}

}

// DefineSchema is a temporary method to satisfy the authentication process.
func DefineSchema(args models.ScammerQueryArgs) ([]byte, error) {
	cmdStr := "sudo sh define_schema.sh"
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	color.Yellow(string(out))
	return out, nil
}

func DockerComposeUp(args models.ScammerQueryArgs) ([]byte, error) {
	cmdStr := "sudo docker compose -f docker-compose.yml up -d"
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	color.Yellow(string(out))
	return out, nil
}

func checkBinaryExists(binaryName string) error {
	_, err := exec.LookPath(binaryName)
	if err != nil {
		return fmt.Errorf("%s binary not found", binaryName)
	}
	return nil
}
