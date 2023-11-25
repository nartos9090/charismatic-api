package main

import (
	"go-api-echo/config"
	"go-api-echo/internal/pkg/db/sqlite"
	"go-api-echo/internal/server"
)

func main() {
	config.InitConfig()

	sqlite.InitSqlite(config.GlobalEnv.SQLite)

	server.InitServer(config.GlobalEnv.HttpPort)
}
