package ethclient

import (
	"fmt"
)

// ChainID represents a chain id of ethereum network.
type ChainID int64

const (
	// Mainnet chain id
	Mainnet = 1

	// Ropsten chain id
	Ropsten = 3

	// Localhost chain id
	Localhost = 0

	// Getho chain id
	Getho = 1010
)

// URL generates an url based on chain id with a provided api key
func (id ChainID) URL(apiKey string) string {
	switch id {
	case Mainnet:
		return fmt.Sprintf("https://mainnet.infura.io/%s", apiKey)
	case Ropsten:
		return fmt.Sprintf("https://ropsten.infura.io/%s", apiKey)
	case Localhost:
		return "http://localhost:7545"
	case Getho:
		return "https://tame-kitten-68170.getho.io/jsonrpc"
	default:
		return ""
	}
}
