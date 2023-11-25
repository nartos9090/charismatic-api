package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	SQLiteConf struct {
		DataSourceName string
	}

	Env struct {
		HttpPort  string
		JWTSecret string

		SQLite SQLiteConf
	}
)

var GlobalEnv Env

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Panic(`Can't find .env`)
	}

	var ok bool

	GlobalEnv.HttpPort, ok = os.LookupEnv(`HTTP_PORT`)
	if !ok {
		panic(`HTTP_PORT env not set`)
	}

	GlobalEnv.JWTSecret, ok = os.LookupEnv(`JWT_SECRET`)
	if !ok {
		panic(`JWT_SECRET env not set`)
	}

	GlobalEnv.SQLite.DataSourceName, ok = os.LookupEnv(`SQLITE_DSN`)
	if !ok {
		panic(`SQLITE_DSN env not set`)
	}

	fmt.Print(":: Config loaded\n")
}
