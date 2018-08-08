package confirmedtransaction

import "github.com/D-Technologies/go-tokentracker/domain/confirmedtransaction"

// Entity represents an entity
type Entity struct {
	TxHash  string `db:"hash"`
	From    string `db:"from"`
	TokenID int64  `db:"token_id"`
}

// NewEntity creates a new entity
func NewEntity(ct *confirmedtransactiondomain.ConfirmedTransaction) *Entity {
	return &Entity{
		TxHash:  ct.TxHash,
		From:    ct.From,
		TokenID: ct.TokenID,
	}
}

// Domain converts an entity to domain
func (e *Entity) Domain() *confirmedtransactiondomain.ConfirmedTransaction {
	return &confirmedtransactiondomain.ConfirmedTransaction{
		TxHash:  e.TxHash,
		From:    e.From,
		TokenID: e.TokenID,
	}
}
