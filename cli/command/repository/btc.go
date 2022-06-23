package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Usd struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"USD"`
}

func checkBTCFormat(walletAddr string) bool {
	if strings.HasPrefix(walletAddr, "1") || strings.HasPrefix(walletAddr, "3") || strings.HasPrefix(walletAddr, "bc1") {
		return true
	}
	return false
}

func GetBitcoinPrice() float64 {
	resp, err := http.Get("https://blockchain.info/ticker")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return result.Usd.Last
}