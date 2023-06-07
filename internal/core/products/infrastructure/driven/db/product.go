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

// Find method searches for all the records in postgresql and returns them in a slice
func (p ProductStorer) Find(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	return products, nil
}

// FindById method searches for a record by identifier in postgresql
func (p ProductStorer) FindById(ctx context.Context, id string) (domain.Product, error) {
	var product domain.Product

	return product, nil
}
