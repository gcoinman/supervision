package eth_domain

import (
	"strconv"
)

// TransactionReceipt represents a receipt of a tx
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
	return num == 1
}

func dropHexPrefix(s string) string {
	if s[:2] == "0x" {
		return s[2:]
	}
	return s
}
