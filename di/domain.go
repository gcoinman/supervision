package di

import (
	"github.com/D-Technologies/supervision/domain/block_number_domain"
	"github.com/D-Technologies/supervision/domain/confirmed_tx_domain"
	"github.com/D-Technologies/supervision/domain/received_tx_domain"
	"github.com/D-Technologies/supervision/infrastructure/db/mysql/blocknumber"
	"github.com/D-Technologies/supervision/infrastructure/db/mysql/confirmed_tx"
	"github.com/D-Technologies/supervision/infrastructure/db/mysql/received_tx"
)

var blockNumRepository block_number_domain.Repository

// InjectBlockNumRepository injects a blocknum repository
func InjectBlockNumRepository() block_number_domain.Repository {
	if blockNumRepository != nil {
		return blockNumRepository
	}

	blockNumRepository = blocknumber.NewRepository()

	return blockNumRepository
}

var receivedTransactionRepository received_tx_domain.Repository

// InjectReceivedTransactionRepository injects a received tx repository
func InjectReceivedTransactionRepository() received_tx_domain.Repository {
	if receivedTransactionRepository != nil {
		return receivedTransactionRepository
	}

	receivedTransactionRepository = received_tx.NewRepository()

	return receivedTransactionRepository
}

var confirmedTransactionRepository confirmed_tx_domain.Repository

// InjectConfirmedTransactionRepository injects a confirmed tx repository
func InjectConfirmedTransactionRepository() confirmed_tx_domain.Repository {
	if confirmedTransactionRepository != nil {
		return confirmedTransactionRepository
	}

	confirmedTransactionRepository = confirmed_tx.NewRepository()

	return confirmedTransactionRepository
}
