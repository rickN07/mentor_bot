package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"

	"github.com/rickN07/mentor_bot/internal/clients/telegram"
	"github.com/rickN07/mentor_bot/internal/config"
	storage "github.com/rickN07/mentor_bot/internal/storage/postgresql"
)

func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return logger
}

func main() {
	logger := setupLogger()

	// Получаем путь к текущему файлу Go
	_, filename, _, _ := runtime.Caller(0)
	// Вычисляем корень (поднимитесь на столько уровней вверх, сколько нужно)
	dir := filepath.Join(filepath.Dir(filename), "../")
	// Загружаем файл по абсолютному пути
	if err := godotenv.Load(filepath.Join(dir, ".env")); err != nil {
		logger.Warn("ENV is not found")
		os.Exit(1)
	}

	// Создаем контекст с обработкой сигналов для graceful shutdown
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer cancel()

	cfg, err := config.LoadConfig(logger)
	if err != nil {
		logger.Error("Can't get config", "err", err)
		os.Exit(1)
	}

	tgClient := telegram.New(cfg.TgBotToken, cfg.TgHost)
	_ = tgClient

	dbPool, err := storage.NewPostgresClient(logger, context.Background(), cfg.DB, 3)
	if err != nil {
		logger.Error("", "err", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	fmt.Println("EEEEE")
}
