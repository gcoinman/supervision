package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/D-Technologies/go-tokentracker/ethclient"
)

func main() {
	client := http.DefaultClient
	ethclient := ethclient.New("z1sEfnzz0LLMsdYMX4PV", ethclient.Ropsten)
	number, err := ethclient.GetBlockNumber(client)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	block, err := ethclient.GetBlockByBlockNumber(client, number, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	fmt.Printf("Block Hash is %s\n", block.Hash)
}
