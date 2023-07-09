package configurations

import (
	_ "embed"
	"os"
)

//go:embed uniswapv3_abi.json
var Uniswapv3ABI string

func GetRPCEndpoint() string {
	return os.Getenv("RPC_ENDPOINT")
}

func GetDatabaseConnectionString() string {
	return os.Getenv("DATABASE_CONNECTION_STRING")
}
