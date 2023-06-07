package ports

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
)

type (
	// ProductStorer represents the right side, for a adapter to persist data
	ProductStorer interface {
		// Save method persists a product in the database
		Save(ctx context.Context, identifier string, product domain.Product) error
		// Find method searches for all the products in the database
		Find(ctx context.Context) ([]domain.Product, error)
		// FindById method searches for a product by its identifier in the database
		FindById(ctx context.Context, identifier string) (domain.Product, error)
		// Delete method deletes a product by it is identifier in the database
		Delete(ctx context.Context, identifier string) error
	}
)
