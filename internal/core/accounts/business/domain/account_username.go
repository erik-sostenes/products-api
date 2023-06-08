package domain

import (
	"strings"

	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
)

// AccountUserName(Value Object) represent the user name of account
type AccountUserName struct {
	value string
}

func NewAccountUserName(value string) (AccountUserName, error) {
	if strings.TrimSpace(value) == "" {
		return AccountUserName{}, wrongs.StatusBadRequest("The password is missing.")
	}
	return AccountUserName{
		value: value,
	}, nil
}

func (a AccountUserName) String() string {
	return a.value
}
