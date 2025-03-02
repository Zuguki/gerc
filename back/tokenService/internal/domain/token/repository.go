package domain

import "context"

type Repository interface {
	Create(ctx context.Context, token *Token) error
	FindAll(ctx context.Context) (token []Token, err error)
	FindById(ctx context.Context, tokenId string) (Token, error)
	FindBySymbol(ctx context.Context, symbol string) (Token, error)
}
