package repository

import (
	"encoding/json"
	"flag"
	"fmt"
	models "github.com/aydinnyunus/wallet-tracker/domain/repository"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type responseJSON struct {
	Op string `json:"op"`
	X  struct {
		LockTime int `json:"lock_time"`
		Ver      int `json:"ver"`
		Size     int `json:"size"`
		Inputs   []struct {
			Sequence int64 `json:"sequence"`
			PrevOut  struct {
				Spent   bool   `json:"spent"`
				TxIndex int    `json:"tx_index"`
				Type    int    `json:"type"`
				Addr    string `json:"addr"`
				Value   int    `json:"value"`
				N       int    `json:"n"`
				Script  string `json:"script"`
			} `json:"prev_out"`
			Script string `json:"script"`
		} `json:"inputs"`
		Time      int    `json:"time"`
		TxIndex   int    `json:"tx_index"`
		VinSz     int    `json:"vin_sz"`
		Hash      string `json:"hash"`
		VoutSz    int    `json:"vout_sz"`
		RelayedBy string `json:"relayed_by"`
		Out       []struct {
			Spent   bool   `json:"spent"`
			TxIndex int    `json:"tx_index"`
			Type    int    `json:"type"`
			Addr    string `json:"addr"`
			Value   int    `json:"value"`
			N       int    `json:"n"`
			Script  string `json:"script"`
		} `json:"out"`
	} `json:"x"`
}

var addr = flag.String("addr", "ws.blockchain.info", "http service address")

func ConnectWebsocketAllTransactions(dbConfig models.Database, args models.ScammerQueryArgs) {
	flag.Parse()
	count := 0
	x := 0
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/inv"}
	color.Blue("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		color.Red("dial:", err)
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			color.Red(err.Error())
		}
	}(c)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				color.Red("read:", err)
				return
			}
			var r responseJSON
			err = json.Unmarshal(message, &r)

			if err != nil {
				color.Red(err.Error())
			}
			color.Green("recv: %s", r.X.Hash)
			if x%1000 == 0 {
				color.Yellow("Imported " + strconv.Itoa(x) + " transactions")
			}
			x += 1

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			btcToUsd := GetBitcoinPrice()
			color.Yellow("Current BTC Price : %f", btcToUsd)
			Hash = r.X.Hash
			tm, err := strconv.ParseInt(strconv.Itoa(r.X.Time), 10, 64)
			if err != nil {
				panic(err)
			}
			Timestamp = time.Unix(tm, 0)

			for j := range r.X.Inputs {
				if len(r.X.Inputs[j].PrevOut.Addr) == 0 {
					continue
				}
				IgnoreAddress = append(IgnoreAddress, r.X.Inputs[j].PrevOut.Addr)
				FromAddress = map[int]map[string]string{
					j + count + r1.Intn(100): {"address": r.X.Inputs[j].PrevOut.Addr, "value": strconv.FormatFloat(float64(r.X.Inputs[j].PrevOut.Value)/SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(float64(r.X.Inputs[j].PrevOut.Value)/SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
				}
			}

			for k := range r.X.Out {
				TotalAmount += float64(r.X.Out[k].Value) / SatoshiToBitcoin
				TotalUSD = TotalAmount * btcToUsd
				ToAddress = map[int]map[string]string{
					k + count + r1.Intn(100): {"address": r.X.Out[k].Addr, "value": strconv.FormatFloat(float64(r.X.Out[k].Value)/SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(float64(r.X.Out[k].Value)/SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
				}

				if !StringInSlice(r.X.Out[k].Addr, IgnoreAddress) {
					FlowBTC += float64(r.X.Out[k].Value) / SatoshiToBitcoin
				}

				FlowUSD = FlowBTC * btcToUsd
			}
			color.Blue("Write Transactions to Neo4j")

			_, err = Neo4jDatabase(Hash, Timestamp.Format("2006-01-02"), strconv.FormatFloat(TotalUSD, 'E', -1, 64), strconv.FormatFloat(TotalAmount, 'E', -1, 64), strconv.FormatFloat(FlowBTC, 'E', -1, 64), strconv.FormatFloat(FlowUSD, 'E', -1, 64), FromAddress, ToAddress)
			if err != nil {
				fmt.Println(err)
				return
			}
			color.Blue("Finished Write Transactions to Neo4j")

			count += 1

		}

	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte("{\"op\": \"unconfirmed_sub\"}"))
			if err != nil {
				fmt.Println(t)
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func ConnectWebsocketSpecificAddress(btcAddress string) {
	flag.Parse()
	count, x := 0, 0
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/inv"}
	color.Blue("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		color.Red("dial:", err)
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			color.Red(err.Error())
		}
	}(c)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				color.Red("read:", err)
				return
			}
			var r responseJSON
			err = json.Unmarshal(message, &r)

			if err != nil {
				color.Red(err.Error())
			}
			color.Green("recv: %s", r.X.Hash)
			if x%1000 == 0 {
				color.Yellow("Imported " + strconv.Itoa(x) + " transactions")
			}
			x += 1

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			btcToUsd := GetBitcoinPrice()
			color.Yellow("Current BTC Price : %f", btcToUsd)
			Hash = r.X.Hash
			tm, err := strconv.ParseInt(strconv.Itoa(r.X.Time), 10, 64)
			if err != nil {
				panic(err)
			}
			Timestamp = time.Unix(tm, 0)

			for j := range r.X.Inputs {
				if len(r.X.Inputs[j].PrevOut.Addr) == 0 {
					continue
				}
				IgnoreAddress = append(IgnoreAddress, r.X.Inputs[j].PrevOut.Addr)
				FromAddress = map[int]map[string]string{
					j + count + r1.Intn(100): {"address": r.X.Inputs[j].PrevOut.Addr, "value": strconv.FormatFloat(float64(r.X.Inputs[j].PrevOut.Value)/SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(float64(r.X.Inputs[j].PrevOut.Value)/SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
				}
			}

			for k := range r.X.Out {
				TotalAmount += float64(r.X.Out[k].Value) / SatoshiToBitcoin
				TotalUSD = TotalAmount * btcToUsd
				ToAddress = map[int]map[string]string{
					k + count + r1.Intn(100): {"address": r.X.Out[k].Addr, "value": strconv.FormatFloat(float64(r.X.Out[k].Value)/SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(float64(r.X.Out[k].Value)/SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
				}

				if !StringInSlice(r.X.Out[k].Addr, IgnoreAddress) {
					FlowBTC += float64(r.X.Out[k].Value) / SatoshiToBitcoin
				}

				FlowUSD = FlowBTC * btcToUsd
			}
			color.Blue("Write Transactions to Neo4j")

			_, err = Neo4jDatabase(Hash, Timestamp.Format("2006-01-02"), strconv.FormatFloat(TotalUSD, 'E', -1, 64), strconv.FormatFloat(TotalAmount, 'E', -1, 64), strconv.FormatFloat(FlowBTC, 'E', -1, 64), strconv.FormatFloat(FlowUSD, 'E', -1, 64), FromAddress, ToAddress)
			if err != nil {
				fmt.Println(err)
				return
			}
			color.Blue("Finished Write Transactions to Neo4j")

			count += 1

		}

	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			address := "{\"op\": \"addr_sub\",\"addr\": " + btcAddress + "}"
			err := c.WriteMessage(websocket.TextMessage, []byte(address))
			if err != nil {
				fmt.Println(t)
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
