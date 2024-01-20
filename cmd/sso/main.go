package main

import (
	"sso/internal/config"
	"sso/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	_ = log
}
