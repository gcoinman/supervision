package di

import (
	"github.com/D-Technologies/go-tokentracker/domain/blocknumber"
	"github.com/D-Technologies/go-tokentracker/domain/confirmedtransaction"
	"github.com/D-Technologies/go-tokentracker/domain/receivedtransaction"
	"github.com/D-Technologies/go-tokentracker/infrastructure/db/mysql/blocknumber"
	"github.com/D-Technologies/go-tokentracker/infrastructure/db/mysql/confirmed_transaction"
	"github.com/D-Technologies/go-tokentracker/infrastructure/db/mysql/received_transaction"
)

var blockNumRepository blocknumberdomain.BlockNumRepository

// InjectBlockNumRepository injects a blocknum repository
func InjectBlockNumRepository() blocknumberdomain.BlockNumRepository {
	if blockNumRepository != nil {
		return blockNumRepository
	}

	blockNumRepository = blocknumber.NewRepository()

	return blockNumRepository
}

var receivedTransactionRepository receivedtransactiondomain.ReceivedTransactionRepository

// InjectReceivedTransactionRepository injects a received transaction repository
func InjectReceivedTransactionRepository() receivedtransactiondomain.ReceivedTransactionRepository {
	if receivedTransactionRepository != nil {
		return receivedTransactionRepository
	}

	receivedTransactionRepository = receivedtransaction.NewRepository()

	return receivedTransactionRepository
}

var confirmedTransactionRepository confirmedtransactiondomain.ConfirmedTransactionRepository

// InjectConfirmedTransactionRepository injects a confirmed transaction repository
func InjectConfirmedTransactionRepository() confirmedtransactiondomain.ConfirmedTransactionRepository {
	if confirmedTransactionRepository != nil {
		return confirmedTransactionRepository
	}

	confirmedTransactionRepository = confirmedtransaction.NewRepository()

	return confirmedTransactionRepository
}
