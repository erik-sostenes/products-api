package db

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
)

// ProductStorer implements the ports.ProductStorer interface and persist the data in postgresql
type ProductStorer struct {
}

// NewProductStorer returns an instance ports.ProductStorer
func NewProductStorer() ports.ProductStorer {
	return ProductStorer{}
}

// Save method that persist the product a postgresql
func (p ProductStorer) Save(ctx context.Context, identifier string, product domain.Product) error {
	return nil
}
