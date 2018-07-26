package ethclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// EthClient handles an interacting with ethereum node
type EthClient struct {
	apiKey  string
	chainID ChainID

	EthAPIClient
}

// New creates an ethclient
func New(apiKey string, chainID ChainID) *EthClient {
	return &EthClient{
		apiKey:  apiKey,
		chainID: chainID,
	}
}

type rpcRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

func (c *EthClient) do(client *http.Client, method string, params interface{}) (json.RawMessage, error) {
	body, err := json.Marshal(rpcRequest{
		JsonRpc: "2.0",
		ID:      "1",
		Method:  method,
		Params:  params,
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.chainID.URL(c.apiKey), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		const format = "failed to HTTP request with status code %d: %s"
		msg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			msg = []byte("no error message returned from ethereum node")
		}
		return nil, errors.Errorf(format, resp.StatusCode, msg)
	}

	var rawMsg json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&rawMsg); err != nil {
		return nil, err
	}
	return rawMsg, nil
}
