package tracker

import (
	"fmt"
	"github.com/aydinnyunus/blockchain"
	"github.com/aydinnyunus/wallet-tracker/cli/command/repository"
	models "github.com/aydinnyunus/wallet-tracker/domain/repository"
	"github.com/fatih/color"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// global constants for file
const ()

// global variables (not cool) for this file
var ()

func TrackCommand() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "track",
		Short: "Command to track Wallet",
		Long:  ``,
		RunE:  startTrack,
	}

	// declaring local flags used by track wallet commands.
	getCmd.Flags().String(
		"wallet", "w", "Specify specific wallet",
	)

	getCmd.Flags().String(
		"network", "n", "Specify specific network",
	)

	getCmd.Flags().BoolP(
		"detect-exchanges", "d", false, "Detect Exchange Exits",
	)

	getCmd.Flags().BoolP(
		"verbose", "v", true, "To hide field values like headers, say --verbose=false",
	)

	return getCmd
}

// startTrack implements GetCommand logic.
func startTrack(cmd *cobra.Command, _ []string) error {
	var (
		cliConfig models.Database
		err       error
	)

	// get cli config for authentication
	err = viper.Unmarshal(&cliConfig)
	if err != nil {
		return err
	}

	queryArgs := models.ScammerQueryArgs{Limit: 1}

	// parse and set list of suspects to be queried
	wallet, err := cmd.Flags().GetString("wallet")
	if err != nil {
		return err
	} else if len(wallet) > 0 {
		queryArgs.Wallet = append(queryArgs.Wallet, wallet)
	}

	network, err := cmd.Flags().GetString("network")
	if err != nil {
		return err
	} else if len(network) > 1 {
		if repository.CheckWalletNetwork(wallet) == 1 {
			if network != "BTC" {
				color.Red("Please enter correct Network")
				return nil
			}
		} else if repository.CheckWalletNetwork(wallet) == 0 {
			if network != "ETH" {
				color.Red("Please enter correct Network")
				return nil
			}
		}
		queryArgs.Network = network
	}

	// parse verboseness flag.
	detect, err := cmd.Flags().GetBool("detect-exchanges")
	if err != nil {
		return err
	} else if detect {
		queryArgs.Detect = detect
	}

	// parse verboseness flag.
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	} else if verbose {
		queryArgs.Verbose = verbose
	}

	out, err := TrackWallet(cliConfig, queryArgs)
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

