package domain

import "time"

type Token struct {
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	Symbol             string     `json:"symbol"`
	Decimals           int        `json:"decimals"`
	ContractAddress    string     `json:"contract_address"`
	OwnerWalletAddress string     `json:"owner_wallet_address"`
	DateOfCreate       time.Time  `json:"date_of_create"`
	DateOfUpdate       *time.Time `json:"date_of_update"`
}
