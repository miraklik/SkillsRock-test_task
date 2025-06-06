package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Db_host string
		Db_port string
		Db_user string
		Db_name string
		Db_pass string
	}

	Server struct {
		Port string
	}
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load .env file: %v", err)
		return nil, err
	}

	var cfg Config

	cfg.Database.Db_host = os.Getenv("DB_HOST")
	cfg.Database.Db_name = os.Getenv("DB_NAME")
	cfg.Database.Db_pass = os.Getenv("DB_PASS")
	cfg.Database.Db_port = os.Getenv("DB_PORT")
	cfg.Database.Db_user = os.Getenv("DB_USER")

	cfg.Server.Port = os.Getenv("PORT")

	return &cfg, nil
}
