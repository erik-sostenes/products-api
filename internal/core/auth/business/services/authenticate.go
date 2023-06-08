package services

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/services"
	"github.com/erik-sostenes/products-api/internal/core/auth/business/domain"
	wrongsAuth "github.com/erik-sostenes/products-api/internal/core/auth/business/domain/wrongs"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/jwt"
)

type AccountAuthenticator struct {
	query.Bus[services.FindAccountQuery, services.AccountResponse]
	jwt.Token[jwt.Claims]
}

// NewAccountAuthenticator returns an instance of accountAuthenticator with the dependencies that need to authenticate the account
func NewAccountAuthenticator(bus query.Bus[services.FindAccountQuery, services.AccountResponse], token jwt.Token[jwt.Claims]) AccountAuthenticator {
	return AccountAuthenticator{
		Bus:   bus,
		Token: token,
	}
}

// Authenticate need to create a FindAccountQuery query to get the account of Data Base and valid the account and credentials math with the
// values that try check the authentication
func (a *AccountAuthenticator) Authenticate(ctx context.Context, id domain.AuthId, password domain.AuthPassword) (authResponse AuthResponse, err error) {
	accountQuery := services.FindAccountQuery{
		AccountId: id.String(),
	}

	accountResponse, err := a.Bus.Ask(ctx, accountQuery)
	if err != nil {
		return
	}

	authAccount, err := domain.NewAuthAccount(accountResponse.AccountId, accountResponse.AccountPassword)
	if err != nil {
		return
	}

	if a.ensureUserExist(authAccount, id) {
		err = wrongsAuth.InvalidAuthAccount(fmt.Sprintf("The user '%s' does not exists", authAccount.AuthId().String()))
		return
	}

	if a.ensureCredentialsAreValid(authAccount, password) {
		err = wrongsAuth.InvalidAuthCredentials(fmt.Sprintf("The credentials for '%s' are invalid", authAccount.AuthId().String()))
		return
	}

	claims := jwt.NewClaims(accountResponse)

	token, err := a.Token.Generate(claims)

	fmt.Println(err)
	return AuthResponse{
		Token: token,
	}, err
}

func (a *AccountAuthenticator) ensureUserExist(authAccount domain.AuthAccount, id domain.AuthId) bool {
	return authAccount.AuthId().String() != id.String()
}

func (a *AccountAuthenticator) ensureCredentialsAreValid(authAccount domain.AuthAccount, password domain.AuthPassword) bool {
	return authAccount.PasswordMatches(password) != nil
}
