package main

import (
	"fmt"
	generic "github.com/aydinnyunus/wallet-tracker/cli/command/repository"
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

	neo4jUser := generic.GetEnv("NEO4J_USERNAME", "neo4j")
	neo4jPass := generic.GetEnv("NEO4J_PASS", "letmein")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dockerEnvVarValue, err := generic.GetDockerEnvVar(generic.ContainerName, generic.Neo4jAuth)
	if err != nil {
		log.Fatalf("Error getting Docker env var: %v", err)
	}

	envVarValue := neo4jUser + "/" + neo4jPass

	if envVarValue == dockerEnvVarValue {
		fmt.Println("The .env NEO4J_AUTH value matches the Docker container NEO4J_AUTH value.")
	} else {
		generic.RestartDockerCompose()
	}

	rootCmd = commands.NewWalletTrackerCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
