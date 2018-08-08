package confirmedtransactiondomain

import "github.com/D-Technologies/go-tokentracker/lib/mysqlutil"

// ConfirmedTransactionRepository represents an interface for infrastructure
type ConfirmedTransactionRepository interface {
	Create(sqle *mysqlutil.SQL, ct *ConfirmedTransaction) error
	GetAll(sqle *mysqlutil.SQL) ([]*ConfirmedTransaction, error)
	Delete(sqle *mysqlutil.SQL, cts []*ConfirmedTransaction) error
}

// ConfirmedTransaction represents a confirmed transaction
type ConfirmedTransaction struct {
	TxHash  string
	From    string
	TokenID int64
}
