package eth_domain

// Tx represents a structure of a tx in ethereum network
type Tx struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Gas       string `json:"gas"`
	GasPrice  string `json:"gasPrice"`
	Value     string `json:"value"`
	Data      string `json:"input"`
	BlockHash string `json:"blockHash"`
	BlockNum  string `json:"blockNumber"`
	Hash      string `json:"hash"`
	Nonce     string `json:"nonce"`
}
