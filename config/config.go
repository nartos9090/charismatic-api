package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Env struct {
		HttpPort  string
		JWTSecret string
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
}
