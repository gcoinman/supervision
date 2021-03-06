package ethclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-Technologies/supervision/domain/eth_domain"
	"github.com/pkg/errors"
)

// EthAPIClient is an interface of ethereum apis
type EthAPIClient interface {
	GetBlockNumber() (int64, error)
	GetBlockByBlockNumber(bnum int64) (*eth_domain.Block, error)
	GetTransactionReceipt(client *http.Client, txhash string) (*eth_domain.TransactionReceipt, error)
}

// GetBlockNumber fetches the latest block number in ethereum network
func (c *EthClient) GetBlockNumber(client *http.Client) (int64, error) {
	raw, err := c.do(client, "eth_blockNumber", []string{})
	if err != nil {
		return 0, err
	}

	resp := new(eth_domain.BlockNumberResponse)
	if err := json.Unmarshal(raw, resp); err != nil {
		return 0, errors.Wrap(err, "failed to fetch blocknumber")
	}

	if resp.Error != nil {
		return 0, errors.New(resp.Error.Message)
	}

	num, err := strconv.ParseInt(dropHexPrefix(resp.BlockNum), 16, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert string to int")
	}

	return num, nil
}

// GetBlockByBlockNumber fetches the block with a specified block number.
func (c *EthClient) GetBlockByBlockNumber(client *http.Client, bnum int64, isFullBlock bool) (*eth_domain.Block, error) {
	raw, err := c.do(client, "eth_getBlockByNumber", []interface{}{toHex(bnum), isFullBlock})
	if err != nil {
		return nil, err
	}

	resp := new(eth_domain.BlockResponse)
	if err := json.Unmarshal(raw, resp); err != nil {
		return nil, errors.Wrap(err, "failed to fetch a block")
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	return resp.Block, nil
}

// GetTransactionReceipt fetches a receipt of a specified tx hash
func (c *EthClient) GetTransactionReceipt(client *http.Client, txhash string) (*eth_domain.TransactionReceipt, error) {
	raw, err := c.do(client, "eth_getTransactionReceipt", []interface{}{txhash})
	if err != nil {
		return nil, err
	}

	resp := new(eth_domain.TransactionReceiptResponse)
	if err := json.Unmarshal(raw, resp); err != nil {
		return nil, errors.Wrap(err, "failed to fetch a tx receipt")
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	return resp.Receipt, nil
}

func toHex(n int64) string {
	return fmt.Sprintf("0x%s", fmt.Sprintf("%0x", n))
}

func dropHexPrefix(s string) string {
	if s[:2] == "0x" {
		return s[2:]
	}
	return s
}
