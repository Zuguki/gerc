package main

import (
	"context"
	"fmt"
	"time"
	"tokenService/internal/config"
	"tokenService/internal/domain/token"
	"tokenService/internal/infrastructure/db"
	"tokenService/pkg/client/postgres"
	"tokenService/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := logger.SetupLogger(cfg.Env)
	log.Debug("Logger initialized")

	postgreSQLClient, err := postgres.NewClient(context.TODO(), cfg.RetryConfig, cfg.Storage.Postgres)
	if err != nil {
		log.Error(fmt.Sprintf("Error initializing postgres client: %v", err))
		return
	}

	repository := db.NewRepository(postgreSQLClient, log)

	newToken := domain.Token{
		Name:               "GG",
		Symbol:             "GGG",
		Decimals:           18,
		ContractAddress:    "contract",
		OwnerWalletAddress: "owner",
		DateOfCreate:       time.Now(),
		DateOfUpdate:       nil,
	}

	err = repository.Create(context.TODO(), &newToken)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating token: %v", err))
		return
	}

	log.Info(fmt.Sprintf("Created token: %v", newToken))
}
