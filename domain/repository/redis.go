package repository

import (
	"fmt"
)

/*
	Created by aomerk at 5/21/21 for project cli
*/

// global constants for file
const ()

// global variables (not cool) for this file
var ()

// ScammerQueryArgs are arguments you can use to customize your queries. Multiple fields can be used at once,
// also empty query args is not a problem.
type ScammerQueryArgs struct {
	// how many results do you want to retrieve
	Limit int

	Exchanges []string

	Port int

	Wallet []string

	Network string

	Detect bool

	// If true, queries nested fields like request headers.
	Verbose bool
}

func (q ScammerQueryArgs) String() string {
	var query string

	query = fmt.Sprintf("%s\nDisplaying maximum %d rows", query, q.Limit)

	if q.Exchanges != nil && len(q.Exchanges) != 0 {
		query = fmt.Sprintf(
			"%s\nQuerying %d Exchanges with names: %s", query, len(q.Exchanges), q.Exchanges,
		)
	}

	return query
}
