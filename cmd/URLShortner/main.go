package main

import (
	"github.com/strCarne/URLShortner/internal/config"
	"github.com/strCarne/URLShortner/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.SetUp(cfg.Env)
}
