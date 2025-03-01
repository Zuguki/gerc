package token

import "context"

type Repository interface {
	Create(ctx context.Context, token *Token) error
	FindAll(ctx context.Context) (token []Token, err error)
	FindById(ctx context.Context, tokenId string) (Token, error)
	FindBySymbol(ctx context.Context, symbol string) (Token, error)
	Update(ctx context.Context, token Token) error
	Delete(ctx context.Context, tokenId string) error
}
