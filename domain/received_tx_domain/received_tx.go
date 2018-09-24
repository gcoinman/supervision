package received_tx_domain

import (
	"github.com/D-Technologies/supervision/lib/mysqlutil"
)

// Repository is an interface for ReceivedTransaction infrastrusture
type Repository interface {
	Create(sqle mysqlutil.SQLExecutor, rt *ReceivedTx) error
	CreateMulti(sqle mysqlutil.SQLExecutor, rts []*ReceivedTx) error
	Update(sqle mysqlutil.SQLExecutor, rt *ReceivedTx) error
	Exist(sqle mysqlutil.SQLExecutor, hash string) bool
	GetSuccessAndPendingTransactions(sqle mysqlutil.SQLExecutor) ([]*ReceivedTx, error)
}

var (
	// Pending represnets pending status
	Pending = "pending"

	// Success represents success status
	Success = "success"

	// Error represents error status
	Error = "error"

	// Completed represents complete status
	Completed = "completed"
)

// ReceivedTx represents found transactions in Scan method
type ReceivedTx struct {
	Hash         string
	BlockNum     int64
	From         string
	ReceiveAddir string
	TokenID      int64
	Status       string
}

// Confirmed checks if a received-tx is confirmed by n th blcok
func (rt *ReceivedTx) Confirmed(blockNum int64) bool {
	return blockNum-rt.BlockNum > 5
}

// Complete changes a status to complete
func (rt *ReceivedTx) Complete() *ReceivedTx {
	rt.Status = Completed
	return rt
}

// Success changes a status to success
func (rt *ReceivedTx) Success() *ReceivedTx {
	rt.Status = Success
	return rt
}

// Error changes a status to error
func (rt *ReceivedTx) Error() *ReceivedTx {
	rt.Status = Error
	return rt
}
