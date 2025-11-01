package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TELEGRAM_API_TOKEN   string
	URL_GET_SCHEDULE_API string
	ULOGIN               string
	UPASSWORD            string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}

	telegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	urlGetScheduleAPI := os.Getenv("URL_GET_SCHEDULE_API")
	ulogin := os.Getenv("ULOGIN")
	upassword := os.Getenv("UPASSWORD")

	return &Config{telegramApiToken, urlGetScheduleAPI, ulogin, upassword}
}
