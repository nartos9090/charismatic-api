package main

import (
	"go-api-echo/config"
	db_mysql "go-api-echo/internal/pkg/db/mysql"
	"go-api-echo/internal/server"
)

func main() {
	config.InitConfig()

	db_mysql.InitMysql()

	server.InitServer(config.GlobalEnv.HttpPort)
}
