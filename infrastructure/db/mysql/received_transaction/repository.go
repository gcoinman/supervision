package receivedtransaction

import (
	"strings"

	"github.com/D-Technologies/go-tokentracker/domain/receivedtransaction"
	"github.com/D-Technologies/go-tokentracker/lib/mysqlutil"
	"github.com/pkg/errors"
)

const (
	// TableName for mysql
	TableName = "received_transactions"
)

// Repository represents a repository for block_number
type Repository struct {
}

// NewRepository creates a repository
func NewRepository() *Repository {
	return &Repository{}
}

// Create creates a new entity
func (r *Repository) Create(sqle mysqlutil.SQLExecutor, rt *receivedtransactiondomain.ReceivedTransaction) error {
	const errtag = "Repository.Create failed "
	e := NewEntity(rt)

	if err := sqle.DB().Insert(e); err != nil {
		return errors.Wrapf(err, errtag)
	}

	return nil
}

// Update updates an entity
func (r *Repository) Update(sqle mysqlutil.SQLExecutor, rt *receivedtransactiondomain.ReceivedTransaction) error {
	const errtag = "Repository.Update failed"

	e := NewEntity(rt)

	if _, err := sqle.DB().Update(e); err != nil {
		return errors.Wrapf(err, errtag)
	}

	return nil
}

// GetUncompletedTransaction fetches entities with status of not completed (pending, success, error)
func (r *Repository) GetUncompletedTransaction(sqle mysqlutil.SQLExecutor) ([]*receivedtransactiondomain.ReceivedTransaction, error) {
	const errtag = "Repository.GetPendingTransactions failed"

	var es []*Entity

	q := strings.Join([]string{"SELECT * FROM", TableName, "WHERE `status`!=? FOR UPDATE"}, " ")
	if _, err := sqle.DB().Select(&es, q, "completed"); err != nil {
		return nil, errors.Wrapf(err, errtag)
	}

	rts := make([]*receivedtransactiondomain.ReceivedTransaction, 0, len(es))
	for _, v := range es {
		rts = append(rts, v.Domain())
	}
	return rts, nil
}

// GetTransactions fetches entities with status of pending
func (r *Repository) GetTransactions(sqle mysqlutil.SQLExecutor, status string) ([]*receivedtransactiondomain.ReceivedTransaction, error) {
	const errtag = "Repository.GetPendingTransactions failed"

	var es []*Entity

	q := strings.Join([]string{"SELECT * FROM", TableName, "WHERE `status`=? FOR UPDATE"}, " ")
	if _, err := sqle.DB().Select(&es, q, status); err != nil {
		return nil, errors.Wrapf(err, errtag)
	}

	rts := make([]*receivedtransactiondomain.ReceivedTransaction, 0, len(es))
	for _, v := range es {
		rts = append(rts, v.Domain())
	}
	return rts, nil
}

// Exist checks if an entity exists with the same primary key
func (r *Repository) Exist(sqle mysqlutil.SQLExecutor, hash string) bool {
	const errtag = "Repository.Exist failed"
	var e Entity

	q := strings.Join([]string{"SELECT * FROM", TableName, "WHERE `hash`=? FOR UPDATE"}, " ")
	if err := sqle.DB().SelectOne(&e, q, hash); err != nil {
		return false
	}
	return true
}

// CreateMulti creates multiple new entities
func (r *Repository) CreateMulti(sqle mysqlutil.SQLExecutor, rts []*receivedtransactiondomain.ReceivedTransaction) error {
	for _, rt := range rts {
		if r.Exist(sqle, rt.Hash) {
			return nil
		}
		if err := r.Create(sqle, rt); err != nil {
			return err
		}
	}
	return nil
}
