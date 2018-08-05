package ethclient

// Transaction represents a structure of a transaction in ethereum network
type Transaction struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Gas       string `json:"gas"`
	GasPrice  string `json:"gasPrice"`
	Value     string `json:"value"`
	Data      string `json:"data"`
	BlockHash string `json:"blockHash"`
	BlockNum  string `json:"blockNumber"`
	Hash      string `json:"hash"`
	Nonce     string `json:"nonce"`
}
