package main

import (
	"fmt"
	"tokenService/internal/config"
	"tokenService/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := logger.SetupLogger(cfg.Env)
	log.Debug("Logger initialized")
}
