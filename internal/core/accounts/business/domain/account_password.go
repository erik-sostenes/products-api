package domain

import (
	"strings"

	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
	"golang.org/x/crypto/bcrypt"
)

// AccountPassword (Value Object) represent the password of account
type AccountPassword struct {
	value string
}

func NewAccountPassword(value string) (AccountPassword, error) {
	if strings.TrimSpace(value) == "" {
		return AccountPassword{}, wrongs.StatusBadRequest("The password is missing.")
	}

	accountPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return AccountPassword{}, err
	}

	return AccountPassword{
		value: string(accountPassword),
	}, nil
}

// NewEncryptedAccountPassword receives an already encrypted string
func NewEncryptedAccountPassword(value string) (string, error) {
	if strings.TrimSpace(value) == "" {
		return "", wrongs.StatusBadRequest("The password is missing.")
	}

	return value, nil
}

func (a AccountPassword) String() string {
	return a.value
}
