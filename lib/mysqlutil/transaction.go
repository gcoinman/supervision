package mysqlutil

import (
	"gopkg.in/gorp.v2"
)

// SQLExecutor represents an executor interface
type SQLExecutor interface {
	DB() gorp.SqlExecutor
}

// SQL represents an structure for sql
type SQL struct {
	dbmap *gorp.DbMap
}

// NewSQL creates a new sql
func NewSQL(dbmap *gorp.DbMap) *SQL {
	return &SQL{
		dbmap: dbmap,
	}
}

// DB returns a sql executor
func (s *SQL) DB() gorp.SqlExecutor {
	return s.dbmap
}

// Begin returns a new tx
func (s *SQL) Begin() (*Tx, error) {
	tx, err := s.dbmap.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{
		tx: tx,
	}, nil
}

// Tx represents a structure for tx
type Tx struct {
	tx *gorp.Transaction
}

// DB returns a new sql executor
func (t *Tx) DB() gorp.SqlExecutor {
	return t.tx
}

// Commit commits a tx
func (t *Tx) Commit() error {
	return t.tx.Commit()
}

// Rollback executes rollback
func (t *Tx) Rollback() error {
	return t.tx.Rollback()
}

// RunInTransaction runs in tx
func RunInTransaction(db *SQL, txFunc func(*Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	// 関数を抜けた際に実行される。 errには抜ける瞬間のerrがキャプチャされている
	defer func() {
		if p := recover(); p != nil {
			// panicが発生したときのためRollbackした後に再度panicを投げ直す
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// エラーは必ずRollbackする
			tx.Rollback()
		} else {
			// エラーでなければCommitする
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)

	return err
}
