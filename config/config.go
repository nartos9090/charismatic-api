package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	GeminiConf struct {
		ApiKey string
	}

	ElevenLabsConf struct {
		ApiKey string
	}

	DalleConf struct {
		ApiKey string
	}

	Env struct {
		HttpPort  string
		JWTSecret string

		GeminiConf     GeminiConf
		ElevenLabsConf ElevenLabsConf
		DalleConf      DalleConf
	}
)

var GlobalEnv Env

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Panic("Can't find .env")
	}

	var ok bool

	GlobalEnv.HttpPort, ok = os.LookupEnv("HTTP_PORT")
	if !ok {
		panic("HTTP_PORT env not set")
	}

	GlobalEnv.JWTSecret, ok = os.LookupEnv("JWT_SECRET")
	if !ok {
		panic("JWT_SECRET env not set")
	}

	GlobalEnv.ElevenLabsConf.ApiKey, ok = os.LookupEnv("ELEVENLABS_API_KEY")
	if !ok {
		panic("ELEVENLABS_API_KEY env not set")
	}

	GlobalEnv.DalleConf.ApiKey, ok = os.LookupEnv("DALLE_API_KEY")
	if !ok {
		panic("DALLE_API_KEY env not set")
	}

	GlobalEnv.GeminiConf.ApiKey, ok = os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		panic("GEMINI_API_KEY env not set")
	}
}
