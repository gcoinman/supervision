package block_number_domain

import "github.com/D-Technologies/supervision/lib/mysqlutil"

// Repository is an interface for a repositroy of blocknumber
type Repository interface {
	GetLatest(sqle mysqlutil.SQLExecutor) (*BlockNum, error)
	Create(sqle mysqlutil.SQLExecutor, b *BlockNum) error
	Exist(sqle mysqlutil.SQLExecutor, num int64) bool
}

// BlockNum represents a block number in domain layer
type BlockNum struct {
	Num int64
}
