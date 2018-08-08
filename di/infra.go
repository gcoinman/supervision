package di

import (
	"database/sql"
	"fmt"
	stdlog "log"
	"os"

	"github.com/D-Technologies/supervision/infrastructure/db/mysql/blocknumber"
	"github.com/D-Technologies/supervision/infrastructure/db/mysql/confirmed_tx"
	"github.com/D-Technologies/supervision/infrastructure/db/mysql/received_tx"
	"github.com/D-Technologies/supervision/infrastructure/ethclient"
	"github.com/D-Technologies/supervision/lib/config"
	"github.com/D-Technologies/supervision/lib/mysqlutil"

	// empty import
	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v2"
)

var db *sql.DB

// InjectDB injects a DB
func InjectDB() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/supervision",
		config.DefaultConfig.DbUser,
		config.DefaultConfig.DbPassword,
		config.DefaultConfig.DbHost,
	))

	if err != nil {
		panic(err)
	}
	return db
}

var mysql *mysqlutil.SQL

// InjectSQL injects a sql
func InjectSQL() *mysqlutil.SQL {
	if mysql != nil {
		return mysql
	}
	dbmap := &gorp.DbMap{
		Db: InjectDB(),
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8MB4",
		},
	}

	dbmap.TraceOn("gorp", stdlog.New(os.Stderr, "gorptest: ", stdlog.Lmicroseconds))

	dbmap.AddTableWithName(blocknumber.Entity{}, blocknumber.TableName).SetKeys(false, "BlockNum")
	dbmap.AddTableWithName(received_tx.Entity{}, received_tx.TableName).SetKeys(false, "Hash")
	dbmap.AddTableWithName(confirmed_tx.Entity{}, confirmed_tx.TableName).SetKeys(false, "TxHash")

	if err := dbmap.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}

	mysql = mysqlutil.NewSQL(dbmap)

	return mysql
}

var client *ethclient.EthClient

// InjectEthClient injects ethclient
func InjectEthClient() *ethclient.EthClient {
	if client != nil {
		return client
	}

	// Ropsten
	//client = ethclient.New("z1sEfnzz0LLMsdYMX4PV", ethclient.Ropsten)

	// Localhost
	client = ethclient.New("", ethclient.Localhost)

	return client
}
