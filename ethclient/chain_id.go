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
)

// URL generates an url based on chain id with a provided api key
func (id ChainID) URL(apiKey string) string {
	switch id {
	case Mainnet:
		return fmt.Sprintf("https://mainnet.infura.io/%s", apiKey)
	case Ropsten:
		return fmt.Sprintf("https://ropsten.infura.io/%s", apiKey)
	default:
		return ""
	}
}
