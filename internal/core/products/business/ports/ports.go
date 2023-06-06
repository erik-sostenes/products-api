package ports

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
)

type (
	// ProductStorer represents the right side, for a adapter to persist data
	ProductStorer interface {
		// Save method that persists a product in the database
		Save(context.Context, string, domain.Product) error
		// Find method that searches for all the products in the database
		Find(context.Context) ([]domain.Product, error)
	}
)
