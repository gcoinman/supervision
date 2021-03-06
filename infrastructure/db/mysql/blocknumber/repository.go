package blocknumber

import (
	"strings"

	"github.com/D-Technologies/supervision/domain/block_number_domain"
	"github.com/D-Technologies/supervision/lib/mysqlutil"
	"github.com/pkg/errors"
)

const (
	// TableName for mysql
	TableName = "block_numbers"
)

// Repository represents a repository for block_number
type Repository struct {
}

// NewRepository creates a repository
func NewRepository() *Repository {
	return &Repository{}
}

// GetLatest fetches the latest element from a DB
func (r *Repository) GetLatest(sqle mysqlutil.SQLExecutor) (*block_number_domain.BlockNum, error) {
	const errtag = "Repository.Get failed"
	var e Entity

	q := strings.Join([]string{"SELECT * FROM", TableName, "WHERE block_number=(SELECT Max(block_number) FROM block_numbers)"}, " ")
	if err := sqle.DB().SelectOne(&e, q); err != nil {
		return nil, errors.Wrap(err, errtag)
	}

	return e.Domain(), nil
}

// Create creates a new entity of blocknumber
func (r *Repository) Create(sqle mysqlutil.SQLExecutor, b *block_number_domain.BlockNum) error {
	const errtag = "Repository.Create failed "
	e := NewEntity(b)

	if err := sqle.DB().Insert(e); err != nil {
		return errors.Wrapf(err, errtag)
	}

	return nil
}

// Exist checks if an entity with same primary key exists
func (r *Repository) Exist(sqle mysqlutil.SQLExecutor, num int64) bool {
	const errtag = "Repository.Exist failed"
	var e Entity

	q := strings.Join([]string{"SELECT * FROM", TableName, "WHERE `block_number`=? FOR UPDATE"}, " ")
	if err := sqle.DB().SelectOne(&e, q, num); err != nil {
		return false
	}
	return true
}
