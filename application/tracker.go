package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/D-Technologies/supervision/proto"

	"github.com/pkg/errors"

	"github.com/D-Technologies/supervision/domain/block_domain"
	"github.com/D-Technologies/supervision/domain/block_number_domain"
	"github.com/D-Technologies/supervision/domain/confirmed_tx_domain"
	"github.com/D-Technologies/supervision/domain/received_tx_domain"
	"github.com/D-Technologies/supervision/infrastructure/ethclient"
	"github.com/D-Technologies/supervision/lib/mysqlutil"
)

// TrackerApp is an application layer that tracks tokens
type TrackerApp struct {
	StartBlockNum                  int64
	ContractAddr                   string
	ReceiveAddrs                   []string
	BlockNumRepository             block_number_domain.Repository
	ReceivedTransactionRepository  received_tx_domain.Repository
	ConfirmedTransactionRepository confirmed_tx_domain.Repository
	EthClient                      *ethclient.EthClient
	Client                         *http.Client
	DepositClient                  deposit.DepositServiceClient
	SQL                            *mysqlutil.SQL
}

// NewApp creates a new TrackerApp
func NewApp(
	startBlockNum int64,
	contractAddr string,
	receiveAddrs []string,
	br block_number_domain.Repository,
	rr received_tx_domain.Repository,
	cr confirmed_tx_domain.Repository,
	c *http.Client,
	ec *ethclient.EthClient,
	dc deposit.DepositServiceClient,
	sql *mysqlutil.SQL,
) *TrackerApp {

	return &TrackerApp{
		StartBlockNum:                  startBlockNum,
		ContractAddr:                   contractAddr,
		ReceiveAddrs:                   receiveAddrs,
		BlockNumRepository:             br,
		ReceivedTransactionRepository:  rr,
		ConfirmedTransactionRepository: cr,
		EthClient:                      ec,
		Client:                         c,
		DepositClient:                  dc,
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
			lastBlockNum = &block_number_domain.BlockNum{
				Num: t.StartBlockNum,
			}
		} else {
			return err
		}
	}

	if lastBlockNum.Num == blockNum {
		return nil
	}

	fmt.Printf("\n\nScanning blocks between %d and %d\n\n", lastBlockNum.Num, blockNum)

	for num := lastBlockNum.Num + 1; num <= blockNum; num++ {
		b, err := t.EthClient.GetBlockByBlockNumber(t.Client, num, true)
		if err != nil {
			return err
		}

		domainBlock := block_domain.Block{
			Transactions: b.Transactions,
		}
		rts := domainBlock.Scan(t.ContractAddr, t.ReceiveAddrs)
		if err := t.ReceivedTransactionRepository.CreateMulti(t.SQL, rts); err != nil {
			return err
		}

		if t.BlockNumRepository.Exist(t.SQL, num) {
			break
		}

		if err := t.BlockNumRepository.Create(t.SQL, &block_number_domain.BlockNum{Num: num}); err != nil {
			return err
		}
	}

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
		case received_tx_domain.Pending:
			receipt, err := t.EthClient.GetTransactionReceipt(t.Client, rt.Hash)
			if err != nil {
				return err
			}

			if receipt.IsSuccess() {
				rt.Success()
			} else {
				rt.Error()
			}

		case received_tx_domain.Success:
			if rt.Confirmed(blockNum) {
				rt = rt.Complete()
			}

			ct := &confirmed_tx_domain.ConfirmedTx{
				TxHash:       rt.Hash,
				From:         rt.From,
				ReceivedAddr: rt.ReceiveAddir,
				TokenID:      rt.TokenID,
			}

			if err := t.ConfirmedTransactionRepository.Create(t.SQL, ct); err != nil {
				return err
			}
		}

		if err := t.ReceivedTransactionRepository.Update(t.SQL, rt); err != nil {
			return err
		}

		fmt.Printf("\n\nUpdate tx status with tokenID %d. Status: %s\n\n", rt.TokenID, rt.Status)
	}

	return nil
}

func (t *TrackerApp) pushConfirmedTx() error {
	cts, err := t.ConfirmedTransactionRepository.GetAll(t.SQL)
	if err != nil {
		return err
	}

	for _, ct := range cts {
		pd := &deposit.Deposit{
			TokenId:      ct.TokenID,
			From:         ct.From,
			ReceivedAddr: ct.ReceivedAddr,
		}
		_, err := t.DepositClient.PushDeposit(context.Background(), pd)
		if err != nil {
			return err
		}
		fmt.Printf("\n\nTx confirmed: %s\n\n", ct.TxHash)
	}

	if err := t.ConfirmedTransactionRepository.Delete(t.SQL, cts); err != nil {
		return err
	}

	return nil
}
