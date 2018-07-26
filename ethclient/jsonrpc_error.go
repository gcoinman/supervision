package ethclient

// JSONRPCError represents an error in jsonrpc request
type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
