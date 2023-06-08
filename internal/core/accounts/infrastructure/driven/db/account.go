package db

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/accounts/business/ports"
)

// accountStorer implements ports.AccountStorer interface
type accountStorer struct {
}

// NewAccountStorer returns an instance of the ports.AccountStorer interface with the generic data initialized
//
// injects the database type to be used to persist the data
func NewAccountStorer() ports.AccountStorer {
	return &accountStorer{}
}

// Save method that persists in postgresql
func (a *accountStorer) Save(ctx context.Context, id domain.AccountId, account domain.Account) error {
	return nil
}

// Find method that searches in postgresql for a user account by an identifier
func (a *accountStorer) Find(ctx context.Context, accountId domain.AccountId) (account domain.Account, err error) {
	return
}

// Remove method that removes a postgrsql account by means of a key created by the account ID
func (a *accountStorer) Remove(ctx context.Context, id domain.AccountId) (err error) {
	return
}
