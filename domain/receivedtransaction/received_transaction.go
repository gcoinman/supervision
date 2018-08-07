package receivedtransactiondomain

import (
	"github.com/D-Technologies/go-tokentracker/lib/mysqlutil"
)

// ReceivedTransactionRepository is an interface for ReceivedTransaction infrastrusture
type ReceivedTransactionRepository interface {
	Create(sqle mysqlutil.SQLExecutor, rt *ReceivedTransaction) error
	CreateMulti(sqle mysqlutil.SQLExecutor, rts []*ReceivedTransaction) error
	Update(sqle mysqlutil.SQLExecutor, rt *ReceivedTransaction) error
	Exist(sqle mysqlutil.SQLExecutor, hash string) bool
	GetSuccessAndPendingTransactions(sqle mysqlutil.SQLExecutor) ([]*ReceivedTransaction, error)
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

// ReceivedTransaction represents found transactions in Scan method
type ReceivedTransaction struct {
	Hash     string
	BlockNum int64
	From     string
	TokenID  int64
	Status   string
}

// Confirmed checks if a received-transaction is confirmed by n th blcok
func (rt *ReceivedTransaction) Confirmed(blockNum int64) bool {
	return blockNum-rt.BlockNum > 5
}

// Complete changes a status to complete
func (rt *ReceivedTransaction) Complete() *ReceivedTransaction {
	rt.Status = Completed
	return rt
}

// Success changes a status to success
func (rt *ReceivedTransaction) Success() *ReceivedTransaction {
	rt.Status = Success
	return rt
}

// Error changes a status to error
func (rt *ReceivedTransaction) Error() *ReceivedTransaction {
	rt.Status = Error
	return rt
}
