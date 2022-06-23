package repository

import (
	"github.com/fatih/color"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Neo4jDatabase(hash, timestamp, totalUSD, totalAmount, flowBTC, flowUSD string, from, toAddress map[int]map[string]string) (string, error) {
	driver, err := neo4j.NewDriver(Neo4jUri, neo4j.BasicAuth(Neo4jUser, Neo4jPass, ""))
	if err != nil {
		color.Red(err.Error())
		return "", err
	}
	defer func(driver neo4j.Driver) {
		err := driver.Close()
		if err != nil {
			color.Red(err.Error())
		}
	}(driver)

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			color.Red(err.Error())
		}
	}(session)
	for _, value := range from {
		for _, value2 := range toAddress {
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				result, err := transaction.Run(
					importQuery,
					map[string]interface{}{"hash": hash, "timestamp": timestamp, "total_amount": totalAmount, "total_usd": totalUSD, "from_address": value["address"], "to_address": value2["address"], "flow_btc": flowBTC, "flow_usd": flowUSD, "from_value": value["value"], "from_value_usd": value["value_usd"], "to_value": value2["value"], "to_value_usd": value2["to_value_usd"]})

				if err != nil {
					color.Red(err.Error())
					return nil, err
				}

				if result.Next() {
					return result.Record().Values[0], nil
				}

				return nil, result.Err()
			})
			if err != nil {
				color.Red(err.Error())
				return "", err
			}

			return "", nil
		}

		return "", nil
	}
	return "", nil
}
