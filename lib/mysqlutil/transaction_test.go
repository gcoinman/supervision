// +build local

package mysqlutil

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gopkg.in/gorp.v2"
)

var (
	dbmap *gorp.DbMap

	createTable = `
CREATE TABLE IF NOT EXISTS balance_test (
  user_id varchar(255) NOT NULL,
  amount decimal(65,20) DEFAULT NULL,
  PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
`
	dropTable = `
DROP TABLE IF EXISTS balance_test;
`
	insert = `
INSERT INTO balance_test VALUES(?, ?);
`
	updateAmount = `
UPDATE balance_test SET amount = ? WHERE user_id = ?;
`
)

func init() {
	db, err := sql.Open("mysql", "root:anchordev@tcp/crypto?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}

	dbmap = &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{Engine: "InnoDB",
			Encoding: "UTF8MB4",
		},
	}
}

func setup() {
	// create table
	if _, err := dbmap.Query(dropTable); err != nil {
		log.Fatalln(err)
	}
	if _, err := dbmap.Query(createTable); err != nil {
		log.Fatalln(err)
	}

	// init dasta
	if _, err := dbmap.Query(insert, "local_dummy", 0); err != nil {
		log.Fatalln(err)
	}
}

func teardown() {
	if _, err := dbmap.Query(dropTable); err != nil {
		log.Fatalln(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	// teardown()
	os.Exit(ret)
}

// TestRunInTransaction はトランザクションが正常に働いているかを確認するテスト。
func TestRunInTransaction(t *testing.T) {
	sql := &SQL{
		dbmap: dbmap,
	}

	wg := sync.WaitGroup{}
	// Queryを同時に実行しているQueryの数
	count := 0
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) error {

			// トランザクションの発行
			if err := RunInTransaction(sql, func(tx *Tx) error {
				defer wg.Done()

				// local_dummyユーザーに対してAmountのUpdate
				if _, err := tx.DB().Query(updateAmount, n, "local_dummy"); err != nil {
					return err
				}

				// 擬似的に長いQueryを表現するためのSleepと
				// 実行中であることを示すcountのインクリメント
				count++
				defer func() {
					count--
				}()
				time.Sleep(1 * time.Second)

				// count が1より大きい時はトランザクションが張れていないためエラー
				if count > 1 {
					err := errors.Errorf("failed tx: count = %d", count)
					t.Fatalf("%s", err)
					return err
				}

				return nil
			}); err != nil {
				t.Fatalf("%s", err)
			}

			return nil
		}(i)
	}
	wg.Wait()
}
