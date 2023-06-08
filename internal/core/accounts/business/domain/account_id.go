package domain

import "github.com/erik-sostenes/products-api/internal/shared/domain"

// AccountId(Value Object) represent the id of account
type AccountId struct {
	value string
}

func NewAccountId(value string) (AccountId, error) {
	err := domain.UuID(value).EnsureIdIsValid()
	if err != nil {
		return AccountId{}, err
	}

	return AccountId{
		value: value,
	}, nil
}

func (id AccountId) String() string {
	return id.value
}
