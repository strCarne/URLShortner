package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/strCarne/URLShortner/internal/config"
	"github.com/strCarne/URLShortner/internal/logger"
	"github.com/strCarne/URLShortner/internal/logger/sl"
	"github.com/strCarne/URLShortner/internal/storage"
	"github.com/strCarne/URLShortner/internal/storage/sqlite"
)

func main() {
	godotenv.Load(".env")

	cfg := config.MustLoad()

	logger := logger.SetUp(cfg.Env)
	logger.Info("starting URLShortner", slog.String("env", cfg.Env))
	logger.Debug("debug messeges are enabled")

	var storage storage.Storage
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		logger.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	err = storage.SaveURL("htpps://google.com", "google")
	if err != nil {
		logger.Error("didn't saved the link", sl.Err(err))
	}
}
