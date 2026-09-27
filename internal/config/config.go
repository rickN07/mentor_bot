package config

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
)

type DB struct {
	User string
	Password string
	Host string
	Port int
	Database string
}

type Config struct {
	TgBotToken	string
	TgHost string
	DB DB
}

func LoadConfig(logger *slog.Logger, ) (*Config, error) {
	tgBotToken := os.Getenv("TG_TOKEN")
	if tgBotToken == "" {
		logger.Error("tg bot token is empty")
		return nil, errors.New("tg bot token is empty")
	}

	tgHost := os.Getenv("TG_HOST")
	if tgHost == "" {
		logger.Error("tg host is empty")
		return nil, errors.New("tg host is empty")
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	database := os.Getenv("DATABASE")
	if dbUser == "" || dbPassword == "" || dbHost == "" || database == "" {
	  logger.Error("db is empty")
		return nil, errors.New("db is empty")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		logger.Error("dbPort is empty")
		dbPort = 5432
	}

	return &Config{
		TgBotToken: tgBotToken,
		TgHost: tgHost,
		DB: DB{
			User: dbUser,
			Password: dbPassword,
			Host: dbHost,
			Port: dbPort,
			Database: database,
		},
	}, nil
}
