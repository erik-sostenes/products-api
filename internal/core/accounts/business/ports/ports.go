package ports

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/domain"
)

type (
	// AccountStorer represents the Right Side
	// manages the persistence of any data in a DB or in memory, with the required operations
	AccountStorer interface {
		// Save method that saves any data to database
		Save(context.Context, domain.AccountId, domain.Account) error
		//Remove method that removes any data to the database by means of an identifier
		Remove(context.Context, domain.AccountId) error
		// Find method that searches for a resource by identifier
		Find(context.Context, domain.AccountId) (domain.Account, error)
	}
)
