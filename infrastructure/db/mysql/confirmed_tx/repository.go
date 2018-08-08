package confirmed_tx

import (
	"strings"

	"github.com/D-Technologies/supervision/domain/confirmed_tx_domain"
	"github.com/D-Technologies/supervision/lib/mysqlutil"
	"github.com/pkg/errors"
)

const (
	// TableName for mysql
	TableName = "confirmed_transactions"
)

// Repository represents an reporitory for confirmedtransaction
type Repository struct {
}

// NewRepository creates a new repository
func NewRepository() *Repository {
	return &Repository{}
}

// Create creates a new entity
func (r *Repository) Create(sqle mysqlutil.SQLExecutor, ct *confirmed_tx_domain.ConfirmedTx) error {
	const errtag = "Repository.Create failed "
	e := NewEntity(ct)

	if err := sqle.DB().Insert(e); err != nil {
		return errors.Wrapf(err, errtag)
	}

	return nil
}

// GetAll gets all the entities
func (r *Repository) GetAll(sqle mysqlutil.SQLExecutor) ([]*confirmed_tx_domain.ConfirmedTx, error) {
	const errtag = "Repository.GetAll failed"

	var es []*Entity

	q := strings.Join([]string{"SELECT * FROM", TableName, "FOR UPDATE"}, " ")
	if _, err := sqle.DB().Select(&es, q); err != nil {
		return nil, err
	}

	cts := make([]*confirmed_tx_domain.ConfirmedTx, 0, len(es))
	for _, v := range es {
		cts = append(cts, v.Domain())
	}
	return cts, nil
}

// Delete deletes an entity
func (r *Repository) Delete(sqle mysqlutil.SQLExecutor, cts []*confirmed_tx_domain.ConfirmedTx) error {
	for _, v := range cts {
		if err := r.delete(sqle, v); err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) delete(sqle mysqlutil.SQLExecutor, ct *confirmed_tx_domain.ConfirmedTx) error {
	const errtag = "Repository.Delete failed"
	e := NewEntity(ct)

	if _, err := sqle.DB().Delete(e); err != nil {
		return err
	}

	return nil
}
