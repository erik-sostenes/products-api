package drives

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/core/products/infrastructure/driven/db"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
	"github.com/labstack/echo/v4"
)

func TestProductHandler_Create(t *testing.T) {
	type HandlerFunc func() (echo.HandlerFunc, error)

	tsc := map[string]struct {
		HandlerFunc
		*http.Request
		expectedStatusCode int
	}{
		"Given a valid non-existing product, a status code 201 is expected": {
			HandlerFunc: func() (echo.HandlerFunc, error) {
				productCommandHandler := services.NewCreateProductCommandHandler(db.NewMockProductStorer())

				cmdBus := make(command.CommandBus[services.ProductCommand])
				if err := cmdBus.Record(services.ProductCommand{}, &productCommandHandler); err != nil {
					return nil, err
				}

				return CreateProduct(cmdBus), nil
			},
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/products/create/1e737f50-07f1-4d1b-9c3a-62f4d38559a9",
				strings.NewReader(`
					{
						"title": "Celular Samsung Galaxy A54",
						"price": 5499,
						"image_url": "https://http2.mlstatic.com/D_NQ_NP_987811-MLM54517443931_032023-O.webp",
						"offer": true,
						"available": true,
						"rating": 4.8,
						"sales": {
							"days": 60,
							"amount": 30000
						}
					}`,
				),
			),
			expectedStatusCode: http.StatusCreated,
		},
		"Given an invalid non-existing product, a status code 422 is expected": {
			HandlerFunc: func() (echo.HandlerFunc, error) {
				productCommandHandler := services.NewCreateProductCommandHandler(db.NewMockProductStorer())

				cmdBus := make(command.CommandBus[services.ProductCommand])
				if err := cmdBus.Record(services.ProductCommand{}, &productCommandHandler); err != nil {
					return nil, err
				}

				return CreateProduct(cmdBus), nil
			},
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/products/create/1e737f50-07f1-4d1b-9c3a-62f4d38559a9",
				strings.NewReader(`
					{
						"title": "Celular Samsung Galaxy A54",
						"price": 5,499,
						"image_url": "https://http2.mlstatic.com/D_NQ_NP_987811-MLM54517443931_032023-O.webp",
						"offer": true,
						"available": "true",
						"rating": 4.u
						"sales": {
							"days": 60,
							"amount": "ds998j"
						}
					}`,
				),
			),
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			handlerFunc, err := ts.HandlerFunc()
			if err != nil {
				t.Skip(err)
			}

			e.PUT("/api/v1/products/create/:id", handlerFunc)

			resp := httptest.NewRecorder()
			req := ts.Request
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e.ServeHTTP(resp, req)

			if resp.Code != ts.expectedStatusCode {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}
		})
	}
}
