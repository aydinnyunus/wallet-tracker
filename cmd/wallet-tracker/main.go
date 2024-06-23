package main

import (
	generic "github.com/aydinnyunus/wallet-tracker/cli/command/repository"
	"github.com/fsnotify/fsnotify"
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

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Define the .env file to watch
	envFile := ".env"

	// Start watching the .env file for changes
	err = watcher.Add(envFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Watching .env file for changes...")

	// Start a goroutine to handle events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println(".env file modified. Restarting Docker Compose...")
					generic.RestartDockerCompose()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	// Block main goroutine to keep watching
	select {}

	rootCmd = commands.NewWalletTrackerCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
