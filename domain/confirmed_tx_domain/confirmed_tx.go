package confirmed_tx_domain

import "github.com/D-Technologies/supervision/lib/mysqlutil"

// Repository represents an interface for infrastructure
type Repository interface {
	Create(sqle mysqlutil.SQLExecutor, ct *ConfirmedTx) error
	GetAll(sqle mysqlutil.SQLExecutor) ([]*ConfirmedTx, error)
	Delete(sqle mysqlutil.SQLExecutor, cts []*ConfirmedTx) error
}

// ConfirmedTx represents a confirmed tx
type ConfirmedTx struct {
	TxHash  string
	From    string
	TokenID int64
}
