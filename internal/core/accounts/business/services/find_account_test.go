package services

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/erik-sostenes/products-api/internal/core/accounts/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/accounts/infrastructure/driven/db"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
)

func TestAccountHandler_Find(t *testing.T) {
	// id represents an identifier no registered in the store
	var id = "94343721-6baa-4cd5-a0b4-6c5d0419c02d"

	type queryBusFunc func() (query.QueryBus[FindAccountQuery, AccountResponse], error)

	tsc := map[string]struct {
		// id represents the primitive value with which the FindAccountQuery query will be created
		id string
		queryBusFunc
		expectedError error
	}{
		"Given a valid account that does not exist, an error of type wrongs.NotFound is expected": {
			id: id,
			queryBusFunc: func() (bus query.QueryBus[FindAccountQuery, AccountResponse], err error) {
				queryHandler := FindAccountQueryHandler{db.NewMockAccountStorer()}

				bus = make(query.QueryBus[FindAccountQuery, AccountResponse])
				if err = bus.Record(FindAccountQuery{}, &queryHandler); err != nil {
					return
				}

				return bus, nil
			},
			expectedError: wrongs.StatusNotFound(fmt.Sprintf("Resource with id %v not found.", id)),
		},
		"Given a valid account that exists, no errors are expected.": {
			id: "2bdcc2f9-9522-4daa-94db-3d0f8b1cacb8",
			queryBusFunc: func() (bus query.QueryBus[FindAccountQuery, AccountResponse], err error) {
				mock := db.NewMockAccountStorer()

				account, err := domain.NewAccount("2bdcc2f9-9522-4daa-94db-3d0f8b1cacb8", "Erik", "123erik", map[string]any{"": ""})
				if err != nil {
					return nil, err
				}

				if err := mock.Save(context.TODO(), account.AccountId(), account); err != nil {
					return nil, err
				}
				queryHandler := FindAccountQueryHandler{mock}
				bus = make(query.QueryBus[FindAccountQuery, AccountResponse])
				if err = bus.Record(FindAccountQuery{}, &queryHandler); err != nil {
					return
				}

				return bus, nil
			},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			bus, err := ts.queryBusFunc()
			if err != nil {
				t.Skip(err)
			}

			query := FindAccountQuery{
				AccountId: ts.id,
			}

			_, err = bus.Ask(context.Background(), query)
			if !errors.Is(err, ts.expectedError) {
				t.Errorf("error was expected %v, but error it was obtained %v", ts.expectedError, err)
			}
		})
	}
}
