package blocknumber

import "github.com/D-Technologies/supervision/domain/block_number_domain"

// Entity represents an entity for blocknumber
type Entity struct {
	BlockNum int64 `db:"block_number"`
}

// NewEntity creates a new entity
func NewEntity(b *block_number_domain.BlockNum) *Entity {
	return &Entity{
		BlockNum: b.Num,
	}
}

// Domain converts an entity to a BlockNum in domain layer
func (e *Entity) Domain() *block_number_domain.BlockNum {
	return &block_number_domain.BlockNum{
		Num: e.BlockNum,
	}
}
