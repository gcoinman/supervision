package received_tx

import "github.com/D-Technologies/supervision/domain/received_tx_domain"

// Entity represents an entity
type Entity struct {
	Hash     string `db:"hash"`
	BlockNum int64  `db:"block_number"`
	From     string `db:"from"`
	TokenID  int64  `db:"token_id"`
	Status   string `db:"status"`
}

// NewEntity creates a new Entity
func NewEntity(r *received_tx_domain.ReceivedTx) *Entity {
	return &Entity{
		Hash:     r.Hash,
		BlockNum: r.BlockNum,
		From:     r.From,
		TokenID:  r.TokenID,
		Status:   r.Status,
	}
}

// NewEntities creates new entities
func NewEntities(rs []*received_tx_domain.ReceivedTx) []*Entity {
	es := make([]*Entity, 0, len(rs))
	for _, r := range rs {
		es = append(es, NewEntity(r))
	}
	return es
}

// Domain converts an entity to domain
func (e *Entity) Domain() *received_tx_domain.ReceivedTx {
	return &received_tx_domain.ReceivedTx{
		Hash:     e.Hash,
		BlockNum: e.BlockNum,
		From:     e.From,
		TokenID:  e.TokenID,
		Status:   e.Status,
	}
}
