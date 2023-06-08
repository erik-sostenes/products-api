package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/auth/business/domain"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
)

// AuthenticateAccountQuery implements the command.Command interface
var _ query.Query = AuthenticateAccountQuery{}

// AuthenticateAccountQuery represents the DTO with values primitives
type AuthenticateAccountQuery struct {
	Id       string
	Password string
}

func (AuthenticateAccountQuery) QueryId() string {
	return "authenticate_account_query"
}

// AuthenticateAccountQueryHandler implements the command.Handler interface
var _ query.Handler[AuthenticateAccountQuery, AuthResponse] = (*AuthenticateAccountQueryHandler)(nil)

type AuthenticateAccountQueryHandler struct {
	AccountAuthenticator
}

// Handler instance a domain.AuthAccount (Domain Object) with the command primitives values
// after domain.AuthAccount is sent t Authenticator port
func (h *AuthenticateAccountQueryHandler) Handler(ctx context.Context, qry AuthenticateAccountQuery) (AuthResponse, error) {
	id, err := domain.NewAuthID(qry.Id)
	if err != nil {
		return AuthResponse{}, err
	}
	password, err := domain.NewAuthPassword(qry.Password)
	if err != nil {
		return AuthResponse{}, err
	}

	return h.Authenticate(ctx, id, password)
}
