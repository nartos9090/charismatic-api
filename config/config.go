package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	DBConf struct {
		User   string
		Pass   string
		Host   string
		Port   string
		Schema string
	}

	GeminiConf struct {
		ApiKey string
	}

	ElevenLabsConf struct {
		ApiKey string
	}

	DalleConf struct {
		ApiKey string
	}

	Google struct {
		ClientID     string
		ClientSecret string
	}

	Env struct {
		HttpPort  string
		JWTSecret string

		DB DBConf

		GeminiConf     GeminiConf
		ElevenLabsConf ElevenLabsConf
		DalleConf      DalleConf
		Google         Google
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

	GlobalEnv.DB.User, ok = os.LookupEnv("DB_USER")
	if !ok {
		panic("DB_USER env not set")
	}

	GlobalEnv.DB.Pass, ok = os.LookupEnv("DB_PASS")
	if !ok {
		panic("DB_PASS env not set")
	}

	GlobalEnv.DB.Host, ok = os.LookupEnv("DB_HOST")
	if !ok {
		panic("DB_HOST env not set")
	}

	GlobalEnv.DB.Port, ok = os.LookupEnv("DB_PORT")
	if !ok {
		panic("DB_PORT env not set")
	}

	GlobalEnv.DB.Schema, ok = os.LookupEnv("DB_SCHEMA")
	if !ok {
		panic("DB_SCHEMA env not set")
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

	GlobalEnv.Google.ClientID, ok = os.LookupEnv("GOOGLE_CLIENT_ID")
	if !ok {
		panic("GOOGLE_CLIENT_ID env not set")
	}
	GlobalEnv.Google.ClientSecret, ok = os.LookupEnv("GOOGLE_CLIENT_SECRET")
	if !ok {
		panic("GOOGLE_CLIENT_SECRET env not set")
	}
}
