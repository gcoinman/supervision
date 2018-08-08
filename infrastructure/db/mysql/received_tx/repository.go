package received_tx

import (
	"strings"

	"github.com/D-Technologies/supervision/domain/received_tx_domain"
	"github.com/D-Technologies/supervision/lib/mysqlutil"
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
func (r *Repository) Create(sqle mysqlutil.SQLExecutor, rt *received_tx_domain.ReceivedTx) error {
	const errtag = "Repository.Create failed "
	e := NewEntity(rt)

	if err := sqle.DB().Insert(e); err != nil {
		return errors.Wrapf(err, errtag)
	}

	return nil
}

// CreateMulti creates multiple new entities
func (r *Repository) CreateMulti(sqle mysqlutil.SQLExecutor, rts []*received_tx_domain.ReceivedTx) error {
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

// Update updates an entity
func (r *Repository) Update(sqle mysqlutil.SQLExecutor, rt *received_tx_domain.ReceivedTx) error {
	const errtag = "Repository.Update failed"

	e := NewEntity(rt)

	if _, err := sqle.DB().Update(e); err != nil {
		return errors.Wrapf(err, errtag)
	}

	return nil
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

// GetSuccessAndPendingTransactions fetches entities with status of pending or success
func (r *Repository) GetSuccessAndPendingTransactions(sqle mysqlutil.SQLExecutor) ([]*received_tx_domain.ReceivedTx, error) {
	const errtag = "Repository.GetSuccessAndPendingTransactions failed"

	var es []*Entity

	q := strings.Join([]string{"SELECT * FROM", TableName, "WHERE `status`=? OR `status`=? FOR UPDATE"}, " ")
	if _, err := sqle.DB().Select(&es, q, "pending", "success"); err != nil {
		return nil, errors.Wrapf(err, errtag)
	}

	rts := make([]*received_tx_domain.ReceivedTx, 0, len(es))
	for _, v := range es {
		rts = append(rts, v.Domain())
	}
	return rts, nil
}
