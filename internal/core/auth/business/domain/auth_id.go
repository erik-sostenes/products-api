package domain

import (
	"github.com/erik-sostenes/products-api/internal/shared/domain"
)

// AuthId (Value Object) represents auth id account
type AuthId struct {
	value string
}

// NewAuthID returns an instance of AuthID
func NewAuthID(value string) (AuthId, error) {
	err := domain.UuID(value).EnsureIdIsValid()
	if err != nil {
		return AuthId{}, err
	}

	return AuthId{
		value: value,
	}, nil
}

func (id AuthId) String() string {
	return id.value
}
