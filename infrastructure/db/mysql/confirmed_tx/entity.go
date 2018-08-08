package confirmed_tx

import "github.com/D-Technologies/supervision/domain/confirmed_tx_domain"

// Entity represents an entity
type Entity struct {
	TxHash  string `db:"hash"`
	From    string `db:"from"`
	TokenID int64  `db:"token_id"`
}

// NewEntity creates a new entity
func NewEntity(ct *confirmed_tx_domain.ConfirmedTx) *Entity {
	return &Entity{
		TxHash:  ct.TxHash,
		From:    ct.From,
		TokenID: ct.TokenID,
	}
}

// Domain converts an entity to domain
func (e *Entity) Domain() *confirmed_tx_domain.ConfirmedTx {
	return &confirmed_tx_domain.ConfirmedTx{
		TxHash:  e.TxHash,
		From:    e.From,
		TokenID: e.TokenID,
	}
}
