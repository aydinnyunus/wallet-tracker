package repository

import "time"

const (
	BtcNetwork  = 1
	EthNetwork  = 0
	importQuery = "MERGE (t:Transaction {hash:$hash}) CREATE (t2:Transaction) SET t.hash = $hash, t.timestamp = datetime($timestamp), " +
		"t.totalBTC = toFloat($total_amount), t.totalUSD = toFloat($total_usd),t.flowBTC = toFloat($flow_btc), " +
		"t.flowUSD = toFloat($flow_usd) MERGE (a:Address {id:$from_address}) " +
		"CREATE (a)-[:SENT {valueBTC:toFloat($from_value), valueUSD: toFloat($from_value_usd)}]->(t) " +
		"MERGE (b:Address {id:$to_address}) CREATE (t)-[:SENT {valueBTC: toFloat($to_value), valueUSD: toFloat($to_value_usd)}]->(b)"
)

var (
	Neo4jUri         = getEnv("NEO4J_URI", "neo4j://localhost:7687")
	Neo4jUser        = getEnv("NEO4J_USER", "neo4j")
	Neo4jPass        = getEnv("NEO4J_PASS", "letmein")
	exchange         = []string{"uniswap", "bitfinex"}
	SatoshiToBitcoin = float64(100000000)
	ExitNodes        []string
	IgnoreAddress    []string
	FromAddress      = map[int]map[string]string{}
	ToAddress        = map[int]map[string]string{}
	Hash             = ""
	TotalAmount      = float64(0)
	Timestamp        = time.Now()
	TotalUSD         = float64(0)
	FlowBTC          = float64(0)
	FlowUSD          = float64(0)
)
