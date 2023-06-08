package domain

import "github.com/erik-sostenes/products-api/internal/shared/domain"

// Account(Domain Object) represents a user account
type Account struct {
	accountId       AccountId
	accountUserName AccountUserName
	accountPassword AccountPassword
	accountDetails  AccountDetails
}

// NewAccount takes primitive values and converts them into value objects that make up an Account
func NewAccount(id, username, password string, details domain.Map) (Account, error) {
	accounId, err := NewAccountId(id)
	if err != nil {
		return Account{}, err
	}

	accountUserName, err := NewAccountUserName(username)
	if err != nil {
		return Account{}, err
	}

	accountPassword, err := NewAccountPassword(password)
	if err != nil {
		return Account{}, err
	}

	accountDetails, err := NewAccountDetails(details)
	if err != nil {
		return Account{}, err
	}

	return Account{
		accounId,
		accountUserName,
		accountPassword,
		accountDetails,
	}, nil
}

func (a *Account) AccountId() AccountId {
	return a.accountId
}

func (a *Account) AccountUserName() AccountUserName {
	return a.accountUserName
}

func (a *Account) AccountPassword() AccountPassword {
	return a.accountPassword
}

func (a *Account) AccountDetails() AccountDetails {
	return a.accountDetails
}
