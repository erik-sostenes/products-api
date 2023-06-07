package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
)

type FinderProducts struct {
	ports.ProductStorer
}

// Handler method that receives the request to get all product records, requesting it to the adapter implementing the ports.ProductStorer interface
func (f FinderProducts) Handler(ctx context.Context) ([]domain.Product, error) {
	return f.ProductStorer.Find(ctx)
}
