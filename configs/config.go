package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfig() {
	err := godotenv.Load("./configs/config.yml")
	if err != nil {
		log.Fatal("Error loading config file")
	}
}
