package domain

import "github.com/erik-sostenes/products-api/internal/shared/domain"

// AccountDetails(Value Object) represent the account details
type AccountDetails struct {
	value domain.Map
}

func NewAccountDetails(value domain.Map) (AccountDetails, error) {
	return AccountDetails{
		value: value,
	}, nil
}

func (a AccountDetails) Value() domain.Map {
	return a.value
}
