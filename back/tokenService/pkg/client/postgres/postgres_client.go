package postgres

import (
	"context"
	"fmt"
	"log"
	"tokenService/internal/config"
	"tokenService/pkg/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, rc config.RetryConfig, pc config.PostgresConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pc.UserName, pc.Password, pc.Host, pc.Port, pc.Database)
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, rc.Timeout)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, rc.MaxAttempts, rc.Timeout)

	if err != nil {
		log.Fatalf("Failed to connect to postgres database: %v", err)
	}

	return pool, nil
}
