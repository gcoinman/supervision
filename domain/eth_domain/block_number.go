package eth_domain

// BlockNumberResponse represents a response of eth_getBlockNumber request.
type BlockNumberResponse struct {
	Error    *JSONRPCError `json:"error"`
	BlockNum string        `json:"result"`
}
