package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TELEGRAM_API_TOKEN   string
	URL_GET_SCHEDULE_API string
	ULOGIN               string
	UPASSWORD            string
	ADMIN_USERNAME		 string
	TG_ID				 int64
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}

	telegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	urlGetScheduleAPI := os.Getenv("URL_GET_SCHEDULE_API")
	ulogin := os.Getenv("ULOGIN")
	upassword := os.Getenv("UPASSWORD")
	adminUsername := os.Getenv("ADMIN_USERNAME")
	tgId, err := strconv.ParseInt(os.Getenv("TG_ID"), 10, 64)

	if err != nil {
		log.Fatalln("Error converting tg_id from string to int64")
	}

	return &Config{telegramApiToken, urlGetScheduleAPI, ulogin, upassword, adminUsername, tgId}
}
