package ethclient

import "strconv"

// TransactionReceipt represents a receipt of a transaction
type TransactionReceipt struct {
	BlockNum string `json:"blockNumber"`
	Hash     string `json:"blockHash"`
	From     string `json:"from"`
	To       string `json:"to"`
	Status   string `json:"status"`
}

// TransactionReceiptResponse represents a response of eth_getTransactionReceipt
type TransactionReceiptResponse struct {
	Error   *JSONRPCError       `json:"error"`
	Receipt *TransactionReceipt `json:"result"`
}

// IsSuccess returns if the receipt's status is success or not
func (receipt *TransactionReceipt) IsSuccess() bool {
	num, err := strconv.ParseInt(dropHexPrefix(receipt.Status), 10, 64)
	if err != nil {
		return false
	}
	if num == 0 {
		return true
	} else {
		return false
	}
}
