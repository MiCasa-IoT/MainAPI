package configs

import (
	"MiCasa-API/pkg/logging"
	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load("./configs/.env")
	logging.PrintEror(err)
}
