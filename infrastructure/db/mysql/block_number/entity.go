package blocknumber

import (
	"time"

	"github.com/D-Technologies/go-tokentracker/domain/blocknumber"
)

// Entity represents an entity for blocknumber
type Entity struct {
	BlockNum  int64     `db:"block_number"`
	Hash      string    `db:"hash"`
	CreatedAt time.Time `db:"created_at"`
}

// NewEntity creates a new entity
func NewEntity(b *blocknumberdomain.BlockNum) *Entity {
	return &Entity{
		BlockNum:  b.Num,
		Hash:      b.Hash,
		CreatedAt: time.Now(),
	}
}

// Domain converts an entity to a BlockNum in domain layer
func (e *Entity) Domain() *blocknumberdomain.BlockNum {
	return &blocknumberdomain.BlockNum{
		Num:  e.BlockNum,
		Hash: e.Hash,
	}
}
