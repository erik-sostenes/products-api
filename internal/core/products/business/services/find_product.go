package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
)

type FinderProduct struct {
	ports.ProductStorer
}

// Handler method that receives the request to obtain a product record by identifier, requesting it from the adapter that implements the ports.ProductStorer interface
func (f FinderProduct) Handler(ctx context.Context, identifier string) (domain.Product, error) {
	return f.ProductStorer.FindById(ctx, identifier)
}
