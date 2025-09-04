package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"multiplicator/internal/models"
	"os"
)

func LoadConfig() *models.Config {
	const op = "config.LoadConfig"
	var config models.Config

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Loc: %s; Err: Error loading .env file", op)
	}
	configPath := os.Getenv("CONFIG_PATH")

	err = cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("Loc: %s; Err: %v", op, err)
	}
	return &config
}
