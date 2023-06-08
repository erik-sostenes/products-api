package db

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/accounts/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
)

// mockAccountStorer simulates data for acceptance test and unit test
type mockAccountStorer struct {
	cache map[domain.AccountId]domain.Account
}

// NewMockAccountStorer returns a new instance of ports.AccountStorer interface
func NewMockAccountStorer() ports.AccountStorer {
	return &mockAccountStorer{
		cache: make(map[domain.AccountId]domain.Account),
	}
}

// Save saves a resource in a map
// if the resource already exist, returns a StatusBadRequest type error
func (m mockAccountStorer) Save(_ context.Context, id domain.AccountId, account domain.Account) (err error) {
	_, ok := m.cache[id]
	if ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("resource with id %v already existing", id))
		return
	}

	m.cache[id] = account

	return
}

// Remove removes the resources by an identifier of a map
// if the resource is not found, returns a Not Found type error
func (m mockAccountStorer) Remove(_ context.Context, id domain.AccountId) (err error) {
	_, ok := m.cache[id]

	if !ok {
		err = wrongs.StatusNotFound(fmt.Sprintf("resource with id %v not found", id))
		return
	}

	delete(m.cache, id)

	return
}

// Find searches a resource by id from a map
// if the resource is not found, returns a Not Found type error
func (m mockAccountStorer) Find(_ context.Context, id domain.AccountId) (account domain.Account, err error) {
	account, ok := m.cache[id]

	if !ok {
		err = wrongs.StatusNotFound(fmt.Sprintf("resource with id %v not found", id))
		return
	}

	return
}
