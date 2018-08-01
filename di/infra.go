package di

import (
	"database/sql"
	"fmt"
	stdlog "log"
	"os"

	"github.com/D-Technologies/go-tokentracker/infrastructure/db/mysql/blocknumber"
	"github.com/D-Technologies/go-tokentracker/lib/config"
	"github.com/D-Technologies/go-tokentracker/lib/mysqlutil"

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
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/go-tokentracker",
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

	if err := dbmap.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}

	mysql = mysqlutil.NewSQL(dbmap)

	return mysql
}
