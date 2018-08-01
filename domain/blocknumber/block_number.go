package blocknumberdomain

import "github.com/D-Technologies/go-tokentracker/lib/mysqlutil"

// BlockNumRepository is an interface for a repositroy of blocknumber
type BlockNumRepository interface {
	GetLatest(sqle mysqlutil.SQLExecutor) (*BlockNum, error)
	Create(sqle mysqlutil.SQLExecutor, b *BlockNum) error
}

// BlockNum todo
type BlockNum struct {
	Num  int64
	Hash string
}
