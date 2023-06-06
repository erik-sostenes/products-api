package db

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
)

// mockProductStorer persists data in memory
type mockProductStorer struct {
	cache map[string]domain.Product
}

// NewMockProductStorer returns an instance of ports.ProductStorer, initializes the map
func NewMockProductStorer() ports.ProductStorer {
	return &mockProductStorer{
		cache: make(map[string]domain.Product),
	}
}

// Save method that saves a resource in a map if the resource already exist, returns a StatusBadRequest type error
func (m mockProductStorer) Save(_ context.Context, identifier string, product domain.Product) error {
	_, ok := m.cache[identifier]
	if ok {
		return wrongs.StatusBadRequest(fmt.Sprintf("Resource with id %v already existing.", identifier))
	}

	m.cache[identifier] = product

	return nil
}

// Find method searches for all the records in the map and returns them in a slice
func (m mockProductStorer) Find(_ context.Context) ([]domain.Product, error) {
	var products []domain.Product

	for _, v := range m.cache {
		products = append(products, v)
	}

	return products, nil
}
