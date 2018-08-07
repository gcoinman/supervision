package blocknumberdomain

import "github.com/D-Technologies/go-tokentracker/lib/mysqlutil"

// BlockNumRepository is an interface for a repositroy of blocknumber
type BlockNumRepository interface {
	GetLatest(sqle mysqlutil.SQLExecutor) (*BlockNum, error)
	Create(sqle mysqlutil.SQLExecutor, b *BlockNum) error
	Exist(sqle mysqlutil.SQLExecutor, num int64) bool
}

// BlockNum represents a block number in domain layer
type BlockNum struct {
	Num int64
}
