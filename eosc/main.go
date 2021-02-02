package main

import (
	// Load all contracts here, so we can always read and decode
	// transactions with those contracts.
	_ "github.com/FSharesSaaS/fshares.fsgo/msig"
	_ "github.com/FSharesSaaS/fshares.fsgo/system"
	_ "github.com/FSharesSaaS/fshares.fsgo/token"

	"github.com/eoscanada/eosc/eosc/cmd"
)

var version = "dev"

func init() {
	cmd.Version = version
}

func main() {
	cmd.Execute()
}
