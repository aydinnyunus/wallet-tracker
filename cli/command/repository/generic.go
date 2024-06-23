package repository

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func CheckWalletNetwork(walletAddr string) int {
	if len(walletAddr) == 42 && strings.HasPrefix(walletAddr, "0x") {
		return EthNetwork // ETH
	} else if len(walletAddr) > 25 && len(walletAddr) < 36 && checkBTCFormat(walletAddr) {
		return BtcNetwork // BTC
	} else if len(walletAddr) == 42 && strings.HasPrefix(walletAddr, "bc1") {
		return BtcNetwork // BTC
	}
	return -1 // Others
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Function to restart Docker Compose
func RestartDockerCompose() {
	commands := []string{
		"docker compose down",
		"docker compose up -d",
	}

	for _, cmd := range commands {
		command := exec.Command("sh", "-c", cmd)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			log.Printf("Command failed: %s\n", cmd)
			log.Fatal(err)
		}
	}
}

func GetDockerEnvVar(containerName, key string) (string, error) {
	cmd := exec.Command("docker", "exec", containerName, "printenv", key)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