func TrackWallet(dbConfig models.Database, args models.ScammerQueryArgs) ([]byte, error) {

	count := 0
	walletID := args.Wallet

	for i := range walletID {
		network := repository.CheckWalletNetwork(walletID[i])
		graph := repository.New()

		if network == repository.BtcNetwork {

			c, e := blockchain.New()

			resp, e := c.GetAddress(walletID[i])
			if e != nil {
				time.Sleep(5 * time.Second)
			}
			count = 0
			node0 := graph.AddNode(resp.Address, resp.FinalBalance)

			for {
				//connectWebsocketAllTransactions()
				//connectWebsocketSpecificAddress(walletID)
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				for i := range resp.Txs {
					btcToUsd := repository.GetBitcoinPrice()
					color.Yellow("Current BTC Price : %f", btcToUsd)
					repository.Hash = resp.Txs[i].Hash
					color.Yellow("Currently working on %s", repository.Hash)
					tm, err := strconv.ParseInt(strconv.Itoa(resp.Txs[i].Time), 10, 64)
					if err != nil {
						panic(err)
					}
					repository.Timestamp = time.Unix(tm, 0)

					//fmt.Println(resp.Txs[i].Result)

					for j := range resp.Txs[i].Inputs {
						if len(resp.Txs[i].Inputs[j].PrevOut.Addr) == 0 {
							continue
						}
						repository.IgnoreAddress = append(repository.IgnoreAddress, resp.Txs[i].Inputs[j].PrevOut.Addr)
						repository.FromAddress = map[int]map[string]string{
							j + count + r1.Intn(100): {"address": resp.Txs[i].Inputs[j].PrevOut.Addr, "value": strconv.FormatFloat(float64(resp.Txs[i].Inputs[j].PrevOut.Value)/repository.SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(float64(resp.Txs[i].Inputs[j].PrevOut.Value)/repository.SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
						}
						//fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Addr)
						node1 := graph.AddNode(resp.Txs[i].Inputs[j].PrevOut.Addr, resp.Txs[i].Inputs[j].PrevOut.Value)
						graph.AddEdge(node0, node1, 1)

						//fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Value)
					}

					for k := range resp.Txs[i].Out {
						//fmt.Println(resp.Txs[i].Out[k].Addr)
						//fmt.Println(resp.Txs[i].Out[k].Value)
						repository.TotalAmount += float64(resp.Txs[i].Out[k].Value) / repository.SatoshiToBitcoin
						repository.TotalUSD = repository.TotalAmount * btcToUsd
						repository.ToAddress = map[int]map[string]string{
							k + count + r1.Intn(100): {"address": resp.Txs[i].Out[k].Addr, "value": strconv.FormatFloat(float64(resp.Txs[i].Out[k].Value)/repository.SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(float64(resp.Txs[i].Out[k].Value)/repository.SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
						}

						if !repository.StringInSlice(resp.Txs[i].Out[k].Addr, repository.IgnoreAddress) {
							repository.FlowBTC += float64(resp.Txs[i].Out[k].Value) / repository.SatoshiToBitcoin
						}

						repository.FlowUSD = repository.FlowBTC * btcToUsd
						color.Blue("Write Transactions to Neo4j")
						_, err := repository.Neo4jDatabase(repository.Hash, repository.Timestamp.Format("2006-01-02"), strconv.FormatFloat(repository.TotalUSD, 'E', -1, 64), strconv.FormatFloat(repository.TotalAmount, 'E', -1, 64), strconv.FormatFloat(repository.FlowBTC, 'E', -1, 64), strconv.FormatFloat(repository.FlowUSD, 'E', -1, 64), repository.FromAddress, repository.ToAddress)
						if err != nil {
							color.Red(err.Error())
							return nil, nil
						}
						color.Blue("Finished writing Transactions to Neo4j")
						node1 := graph.AddNode(resp.Txs[i].Out[k].Addr, resp.Txs[i].Out[k].Value)
						graph.AddEdge(node0, node1, 1)
					}
				}
				count += 1

				color.Yellow("Currently Working on %s", graph.Nodes[count].WalletId)
				resp, e := c.GetAddress(graph.Nodes[count].WalletId)
				if e != nil {
					fmt.Print(e)
					time.Sleep(1 * time.Minute)
				}
				node0 = graph.AddNode(resp.Address, resp.FinalBalance)

				if len(graph.Nodes[len(graph.Nodes)-1].Edges) == 0 {
					repository.ExitNodes = append(repository.ExitNodes, graph.Nodes[len(graph.Nodes)-1].WalletId)
					break
				}

			}
		} else if network == repository.EthNetwork {
			c, e := blockchain.New()

			resp2, e := c.GetETHAddressSummary(walletID[i], true)
			if e != nil {
				color.Red(e.Error())
				time.Sleep(5 * time.Second)
			}
			count = 0
			balance, _ := strconv.Atoi(resp2.Balance)
			node0 := graph.AddNode(resp2.Hash, balance)

			resp, e := c.GetETHAddress(walletID[i])
			if e != nil {
				color.Red(e.Error())
				time.Sleep(5 * time.Second)
			}

			for {
				//connectWebsocketAllTransactions()
				//connectWebsocketSpecificAddress(walletID)
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				for i := range resp.Transactions {
					btcToUsd := repository.GetBitcoinPrice()
					color.Yellow("Current BTC Price : %d", btcToUsd)

					repository.Hash = resp.Transactions[i].Hash
					color.Yellow("Currently working on %s", repository.Hash)

					tm, err := strconv.ParseInt(resp.Transactions[i].Timestamp, 10, 64)
					if err != nil {
						color.Red(err.Error())
						panic(err)
					}
					repository.Timestamp = time.Unix(tm, 0)

					//fmt.Println(resp.Txs[i].Result)

					if resp.Transactions[i].From == resp.Transactions[i].To {
						repository.IgnoreAddress = append(repository.IgnoreAddress, resp.Transactions[i].From)
						continue
					}
					value, _ := strconv.ParseFloat(resp.Transactions[i].Value, 64)
					repository.FromAddress = map[int]map[string]string{
						i + count + r1.Intn(100): {"address": resp.Transactions[i].From, "value": strconv.FormatFloat(value/repository.SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(value/repository.SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
					}
					//fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Addr)
					node1 := graph.AddNode(resp.Transactions[i].From, int(value))
					graph.AddEdge(node0, node1, 1)

					//fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Value)

					//fmt.Println(resp.Txs[i].Out[k].Addr)
					//fmt.Println(resp.Txs[i].Out[k].Value)
					repository.TotalAmount += value / repository.SatoshiToBitcoin
					repository.TotalUSD = repository.TotalAmount * btcToUsd
					value, _ = strconv.ParseFloat(resp.Transactions[i].Value, 64)
					repository.ToAddress = map[int]map[string]string{
						i + count + r1.Intn(100): {"address": resp.Transactions[i].To, "value": strconv.FormatFloat(value/repository.SatoshiToBitcoin, 'E', -1, 64), "value_usd": strconv.FormatFloat(value/repository.SatoshiToBitcoin*btcToUsd, 'E', -1, 64)},
					}

					if !repository.StringInSlice(resp.Transactions[i].From, repository.IgnoreAddress) {
						repository.FlowBTC += value / repository.SatoshiToBitcoin
					}

					repository.FlowUSD = repository.FlowBTC * btcToUsd
					color.Blue("Write Transactions to Neo4j")

					_, err = repository.Neo4jDatabase(repository.Hash, repository.Timestamp.Format("2006-01-02"), strconv.FormatFloat(repository.TotalUSD, 'E', -1, 64), strconv.FormatFloat(repository.TotalAmount, 'E', -1, 64), strconv.FormatFloat(repository.FlowBTC, 'E', -1, 64), strconv.FormatFloat(repository.FlowUSD, 'E', -1, 64), repository.FromAddress, repository.ToAddress)
					if err != nil {
						color.Red(err.Error())
						return nil, nil
					}
					color.Blue("Finished writing Transactions to Neo4j")

					node1 = graph.AddNode(resp.Transactions[i].From, int(value))
					graph.AddEdge(node0, node1, 1)

				}
				count += 1

				resp, e := c.GetETHAddressSummary(graph.Nodes[count].WalletId, true)
				if e != nil {
					fmt.Print(e)
					color.Red("Blockchain.info Rate Limiting waiting for 1 minute")
					time.Sleep(1 * time.Minute)
				}
				balance, _ = strconv.Atoi(resp.Balance)
				node0 = graph.AddNode(resp.Hash, balance)

				if len(graph.Nodes[len(graph.Nodes)-1].Edges) == 0 {
					repository.ExitNodes = append(repository.ExitNodes, graph.Nodes[len(graph.Nodes)-1].WalletId)
					break
				}

			}
		}
	}

	if args.Detect {
		rdb, ctx, err := repository.ConnectToRedis(dbConfig)
		if err != nil {
			log.Fatal(err)
		}
		uni, bitfinex := repository.DetectExchanges(rdb, ctx)

		for i := range repository.ExitNodes {
			color.Blue("Checking Exchange Exits")
			if repository.StringInSlice(repository.ExitNodes[i], uni) {
				color.Green("This address %s refers to Uniswap\n", repository.ExitNodes[i])
			} else if repository.StringInSlice(repository.ExitNodes[i], bitfinex) {
				color.Green("This address %s refers to Bitfinex\n", repository.ExitNodes[i])
			}
		}
	}

	color.Blue("You can visualize the all data using : ./wallet-tracker neodash start")
	return nil, nil
}
