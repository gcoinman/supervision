package application

import (
	"fmt"
	"net/http"

	"github.com/D-Technologies/go-tokentracker/domain/blocknumber"
	"github.com/D-Technologies/go-tokentracker/infrastructure/ethclient"
	"github.com/D-Technologies/go-tokentracker/lib/mysqlutil"
)

// TrackerApp todo
type TrackerApp struct {
	WatchAddr          string
	BlockNumRepository blocknumberdomain.BlockNumRepository
	EthClient          *ethclient.EthClient
	Client             *http.Client
	SQL                *mysqlutil.SQL
}

// NewApp todo
func NewApp(addr string, b blocknumberdomain.BlockNumRepository, c *http.Client, ec *ethclient.EthClient, sql *mysqlutil.SQL) *TrackerApp {
	return &TrackerApp{
		WatchAddr:          addr,
		BlockNumRepository: b,
		EthClient:          ec,
		Client:             c,
		SQL:                sql,
	}
}

// Do todo
func (t *TrackerApp) Do() error {
	blockNum, err := t.EthClient.GetBlockNumber(t.Client)
	if err != nil {
		return err
	}
	fmt.Printf("current block number is %d\n", blockNum)

	lastBlockNum, err := t.BlockNumRepository.GetLatest(t.SQL)
	if err != nil {
		return err
	}
	fmt.Printf("last block number is %d\n", lastBlockNum.Num)

	for num := lastBlockNum.Num; num <= blockNum; num++ {
		block, err := t.EthClient.GetBlockByBlockNumber(t.Client, num, true)
		if err != nil {
			return err
		}

		fmt.Printf("Block hash is %s\n", block.Hash)
		// TODO: check blocks here

		if !t.BlockNumRepository.Exist(t.SQL, num) {
			if err := t.BlockNumRepository.Create(t.SQL, &blocknumberdomain.BlockNum{
				Hash: block.Hash,
				Num:  num,
			}); err != nil {
				return err
			}
		}
	}

	return nil
}
