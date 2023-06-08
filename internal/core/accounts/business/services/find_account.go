package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/accounts/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
)

// FindAccountQuery implements the command.Query interface
var _ query.Query = FindAccountQuery{}

// FindAccountQuery represents the DTO with the primitive values
//
// represents the request that want performed
type FindAccountQuery struct {
	AccountId string
}

func (FindAccountQuery) QueryId() string {
	return "find_account_query"
}

// FindAccountQueryHandler implements the query.Handler interface
var _ query.Handler[FindAccountQuery, AccountResponse] = &FindAccountQueryHandler{}

type FindAccountQueryHandler struct {
	ports.AccountStorer
}

// Handler instantiates a domain.AccountId (Domain Object) with the query primitive value
func (h *FindAccountQueryHandler) Handler(ctx context.Context, qry FindAccountQuery) (AccountResponse, error) {
	id, err := domain.NewAccountId(qry.AccountId)
	if err != nil {
		return AccountResponse{}, err
	}

	account, err := h.AccountStorer.Find(ctx, id)

	if err != nil {
		return AccountResponse{}, err
	}

	return AccountResponse{
		AccountId:       account.AccountId().String(),
		AccountUserName: account.AccountUserName().String(),
		AccountPassword: account.AccountPassword().String(),
		AccountDetails:  account.AccountDetails().Value(),
	}, err
}
