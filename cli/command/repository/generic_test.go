package repository

import "testing"

func TestCheckWalletNetwork(t *testing.T) {
	tests := []struct {
		wallet   string
		expected int
	}{
		{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F", EthNetwork},
		{"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", BtcNetwork},
		{"3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy", BtcNetwork},
		{"bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfJHX0S", BtcNetwork},
		{"invalid", -1},
	}

	for _, tt := range tests {
		result := CheckWalletNetwork(tt.wallet)
		if result != tt.expected {
			t.Errorf("CheckWalletNetwork(%s) = %d; expected %d", tt.wallet, result, tt.expected)
		}
	}
}
