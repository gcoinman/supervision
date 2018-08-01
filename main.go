package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/D-Technologies/go-tokentracker/domain/blocknumber"

	"github.com/D-Technologies/go-tokentracker/di"
	"github.com/D-Technologies/go-tokentracker/ethclient"
	"github.com/D-Technologies/go-tokentracker/lib/config"
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

	config.Setup()

	repository := di.InjectBlockNumRepository()
	sql := di.InjectSQL()

	bn := &blocknumberdomain.BlockNum{
		Num:  number,
		Hash: block.Hash,
	}
	if err := repository.Create(sql, bn); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	entity, err := repository.GetLatest(sql)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(entity)
}
