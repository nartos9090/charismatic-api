package sqlite

import (
	"fmt"
	"go-api-echo/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sqlx.DB
var dbDriver = "sqlite3"

func InitSqlite(cfg config.SQLiteConf) {
	dsn := cfg.DataSourceName
	Db = sqlx.MustConnect(dbDriver, dsn)

	fmt.Print(":: SQLite connected\n")
}
