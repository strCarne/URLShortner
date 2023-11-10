package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/strCarne/URLShortner/internal/config"
	"github.com/strCarne/URLShortner/internal/logger"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	cfg := config.MustLoad()

	logger := logger.SetUp(cfg.Env)
	logger.Info("starting URLShortner", slog.String("env", cfg.Env))
	logger.Debug("debug messeges are enabled")
}
