package main

import (
	generic "github.com/aydinnyunus/wallet-tracker/cli/command/repository"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/aydinnyunus/wallet-tracker/cli/command/commands"
)

var rootCmd *cobra.Command

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Create a new file watcher
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

	// Start a goroutine to handle file system events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println(".env file modified. Restarting Docker Compose...")
					err := godotenv.Load() // Reload the .env file
					if err != nil {
						log.Printf("Error reloading .env file: %v", err)
					}
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

	// Handle interrupts to gracefully shut down
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Initialize and execute the root command
	rootCmd := commands.NewWalletTrackerCommand()
	go func() {
		if err := rootCmd.Execute(); err != nil {
			log.Fatalf("Error executing root command: %v", err)
		}
	}()

	// Block the main goroutine until an interrupt signal is received
	<-stop
	log.Println("Shutting down...")
}
