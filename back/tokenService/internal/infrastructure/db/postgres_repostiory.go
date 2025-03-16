package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"strings"
	"tokenService/internal/domain/token"
	"tokenService/pkg/client/postgres"
)

type repository struct {
	client postgres.Client
	log    *slog.Logger
}

func (r *repository) Create(ctx context.Context, token *domain.Token) error {
	q := `
INSERT INTO db 
	(name, symbol, decimals, contract_address, owner_wallet_address, date_of_create)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING id;
`
	r.log.Debug(fmt.Sprintf("SQL Request: %s", formatQuery(q)))

	if err := r.client.QueryRow(ctx, q, token.Name, token.Symbol, token.Decimals, token.ContractAddress, token.OwnerWalletAddress, token.DateOfCreate).Scan(&token.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			errors.As(err, &pgErr)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Error(), pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState))
			r.log.Error(newErr.Error())
			return newErr
		}
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (token []domain.Token, err error) {
	q := `
SELECT 
	id, name, symbol, decimals, contract_address, owner_wallet_address, date_of_create, date_of_update
FROM db
`

	r.log.Debug(fmt.Sprintf("SQL Request: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	tokens := make([]domain.Token, 0)
	for rows.Next() {
		var tkn domain.Token

		err = rows.Scan(&tkn.ID, &tkn.Name, &tkn.Symbol, &tkn.Decimals, &tkn.ContractAddress, &tkn.OwnerWalletAddress, &tkn.DateOfCreate, &tkn.DateOfUpdate)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, tkn)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *repository) FindById(ctx context.Context, tokenId string) (domain.Token, error) {
	q := `
SELECT 
	id, name, symbol, decimals, contract_address, owner_wallet_address, date_of_create, date_of_update
FROM db
WHERE id = $1
`

	r.log.Debug(fmt.Sprintf("SQL Request: %s", formatQuery(q)))
	var tkn domain.Token
	err := r.client.QueryRow(ctx, q, tokenId).Scan(&tkn.ID, &tkn.Name, &tkn.Symbol, &tkn.Decimals, &tkn.ContractAddress, &tkn.OwnerWalletAddress, &tkn.DateOfCreate, &tkn.DateOfUpdate)
	if err != nil {
		return domain.Token{}, err
	}

	return tkn, nil
}

func (r *repository) FindBySymbol(ctx context.Context, symbol string) (domain.Token, error) {
	q := `
SELECT 
	id, name, symbol, decimals, contract_address, owner_wallet_address, date_of_create, date_of_update
FROM db
WHERE symbol = $1
`

	r.log.Debug(fmt.Sprintf("SQL Request: %s", formatQuery(q)))
	var tkn domain.Token
	err := r.client.QueryRow(ctx, q, symbol).Scan(&tkn.ID, &tkn.Name, &tkn.Symbol, &tkn.Decimals, &tkn.ContractAddress, &tkn.OwnerWalletAddress, &tkn.DateOfCreate, &tkn.DateOfUpdate)
	if err != nil {
		return domain.Token{}, err
	}

	return tkn, nil
}

func NewRepository(client postgres.Client, log *slog.Logger) domain.Repository {
	return &repository{
		client: client,
		log:    log,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
