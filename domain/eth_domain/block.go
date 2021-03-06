package eth_domain

// Block represents a block in Ethereum
type Block struct {
	Hash         string `json:"hash"`
	Nonce        string `json:"nonce"`
	Timestamp    string `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

// BlockResponse represents an response of eth_getBlockByBlockNumber
type BlockResponse struct {
	Error *JSONRPCError `json:"error"`
	Block *Block        `json:"result"`
}
