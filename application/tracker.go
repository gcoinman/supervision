package application

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/D-Technologies/go-tokentracker/domain/block"
	"github.com/D-Technologies/go-tokentracker/domain/blocknumber"
	"github.com/D-Technologies/go-tokentracker/domain/confirmedtransaction"
	"github.com/D-Technologies/go-tokentracker/domain/receivedtransaction"
	"github.com/D-Technologies/go-tokentracker/infrastructure/ethclient"
	"github.com/D-Technologies/go-tokentracker/lib/mysqlutil"
)

// TrackerApp is an application layer that tracks tokens
type TrackerApp struct {
	ContractAddr                   string
	ReceiveAddr                    string
	BlockNumRepository             blocknumberdomain.BlockNumRepository
	ReceivedTransactionRepository  receivedtransactiondomain.ReceivedTransactionRepository
	ConfirmedTransactionRepository confirmedtransactiondomain.ConfirmedTransactionRepository
	EthClient                      *ethclient.EthClient
	Client                         *http.Client
	SQL                            *mysqlutil.SQL
}

// NewApp creates a new TrackerApp
func NewApp(
	contractAddr string,
	receiveAddr string,
	br blocknumberdomain.BlockNumRepository,
	rr receivedtransactiondomain.ReceivedTransactionRepository,
	cr confirmedtransactiondomain.ConfirmedTransactionRepository,
	c *http.Client,
	ec *ethclient.EthClient,
	sql *mysqlutil.SQL,
) *TrackerApp {

	return &TrackerApp{
		ContractAddr:                   contractAddr,
		ReceiveAddr:                    receiveAddr,
		BlockNumRepository:             br,
		ReceivedTransactionRepository:  rr,
		ConfirmedTransactionRepository: cr,
		EthClient:                      ec,
		Client:                         c,
		SQL:                            sql,
	}
}

// Do executes all the logic for tracking tokens
func (t *TrackerApp) Do() error {
	blockNum, err := t.EthClient.GetBlockNumber(t.Client)
	if err != nil {
		return err
	}

	if err := t.scanBlocks(blockNum); err != nil {
		return err
	}

	if err := t.updateTxStatus(blockNum); err != nil {
		return err
	}

	if err := t.pushConfirmedTx(); err != nil {
		return err
	}

	return nil
}

func (t *TrackerApp) scanBlocks(blockNum int64) error {
	lastBlockNum, err := t.BlockNumRepository.GetLatest(t.SQL)
	if err != nil {
		if errors.Cause(err).Error() == "sql: no rows in result set" {
			lastBlockNum = &blocknumberdomain.BlockNum{
				Num: blockNum - 1,
			}
		} else {
			return err
		}
	}

	if lastBlockNum.Num == blockNum {
		return nil
	}

	for num := lastBlockNum.Num + 1; num <= blockNum; num++ {
		b, err := t.EthClient.GetBlockByBlockNumber(t.Client, num, true)
		if err != nil {
			return err
		}

		domainBlock := blockdomain.Block{
			Transactions: b.Transactions,
		}
		rts := domainBlock.Scan(t.ContractAddr, t.ReceiveAddr)
		if err := t.ReceivedTransactionRepository.CreateMulti(t.SQL, rts); err != nil {
			return err
		}

		if !t.BlockNumRepository.Exist(t.SQL, num) {
			if err := t.BlockNumRepository.Create(t.SQL, &blocknumberdomain.BlockNum{Num: num}); err != nil {
				return err
			}
		}
	}

	fmt.Printf("scaned blocks between %d and %d\n", blockNum, lastBlockNum.Num)

	return nil
}

func (t *TrackerApp) updateTxStatus(blockNum int64) error {
	rts, err := t.ReceivedTransactionRepository.GetSuccessAndPendingTransactions(t.SQL)
	if err != nil {
		return err
	}

	if len(rts) == 0 {
		return nil
	}

	for _, rt := range rts {
		switch rt.Status {
		case receivedtransactiondomain.Pending:
			receipt, err := t.EthClient.GetTransactionReceipt(t.Client, rt.Hash)
			if err != nil {
				return err
			}

			if receipt.IsSuccess() {
				rt.Success()
			} else {
				rt.Error()
			}

		case receivedtransactiondomain.Success:
			if rt.Confirmed(blockNum) {
				rt = rt.Complete()
			}

			ct := &confirmedtransactiondomain.ConfirmedTransaction{
				TxHash:  rt.Hash,
				From:    rt.From,
				TokenID: rt.TokenID,
			}

			if err := t.ConfirmedTransactionRepository.Create(t.SQL, ct); err != nil {
				return err
			}
		}

		if err := t.ReceivedTransactionRepository.Update(t.SQL, rt); err != nil {
			return err
		}

		fmt.Printf("\nUpdate received transaction found at %d, currently at %d. Status: %s\n", rt.BlockNum, blockNum, rt.Status)
	}

	return nil
}

func (t *TrackerApp) pushConfirmedTx() error {
	cts, err := t.ConfirmedTransactionRepository.GetAll(t.SQL)
	if err != nil {
		return err
	}

	for _, ct := range cts {
		// TODO: send a request
		fmt.Print(ct)
	}

	if err := t.ConfirmedTransactionRepository.Delete(t.SQL, cts); err != nil {
		return err
	}

	return nil
}
