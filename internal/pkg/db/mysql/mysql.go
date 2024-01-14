package db_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-api-echo/config"
	"log"
)

var Db *sqlx.DB
var dbDriver = `mysql`

func InitMysql() {
	user := config.GlobalEnv.DB.User
	pass := config.GlobalEnv.DB.Pass
	host := config.GlobalEnv.DB.Host
	port := config.GlobalEnv.DB.Port
	schema := config.GlobalEnv.DB.Schema

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, schema)

	db, err := sql.Open(dbDriver, url)
	Db = sqlx.NewDb(db, dbDriver)

	if err != nil {
		log.Print(err.Error())
		log.Panic(`Can't connect to db`)
	}
}
