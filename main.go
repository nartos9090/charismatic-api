package main

import (
	"go-api-echo/config"
	"go-api-echo/internal/server"
)

func main() {
	config.InitConfig()

	server.InitServer(config.GlobalEnv.HttpPort)
}
