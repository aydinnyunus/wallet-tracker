package repository

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func getEnv(key, fallback string) string {
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
